package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/util"
	"github.com/ismdeep/parser"
)

func init() {
	auth.GET("/api/v1/tasks/:task_id/triggers", TriggerList)
	auth.POST("/api/v1/tasks/:task_id/triggers", TriggerAdd)
	auth.PUT("/api/v1/tasks/:task_id/triggers/:trigger_id", TriggerUpdate)
	auth.DELETE("/api/v1/tasks/:task_id/triggers/:trigger_id", TriggerDelete)
}

// TriggerAdd add a trigger
// @Summary add a trigger
// @Author l.jiang.1024@gmail.com
// @Description add a trigger
// @Tags Trigger
// @Router	/api/v1/tasks/:task_id/triggers [post]
func TriggerAdd(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	req := &request.Trigger{}
	err2 := c.BindJSON(req)
	if err := util.FirstError(err1, err2); err != nil {
		Fail(c, err.Error())
		return
	}

	if _, err := handler.Trigger.Add(c.GetUint("user_id"), taskID, req); err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", nil)
}

// TriggerList get trigger list
// @Summary get trigger list
// @Author l.jiang.1024@gmail.com
// @Description get trigger list
// @Tags Trigger
// @Router /api/v1/tasks/:task_id/triggers [get]
func TriggerList(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	if err := util.FirstError(err1); err != nil {
		Fail(c, err.Error())
		return
	}

	respData, err := handler.Trigger.List(c.GetUint("user_id"), taskID)
	if err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", respData)
}

// TriggerUpdate update a trigger
// @Summary update a trigger
// @Author l.jiang.1024@gmail.com
// @Description update a trigger
// @Tags Trigger
// @Router /api/v1/tasks/:task_id/triggers/:trigger_id [put]
func TriggerUpdate(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	triggerID, err2 := parser.ToUint(c.Param("trigger_id"))
	req := &request.Trigger{}
	err3 := c.BindJSON(req)
	if err := util.FirstError(err1, err2, err3); err != nil {
		Fail(c, err.Error())
		return
	}

	err := handler.Trigger.Update(c.GetUint("user_id"), taskID, triggerID, req)
	if err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", nil)
}

// TriggerDelete delete a trigger
// @Summary delete a trigger
// @Author l.jiang.1024@gmail.com
// @Description delete a trigger
// @Tags Trigger
// @Router /api/v1/tasks/:task_id/triggers/:trigger_id [delete]
func TriggerDelete(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	triggerID, err2 := parser.ToUint(c.Param("trigger_id"))
	if err := util.FirstError(err1, err2); err != nil {
		Fail(c, err.Error())
		return
	}
	if err := handler.Trigger.Delete(taskID, triggerID); err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", nil)
}
