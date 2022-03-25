package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
	"github.com/ismdeep/alchemy-furnace/response"
)

// UserLogin user login
// @Summary user login
// @Author l.jiang.1024@gmail.com
// @Description user login
// @Tags User
// @Router /api/v1/sign-in [post]
func UserLogin(c *gin.Context) {
	var req request.User
	if err := c.BindJSON(&req); err != nil {
		Fail(c, err)
		return
	}

	respData, err := handler.User.Login(req.Username, req.Password)
	if err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", respData)
}

// UserMyProfile user profile
// @Summary user profile
// @Author l.jiang.1024@gmail.com
// @Description user profile
// @Tags User
// @Router /api/v1/my/profile [get]
func UserMyProfile(c *gin.Context) {
	Success(c, "", response.UserProfile{
		Username: config.ROOT.Auth.Username,
	})
}
