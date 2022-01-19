package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
)

// TaskList get task list
// @Summary get task list
// @Author l.jiang.1024@gmail.com
// @Description get task list
// @Tags Task
// @Success 200 {object} []response.Task
// @Router /api/v1/tasks [get]
func TaskList(c *gin.Context) {
	items := handler.Task.List()
	Success(c, "", items)
}
