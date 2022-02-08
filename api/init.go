package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/alchemy-furnace/schema"
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/log"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		bytes, err := jwt.VerifyToken(token)
		if err != nil {
			c.JSON(200, map[string]interface{}{"code": 403, "msg": "token verification failed"})
			c.Abort()
			return
		}

		u := schema.LoginUser{}
		if err := json.Unmarshal([]byte(bytes), &u); err != nil {
			c.JSON(200, map[string]interface{}{"code": 403, "msg": "token verification failed"})
			c.Abort()
			return
		}

		c.Set("user_id", u.ID)
		c.Set("username", u.Username)
		c.Next()
	}
}

func init() {
	gin.SetMode(gin.ReleaseMode)
	eng := gin.Default()
	eng.POST("/api/v1/sign-up", UserRegister)
	auth := eng.Group("/api/v1")
	auth.Use(Authorization())
	auth.GET("/api/v1/tasks", TaskList)
	auth.POST("/api/v1/tasks", TaskCreate)
	auth.GET("/api/v1/tasks/:task_id/runs", RunList)
	auth.GET("/api/v1/tasks/:task_id/runs/:run_id", RunDetail)
	log.Info("main", log.String("info", "started to listening"), log.String("bind", config.Config.Bind))
	go func() {
		if err := eng.Run(config.Config.Bind); err != nil {
			panic(err)
		}
	}()
}
