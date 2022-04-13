package runner

import (
	"bufio"
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

	// 2. 准备数据
	randKey := rand.TimeBasedFormat("{datetime}-{hex}", rand.WithHexLen(32))
	workDir := fmt.Sprintf("/tmp/%v", randKey)
	if err := os.MkdirAll(workDir, 0777); err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, err.Error())
		return 1, err
	}
	if err := ioutil.WriteFile(fmt.Sprintf("%v/main.bash", workDir), []byte(fmt.Sprintf("%v\n\n%v", run.Trigger.Environment, run.Task.BashContent)), 0777); err != nil {
		executor.PushMsg(executorID, executor.TypeStderr, fmt.Sprintf("[ERROR] write bash file failed, err: %v", err.Error()))
		return 1, err
	}

	// 3. 执行命令
	cmdRunBashFile := exec.Command("bash", "main.bash")
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
