package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/util"
)

// TaskList get task list
// @Summary get task list
// @Author l.jiang.1024@gmail.com
// @Description get task list
// @Tags Task
// @Success 200 {object} []response.Task
// @Router /api/v1/tasks [get]
func TaskList(c *gin.Context) {
	items := handler.Task.List(c.GetUint("user_id"))
	Success(c, "", items)
}

// TaskCreate create a task
// @Summary creates a task
// @Author l.jiang.1024@gmail.com
// @Description create a task
// @Tags Task
// @Param Authorization	header	string true "Bearer 31a165ba1be6dec616b1f8f3207b4273"
// @Param req body	request.Task true "JSON数据"
// @Router	/api/v1/tasks [post]
func TaskCreate(c *gin.Context) {
	req := &request.Task{}
	err1 := c.BindJSON(req)
	if err := util.FirstError(err1); err != nil {
		Fail(c, err.Error())
		return
	}

	if _, err := handler.Task.Create(c.GetUint("user_id"), req); err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", nil)
}
