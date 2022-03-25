package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
)

// TriggerAdd add a trigger
// @Summary add a trigger
// @Author l.jiang.1024@gmail.com
// @Description add a trigger
// @Tags Trigger
// @Router	/api/v1/tasks/:task_id/triggers [post]
func TriggerAdd(c *gin.Context) {
	var req request.Trigger
	if err := c.BindJSON(&req); err != nil {
		Fail(c, err)
		return
	}

	if _, err := handler.Trigger.Add(c.GetUint("user_id"), c.GetUint("task_id"), &req); err != nil {
		Fail(c, err)
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
	respData, err := handler.Trigger.List(c.GetUint("user_id"), c.GetUint("task_id"))
	if err != nil {
		Fail(c, err)
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
	var req request.Trigger
	if err := c.BindJSON(&req); err != nil {
		Fail(c, err)
		return
	}

	if err := handler.Trigger.Update(c.GetUint("user_id"), c.GetUint("task_id"), c.GetUint("trigger_id"), &req); err != nil {
		Fail(c, err)
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
	if err := handler.Trigger.Delete(c.GetUint("task_id"), c.GetUint("trigger_id")); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}
