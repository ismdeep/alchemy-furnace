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
		token := c.Request.Header.Get("Authorization")
		if len(token) > 7 && token[0:7] != "Bearer " {
			c.JSON(200, map[string]interface{}{"code": 403, "msg": "token verification failed"})
			c.Abort()
			return
		}

		if len(token) <= 7 {
			c.JSON(200, map[string]interface{}{"code": 403, "msg": "token verification failed"})
			c.Abort()
			return
		}

		bytes, err := jwt.VerifyToken(token[7:])
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

var eng *gin.Engine
var noAuth *gin.RouterGroup
var auth gin.IRoutes

func init() {
	gin.SetMode(gin.ReleaseMode)
	eng = gin.Default()
	noAuth = eng.Group("")
	auth = eng.Group("").Use(Authorization())
}

func Run() {
	log.Info("main", log.String("info", "started to listening"), log.String("bind", config.Bind))
	if err := eng.Run(config.Bind); err != nil {
		panic(err)
	}
}
