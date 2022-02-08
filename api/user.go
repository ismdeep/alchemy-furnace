package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
)

// UserRegister user register
// @Summary user register
// @Author l.jiang.1024@gmail.com
// @Description user register
// @Tags User
// @Router /api/v1/sign-up [post]
func UserRegister(c *gin.Context) {
	req := &request.User{}
	if err := c.BindJSON(req); err != nil {
		Fail(c, err.Error())
		return
	}

	if _, err := handler.User.Register(req.Username, req.Password); err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", nil)
}

// UserLogin user login
// @Summary user login
// @Author l.jiang.1024@gmail.com
// @Description user login
// @Tags User
// @Router /api/v1/sign-in [post]
func UserLogin(c *gin.Context) {
	req := &request.User{}
	if err := c.BindJSON(req); err != nil {
		Fail(c, err.Error())
		return
	}

	respData, err := handler.User.Login(req.Username, req.Password)
	if err != nil {
		Fail(c, err.Error())
		return
	}

	Success(c, "", respData)
}
