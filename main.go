package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/api"
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/alchemy-furnace/executor"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/log"
	"github.com/ismdeep/rand"
	"github.com/robfig/cron/v3"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
)

type Task struct {
	ID       string
	Name     string
	Bash     string
	Cron     string
	ExitCode int
}

func (receiver *Task) RunTask() (exeID string) {
	exeID = executor.GenerateExecutor()

	dir := fmt.Sprintf("%v/%v", config.WorkDir, rand.HexStr(32))
	if err := os.MkdirAll(dir, 0700); err != nil {
		executor.PushMsg(exeID, executor.TypeStderr, err.Error())
		return
	}

	bashBytes, err := ioutil.ReadFile(fmt.Sprintf("%v/shells/%v", config.WorkDir, receiver.Bash))
	if err != nil {
		executor.PushMsg(exeID, executor.TypeStderr, err.Error())
		return
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%v/%v", dir, receiver.ID), bashBytes, 0700); err != nil {
		executor.PushMsg(exeID, executor.TypeStderr, err.Error())
		return
	}
	cmd := exec.Command("bash", fmt.Sprintf("%v/%v", dir, receiver.ID))
	cmd.Dir = dir

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		executor.PushMsg(exeID, executor.TypeStderr, err.Error())
		return
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		executor.PushMsg(exeID, executor.TypeStderr, err.Error())
		return
	}

	if err := cmd.Start(); err != nil {
		executor.PushMsg(exeID, executor.TypeStderr, err.Error())
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	wg.Add(1)
	listenFunc := func(p io.ReadCloser, outType int) {
		line := ""
		for {
			t := bufio.NewReader(p)
			s, isPrefix, err := t.ReadLine()
			if err != nil {
				if line != "" {
					executor.PushMsg(exeID, outType, line)
				}
				if err.Error() != "EOF" {
					executor.PushMsg(exeID, executor.TypeStderr, err.Error())
				}
				break
			}
			line = line + string(s)
			if isPrefix {
				continue
			}
			executor.PushMsg(exeID, outType, line)
			line = ""
		}
		wg.Done()
	}

	go listenFunc(stdoutPipe, executor.TypeStdout)
	go listenFunc(stderrPipe, executor.TypeStderr)

	wg.Wait()

	if err := cmd.Wait(); err != nil {
		executor.PushMsg(exeID, executor.TypeStderr, err.Error())
		return
	}

	receiver.ExitCode = cmd.ProcessState.ExitCode()

	if err := os.RemoveAll(dir); err != nil {
		executor.PushMsg(exeID, executor.TypeStderr, err.Error())
	}
	return
}

func (receiver *Task) Run() {
	exeID := receiver.RunTask()
	content, err := executor.DumpLog(exeID)
	executor.DestroyExecutor(exeID)
	if err != nil {
		log.Error(receiver.ID, log.FieldErr(err))
		return
	}

	if err := model.DB.Create(&model.Run{
		TaskID:   receiver.ID,
		Name:     receiver.Name,
		ExitCode: receiver.ExitCode,
		Content:  content,
	}).Error; err != nil {
		log.Error(receiver.ID, log.FieldErr(err))
		return
	}

}

func main() {
	c := cron.New(cron.WithSeconds())
	for _, task := range config.Config.Tasks {
		job := &Task{
			ID:   task.ID,
			Name: task.Name,
			Bash: task.Bash,
			Cron: task.Cron,
		}
		_, _ = c.AddJob(task.Cron, job)
	}
	c.Start()

	gin.SetMode(gin.ReleaseMode)
	eng := gin.Default()
	eng.GET("/api/v1/tasks", api.TaskList)
	eng.GET("/api/v1/tasks/:task_id/runs", api.RunList)
	eng.GET("/api/v1/tasks/:task_id/runs/:run_id", api.RunDetail)
	if err := eng.Run("0.0.0.0:8080"); err != nil {
		panic(err)
	}
}
