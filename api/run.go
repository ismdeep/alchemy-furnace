package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/util"
	"github.com/ismdeep/parser"
)

// RunList get task run list
// @Summary get task run list
// @Author l.jiang.1024@gmail.com
// @Description get task run list
// @Tags Task
// @Router /api/v1/tasks/:task_id/runs [get]
func RunList(c *gin.Context) {
	taskID := c.Param("task_id")
	page, err1 := parser.ToInt(c.DefaultQuery("page", "1"))
	size, err2 := parser.ToInt(c.DefaultQuery("size", "10"))
	if err := util.FirstError(err1, err2); err != nil {
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
	taskID := c.Param("task_id")
	runID, err := parser.ToUint(c.Param("run_id"))
	if err != nil {
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
