package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
)

// NodeAdd add a node
// @Summary add a node
// @Author l.jiang.1024@gmail.com
// @Description add a node
// @Tags Node
// @Router /api/v1/nodes [post]
func NodeAdd(c *gin.Context) {
	req := &request.Node{}
	if err := c.BindJSON(req); err != nil {
		Fail(c, err)
		return
	}

	if _, err := handler.Node.Add(c.GetUint("user_id"), req); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}

// NodeUpdate update a node
// @Summary update a node
// @Author l.jiang.1024@gmail.com
// @Description update a node
// @Tags Node
// @Router /api/v1/nodes/:node_id [put]
func NodeUpdate(c *gin.Context) {
	var req request.Node
	if err := c.BindJSON(&req); err != nil {
		Fail(c, err)
		return
	}
	if err := handler.Node.Update(c.GetUint("user_id"), c.GetUint("node_id"), &req); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}

// NodeList get node list
// @Summary get node list
// @Author l.jiang.1024@gmail.com
// @Description get node list
// @Tags Node
// @Router /api/v1/nodes [get]
func NodeList(c *gin.Context) {
	respData, err := handler.Node.List(c.GetUint("user_id"))
	if err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", respData)
}

// NodeDelete delete a node
// @Summary delete a node
// @Author l.jiang.1024@gmail.com
// @Description delete a node
// @Tags Node
// @Router /api/v1/nodes/:node_id [delete]
func NodeDelete(c *gin.Context) {
	if err := handler.Node.Delete(c.GetUint("node_id")); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}
