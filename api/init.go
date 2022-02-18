package api

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/alchemy-furnace/model"
	"github.com/ismdeep/alchemy-furnace/schema"
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/log"
	"github.com/ismdeep/parser"
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

func PermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		nodeID, _ := parser.ToUint(c.Param("node_id"))
		taskID, _ := parser.ToUint(c.Param("task_id"))

		var node *model.Node
		var task *model.Task

		if nodeID != 0 {
			var nodes []model.Node
			if err := model.DB.Where("id=?", nodeID).Find(&nodes).Error; err != nil {
				Fail(c, err)
				c.Abort()
				return
			}
			if len(nodes) <= 0 {
				Fail(c, errors.New("node not found"))
				c.Abort()
				return
			}
			node = &nodes[0]
		}

		if taskID != 0 {
			var tasks []model.Task
			if err := model.DB.Where("id=?", taskID).Find(&tasks).Error; err != nil {
				Fail(c, err)
				c.Abort()
				return
			}
			if len(tasks) <= 0 {
				Fail(c, errors.New("task not found"))
				c.Abort()
				return
			}
			task = &tasks[0]
		}

		if node != nil && node.UserID != userID {
			Fail(c, errors.New("permission denied"))
			c.Abort()
			return
		}

		if task != nil && task.UserID != userID {
			Fail(c, errors.New("permission denied"))
			c.Abort()
			return
		}
	}
}

var eng *gin.Engine
var noAuth *gin.RouterGroup
var auth gin.IRoutes
var permCheckAuth gin.IRoutes

func init() {
	gin.SetMode(gin.ReleaseMode)
	eng = gin.Default()
	noAuth = eng.Group("")
	auth = eng.Group("").Use(Authorization())
	permCheckAuth = auth.Use(PermissionCheck())
}

func Run() {
	log.Info("main", log.String("info", "started to listening"), log.String("bind", config.Bind))
	if err := eng.Run(config.Bind); err != nil {
		panic(err)
	}
}
