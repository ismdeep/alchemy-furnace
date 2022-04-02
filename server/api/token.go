package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/util"
)

// TokenList get token list
// @Summary get token list
// @Author l.jiang.1024@gmail.com@uniontech.com
// @Description get token list
// @Tags Token
// @Router /api/v1/tokens [get]
func TokenList(c *gin.Context) {
	respData, err := handler.Token.List()
	if err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", respData)
}

// TokenAdd add a token
// @Summary add a token
// @Author l.jiang.1024@gmail.com
// @Description add a token
// @Tags Token
// @Router /api/v1/tokens [post]
func TokenAdd(c *gin.Context) {
	req := &request.Token{}
	err1 := c.BindJSON(req)
	if err := util.FirstError(err1); err != nil {
		Fail(c, err)
		return
	}

	_, tokenKey, err := handler.Token.Add(req)
	if err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", map[string]interface{}{"token_key": tokenKey})
}

// TokenUpdate update a token
// @Summary update a token
// @Author l.jiang.1024@gmail.com
// @Description update a token
// @Tags Token
// @Router /api/v1/tokens/:token_id [put]
func TokenUpdate(c *gin.Context) {
	var req request.Token
	err1 := c.BindJSON(&req)
	if err := util.FirstError(err1); err != nil {
		Fail(c, err)
		return
	}

	if err := handler.Token.Update(c.GetUint("token_id"), &req); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}

// TokenDelete delete a token
// @Summary delete a token
// @Author l.jiang.1024@gmail.com
// @Description delete a token
// @Tags Token
// @Router /api/v1/tokens/:token_id [delete]
func TokenDelete(c *gin.Context) {
	if err := handler.Token.Delete(c.GetUint("token_id")); err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", nil)
}
