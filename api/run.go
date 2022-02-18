package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/ismdeep/alchemy-furnace/executor"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/util"
	"github.com/ismdeep/log"
	"github.com/ismdeep/parser"
	"net/http"
)

func init() {
	auth.POST("/api/v1/tasks/:task_id/triggers/:trigger_id/runs", RunCreate) // Start to run a task by trigger
	auth.GET("/api/v1/tasks/:task_id/runs", RunList)
	auth.GET("/api/v1/tasks/:task_id/runs/:run_id", RunDetail)
	noAuth.GET("/api/v1/tasks/:task_id/runs/:run_id/log", RunLog) // Get run log with websocket
}

// RunList get task run list
// @Summary get task run list
// @Author l.jiang.1024@gmail.com
// @Description get task run list
// @Tags Task
// @Router /api/v1/tasks/:task_id/runs [get]
func RunList(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	page, err2 := parser.ToInt(c.DefaultQuery("page", "1"))
	size, err3 := parser.ToInt(c.DefaultQuery("size", "10"))
	if err := util.FirstError(err1, err2, err3); err != nil {
		Fail(c, err.Error())
		return
	}

	tasks, total, err := handler.Run.List(taskID, page, size)
	if err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", map[string]interface{}{"total": total, "list": tasks})
}

// RunDetail get run detail
// @Summary get run detail
// @Author l.jiang.1024@gmail.com
// @Description get run detail
// @Tags Task
// @Success 200 {object} response.Run
// @Router /api/v1/tasks/:task_id/runs/:run_id [get]
func RunDetail(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	runID, err2 := parser.ToUint(c.Param("run_id"))
	if err := util.FirstError(err1, err2); err != nil {
		Fail(c, err.Error())
		return
	}

	respData, err := handler.Run.Detail(taskID, runID)
	if err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", respData)
	return
}

// RunCreate create a run for task
// @Summary creates a run for task
// @Author l.jiang.1024@gmail.com
// @Description create a run for task
// @Tags Task
// @Router /api/v1/tasks/:task_id/triggers/:trigger_id/runs [post]
func RunCreate(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	triggerID, err2 := parser.ToUint(c.Param("trigger_id"))
	if err := util.FirstError(err1, err2); err != nil {
		Fail(c, err.Error())
		return
	}

	fmt.Println(taskID)

	if err := handler.Run.Start(triggerID); err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", nil)
}

// RunLog get run log with websocket
// @Summary get run log with websocket
// @Author l.jiang.1024@gmail.com
// @Description get run log with websocket
// @Tags Task
// @Router /api/v1/tasks/:task_id/runs/:run_id/log [get]
func RunLog(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	runID, err2 := parser.ToUint(c.Param("run_id"))
	if err := util.FirstError(err1, err2); err != nil {
		Fail(c, err.Error())
		return
	}

	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func() {
		if err := ws.Close(); err != nil {
			log.Error("RunLog", log.FieldErr(err))
		}
	}()

	// 1. 获取日志信息
	run := &model.Run{}
	if err := model.DB.Where("id=?", runID).First(run).Error; err != nil {
		_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("[ERROR] %v\n", err.Error())))
		_ = ws.WriteMessage(websocket.CloseMessage, []byte(""))
		return
	}

	if run.TaskID != taskID {
		_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("[ERROR] bad request\n")))
		_ = ws.WriteMessage(websocket.CloseMessage, []byte(""))
		return
	}

	if run.Status == model.RunEnumsStatusRunning {
		listener, listenerID, err := executor.GenerateListener(run.ExecutorID)
		if err != nil {
			_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("[ERROR] %v\n", err.Error())))
			_ = ws.WriteMessage(websocket.CloseMessage, []byte(""))
			return
		}

		for {
			v := <-listener
			if v == nil {
				break
			}
			if v.Line == executor.EOF {
				break
			}
			_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("[%v] %v\n", v.OutputAt.Format("2006-01-02 15:04:05"), v.Line)))
		}

		_ = executor.DestroyListener(run.ExecutorID, listenerID)
		_ = ws.WriteMessage(websocket.CloseMessage, []byte(""))
		return
	}

	logs := make([]executor.ExeLog, 0)
	_ = json.Unmarshal([]byte(run.CmdLog), &logs)
	for _, v := range logs {
		_ = ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("[%v] %v\n", v.OutputAt.Format("2006-01-02 15:04:05"), v.Line)))
	}

	_ = ws.WriteMessage(websocket.CloseMessage, []byte(""))
}
