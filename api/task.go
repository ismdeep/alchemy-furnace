package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/util"
	"github.com/ismdeep/parser"
)

func init() {
	auth.GET("/api/v1/tasks", TaskList)
	auth.POST("/api/v1/tasks", TaskCreate)
	permCheckAuth.PUT("/api/v1/tasks/:task_id", TaskUpdate)
	permCheckAuth.GET("/api/v1/tasks/:task_id", TaskDetail)
	permCheckAuth.DELETE("/api/v1/tasks/:task_id", TaskDelete)
}

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
		Fail(c, err)
		return
	}

	if _, err := handler.Task.Create(c.GetUint("user_id"), req); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}

// TaskUpdate update a task
// @Summary updates a task
// @Author l.jiang.1024@gmail.com
// @Description update a task
// @Tags Task
// @Param Authorization	header	string true "Bearer 31a165ba1be6dec616b1f8f3207b4273"
// @Param req body	request.Task true "JSON数据"
// @Router /api/v1/tasks/:task_id [put]
func TaskUpdate(c *gin.Context) {
	req := &request.Task{}
	err1 := c.BindJSON(req)
	taskID, err2 := parser.ToUint(c.Param("task_id"))
	if err := util.FirstError(err1, err2); err != nil {
		Fail(c, err)
		return
	}

	if err := handler.Task.Update(taskID, req); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}

// TaskDetail get task detail
// @Summary get task detail
// @Author l.jiang.1024@gmail.com
// @Description get task detail
// @Tags Task
// @Router /api/v1/tasks/:id [get]
func TaskDetail(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	if err := util.FirstError(err1); err != nil {
		Fail(c, err)
		return
	}

	respData, err := handler.Task.Detail(taskID)
	if err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", respData)
}

// TaskDelete delete a task
// @Summary delete a task
// @Author l.jiang.1024@gmail.com
// @Description delete a task
// @Tags Task
// @Router /api/v1/tasks/:task_id [delete]
func TaskDelete(c *gin.Context) {
	taskID, err1 := parser.ToUint(c.Param("task_id"))
	if err := util.FirstError(err1); err != nil {
		Fail(c, err)
		return
	}

	if err := handler.Task.Delete(taskID); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}
