package runner

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ismdeep/alchemy-furnace/executor"
	"github.com/ismdeep/alchemy-furnace/model"
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
		return 1, err
	}
	if len(nodes) <= 0 {
		return 1, errors.New("no available node")
	}
	node := nodes[rand.Intn(len(nodes))]

	// 2. 准备数据
	randKey := rand.TimeBasedFormat("{datetime}-{hex}", rand.WithHexLen(32))
	workDir := fmt.Sprintf("/tmp/%v", randKey)
	if err := os.MkdirAll(workDir, 0777); err != nil {
		return 1, err
	}
	// 2.1 写入脚本
	if err := ioutil.WriteFile(fmt.Sprintf("%v/main.bash", workDir), []byte(run.Task.BashContent), 0777); err != nil {
		return 1, err
	}
	if err := ioutil.WriteFile(fmt.Sprintf("%v/ssh-key", workDir), []byte(node.SSHKey), 0600); err != nil {
		return 1, err
	}
	// 2.2 转化pem文件
	//cmd0 := exec.Command("ssh-keygen", "-f", "ssh-key", "-e", "-m", "pem")
	//cmd0.Dir = workDir
	//if err := cmd0.Run(); err != nil {
	//	return 1, err
	//}
	// 2.2 创建远程服务器目录
	cmd1 := exec.Command(
		"ssh",
		"-i",
		"ssh-key",
		fmt.Sprintf("%v@%v", node.Username, node.Host),
		"-p",
		fmt.Sprintf("%v", node.Port),
		fmt.Sprintf("mkdir -p %v", workDir))
	cmd1.Dir = workDir
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	if err := cmd1.Run(); err != nil {
		return 1, err
	}
	fmt.Println("cmd1 finished")
	// 2.3 拷贝执行脚本
	cmd2 := exec.Command(
		"scp",
		"-P",
		fmt.Sprintf("%v", node.Port),
		"-i",
		"ssh-key",
		"main.bash",
		fmt.Sprintf("%v@%v:%v", node.Username, node.Host, workDir),
	)
	cmd2.Stdout = os.Stdout
	cmd2.Stderr = os.Stderr
	cmd2.Dir = workDir
	if err := cmd2.Run(); err != nil {
		return 1, err
	}
	fmt.Println("cmd2 finished")

	// 3. 执行命令
	cmd3 := exec.Command("ssh",
		fmt.Sprintf("%v@%v", node.Username, node.Host),
		"-p",
		fmt.Sprintf("%v", node.Port),
		fmt.Sprintf("cd %v && bash main.bash", workDir))
	cmd3.Env = os.Environ()
	for _, env := range strings.Split(run.Trigger.Environment, "\n") {
		cmd3.Env = append(cmd3.Env, env)
	}
	cmd3.Dir = workDir
	stdoutPipe, err := cmd3.StdoutPipe()
	if err != nil {
		return 1, err
	}
	stderrPipe, err := cmd3.StderrPipe()
	if err != nil {
		return 1, err
	}
	if err := cmd3.Start(); err != nil {
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
	if err := cmd3.Wait(); err != nil {
		return 1, err
	}

	executor.PushMsg(executorID, executor.TypeStdout, executor.EOF)

	// 4. 返回数据
	return cmd3.ProcessState.ExitCode(), nil
}
