package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/api"
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/alchemy-furnace/executor"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/schema"
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/rand"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"sync"
)

type Task struct {
	ID       uint
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
		return
	}

	if err := model.DB.Create(&model.Run{
		TaskID:   receiver.ID,
		Name:     receiver.Name,
		ExitCode: receiver.ExitCode,
		Content:  content,
	}).Error; err != nil {
		return
	}

}

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		bytes, err := jwt.VerifyToken(token)
		if err != nil {
			c.JSON(200, map[string]interface{}{"code": 403, "msg": "token verification failed"})
			c.Abort()
			return
		}

		u := schema.LoginUser{}
		if err := json.Unmarshal([]byte(bytes), &u); err != nil {
			c.JSON(200, map[string]interface{}{"code": 403, "msg": "token verification failed"})
			c.Abort()
			return
		}

		c.Set("user_id", u.ID)
		c.Set("username", u.Username)
		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	eng := gin.Default()
	eng.POST("/api/v1/sign-up", api.UserRegister)
	auth := eng.Group("/api/v1")
	auth.Use(Authorization())
	auth.GET("/api/v1/tasks", api.TaskList)
	auth.POST("/api/v1/tasks", api.TaskCreate)
	auth.GET("/api/v1/tasks/:task_id/runs", api.RunList)
	auth.GET("/api/v1/tasks/:task_id/runs/:run_id", api.RunDetail)
	fmt.Printf("Listening... %v\n", config.Config.Bind)
	if err := eng.Run(config.Config.Bind); err != nil {
		panic(err)
	}
}
