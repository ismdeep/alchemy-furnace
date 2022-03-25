package runner

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ismdeep/alchemy-furnace/executor"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/log"
	"github.com/ismdeep/rand"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func Run(runID uint, executorID string) (int, error) {
	// 0. 获取数据
	run := &model.Run{}
	if err := model.DB.Preload("Task").Preload("Trigger").Preload("Trigger").Where("id=?", runID).First(run).Error; err != nil {
		return 1, err
	}

	// 1. 选取运行节点
	var nodes []model.Node
	if err := model.DB.Where("user_id=?", run.Task.UserID).Find(&nodes).Error; err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, err.Error())
		return 1, err
	}
	if len(nodes) <= 0 {
		executor.PushMsg(executorID, executor.TypeStderr, "[ERROR] no available node")
		return 1, errors.New("no available node")
	}
	node := nodes[rand.Intn(len(nodes))]
	executor.PushMsg(executorID, executor.TypeStdout, fmt.Sprintf("[NODE] %v", node.Name))

	// 2. 准备数据
	randKey := rand.TimeBasedFormat("{datetime}-{hex}", rand.WithHexLen(32))
	workDir := fmt.Sprintf("/tmp/%v", randKey)
	if err := os.MkdirAll(workDir, 0777); err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, err.Error())
		return 1, err
	}
	// 2.1 写入脚本
	if err := ioutil.WriteFile(fmt.Sprintf("%v/main.bash", workDir), []byte(fmt.Sprintf("%v\n\n%v", run.Trigger.Environment, run.Task.BashContent)), 0777); err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, fmt.Sprintf("[ERROR] write bash file failed, err: %v", err.Error()))
		return 1, err
	}
	if err := ioutil.WriteFile(fmt.Sprintf("%v/ssh-key", workDir), []byte(node.SSHKey), 0600); err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, fmt.Sprintf("[ERROR] write ssh key file failed. err: %v", err.Error()))
		return 1, err
	}
	// 2.2 创建远程服务器目录
	cmdCreateRemoteDir := exec.Command(
		"ssh",
		"-o", "StrictHostKeyChecking=no",
		"-i", "ssh-key",
		fmt.Sprintf("%v@%v", node.Username, node.Host),
		"-p", fmt.Sprintf("%v", node.Port),
		fmt.Sprintf("mkdir -p %v", workDir))
	cmdCreateRemoteDir.Dir = workDir
	cmdCreateRemoteDir.Stdout = os.Stdout
	cmdCreateRemoteDir.Stderr = os.Stderr
	if err := cmdCreateRemoteDir.Run(); err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, fmt.Sprintf("[ERROR] create remote dir failed. err: %v", err.Error()))
		return 1, err
	}
	executor.PushMsg(executorID, executor.TypeStdout, "[DONE] create remote dir finished")

	defer func() {
		// clean remote workdir
		cmdCleanRemote := exec.Command("ssh",
			"-o", "StrictHostKeyChecking=no",
			"-i", "ssh-key",
			fmt.Sprintf("%v@%v", node.Username, node.Host),
			"-p", fmt.Sprintf("%v", node.Port),
			fmt.Sprintf("rm -rf %v", workDir))
		cmdCleanRemote.Dir = workDir
		cmdCleanRemote.Stdout = os.Stdout
		cmdCleanRemote.Stderr = os.Stderr
		if err := cmdCleanRemote.Run(); err != nil {
			log.Error("run", log.Any("run_id", runID), log.Any("executor_id", executorID), log.FieldErr(err))
		}
	}()

	// 2.3 拷贝执行脚本
	cmdCopyBashFile := exec.Command(
		"scp",
		"-P", fmt.Sprintf("%v", node.Port),
		"-o", "StrictHostKeyChecking=no",
		"-i", "ssh-key",
		"main.bash",
		fmt.Sprintf("%v@%v:%v", node.Username, node.Host, workDir),
	)
	cmdCopyBashFile.Stdout = os.Stdout
	cmdCopyBashFile.Stderr = os.Stderr
	cmdCopyBashFile.Dir = workDir
	if err := cmdCopyBashFile.Run(); err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, fmt.Sprintf("[ERROR] copy bash file failed. err: %v", err.Error()))
		return 1, err
	}
	executor.PushMsg(executorID, executor.TypeStdout, "[DONE] copy bash file finished")

	// 3. 执行命令
	cmdRunBashFile := exec.Command("ssh",
		"-o", "StrictHostKeyChecking=no",
		"-i", "ssh-key",
		fmt.Sprintf("%v@%v", node.Username, node.Host),
		"-p", fmt.Sprintf("%v", node.Port),
		fmt.Sprintf("cd %v && bash ./main.bash", workDir))
	cmdRunBashFile.Env = os.Environ()
	for _, env := range strings.Split(run.Trigger.Environment, "\n") {
		cmdRunBashFile.Env = append(cmdRunBashFile.Env, env)
	}
	cmdRunBashFile.Dir = workDir
	stdoutPipe, err := cmdRunBashFile.StdoutPipe()
	if err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, fmt.Sprintf("[ERROR] redirect stdout failed. err: %v", err.Error()))
		return 1, err
	}
	stderrPipe, err := cmdRunBashFile.StderrPipe()
	if err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, fmt.Sprintf("[ERROR] redirect stderr failed. err: %v", err.Error()))
		return 1, err
	}
	if err := cmdRunBashFile.Start(); err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, fmt.Sprintf("[ERROR] start to run bash file failed. err: %v", err.Error()))
		return 1, err
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Add(1)
	listenFunc := func(p io.ReadCloser, outType int) {
		t := bufio.NewReader(p)
		line := ""
		for {
			s, isPrefix, err := t.ReadLine()
			if err != nil {
				if line != "" {
					executor.PushMsg(executorID, outType, line)
				}
				if err.Error() != "EOF" {
					executor.PushMsg(executorID, executor.TypeStderr, err.Error())
				}
				break
			}
			line = line + string(s)
			if isPrefix {
				continue
			}
			executor.PushMsg(executorID, outType, line)
			line = ""
		}
		wg.Done()
	}
	go listenFunc(stdoutPipe, executor.TypeStdout)
	go listenFunc(stderrPipe, executor.TypeStderr)
	wg.Wait()
	if err := cmdRunBashFile.Wait(); err != nil {
		return 1, err
	}

	executor.PushMsg(executorID, executor.TypeStdout, executor.EOF)

	// 4. 返回数据
	return cmdRunBashFile.ProcessState.ExitCode(), nil
}
