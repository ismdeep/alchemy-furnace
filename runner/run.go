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

	// 1. 创建目录
	randKey := rand.TimeBasedFormat("{datetime}-{hex}", rand.WithHexLen(32))
	workDir := fmt.Sprintf("/tmp/%v", randKey)
	if err := os.MkdirAll(workDir, 0777); err != nil {
		return 1, err
	}

	// 2. 写入脚本
	if err := ioutil.WriteFile(fmt.Sprintf("%v/main.bash", workDir), []byte(run.Task.BashContent), 0777); err != nil {
		return 1, err
	}

	// 3. 执行命令
	cmd := exec.Command("bash", "main.bash")
	cmd.Env = os.Environ()
	for _, env := range strings.Split(run.Trigger.Environment, "\n") {
		cmd.Env = append(cmd.Env, env)
	}
	cmd.Dir = workDir
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return 1, err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return 1, err
	}
	if err := cmd.Start(); err != nil {
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
	if err := cmd.Wait(); err != nil {
		return 1, err
	}

	// 4. 返回数据
	return cmd.ProcessState.ExitCode(), nil
}
