package api

import (
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/handler"
	"github.com/ismdeep/alchemy-furnace/request"
)

func init() {
	noAuth.POST("/api/v1/sign-up", UserRegister)
	noAuth.POST("/api/v1/sign-in", UserLogin)
	auth.GET("/api/v1/my/profile", UserMyProfile) // get login user profile
}

// UserRegister user register
// @Summary user register
// @Author l.jiang.1024@gmail.com
// @Description user register
// @Tags User
// @Router /api/v1/sign-up [post]
func UserRegister(c *gin.Context) {
	req := &request.User{}
	if err := c.BindJSON(req); err != nil {
		Fail(c, err)
		return
	}

	if _, err := handler.User.Register(req.Username, req.Password); err != nil {
		Fail(c, err)
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
// @Author @uniontech.com
// @Description user profile
// @Tags User
// @Router /api/v1/my/profile [get]
func UserMyProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	userProfile, err := handler.User.GetProfile(userID)
	if err != nil {
		Fail(c, err)
		return
	}

	Success(c, "", userProfile)
}
