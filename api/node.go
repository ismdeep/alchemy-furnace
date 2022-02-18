package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/util"
	"github.com/ismdeep/parser"
)

func init() {
	auth.GET("/api/v1/nodes", NodeList)
	auth.POST("/api/v1/nodes", NodeAdd)
	permCheckAuth.PUT("/api/v1/nodes/:node_id", NodeUpdate)
	permCheckAuth.DELETE("/api/v1/nodes/:node_id", NodeDelete)
}

func NodeAdd(c *gin.Context) {
	req := &request.Node{}
	err1 := c.BindJSON(req)
	if err := util.FirstError(err1); err != nil {
		Fail(c, err)
		return
	}

	if _, err := handler.Node.Add(c.GetUint("user_id"), req); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}

func NodeUpdate(c *gin.Context) {
	nodeID, err1 := parser.ToUint(c.Param("node_id"))
	req := &request.Node{}
	err2 := c.BindJSON(req)
	if err := util.FirstError(err1, err2); err != nil {
		Fail(c, err)
		return
	}

	if err := handler.Node.Update(c.GetUint("user_id"), nodeID, req); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}

func NodeList(c *gin.Context) {
	respData, err := handler.Node.List(c.GetUint("user_id"))
	if err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", respData)
}

func NodeDelete(c *gin.Context) {
	nodeID, err1 := parser.ToUint(c.Param("node_id"))
	if err := util.FirstError(err1); err != nil {
		Fail(c, err)
		return
	}

	if err := handler.Node.Delete(nodeID); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}
