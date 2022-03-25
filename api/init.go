package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ismdeep/alchemy-furnace/config"
	"github.com/ismdeep/alchemy-furnace/schema"
	"github.com/ismdeep/jwt"
	"github.com/ismdeep/log"
	"github.com/ismdeep/parser"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		nodeID, _ := parser.ToUint(c.Param("node_id"))
		tokenID, _ := parser.ToUint(c.Param("token_id"))
		taskID, _ := parser.ToUint(c.Param("task_id"))
		triggerID, _ := parser.ToUint(c.Param("trigger_id"))
		runID, _ := parser.ToUint(c.Param("run_id"))
		c.Set("node_id", nodeID)
		c.Set("token_id", tokenID)
		c.Set("task_id", taskID)
		c.Set("trigger_id", triggerID)
		c.Set("run_id", runID)

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

		c.Next()
	}
}

var eng *gin.Engine
var free *gin.RouterGroup
var auth gin.IRoutes

func init() {
	gin.SetMode(gin.ReleaseMode)
	eng = gin.Default()
	free = eng.Group("")
	free.POST("/api/v1/sign-in", UserLogin)
	free.GET("/api/v1/tasks/:task_id/runs/:run_id/log", RunLog) // Get run log with websocket

	auth = eng.Group("").Use(Authorization())
	auth.GET("/api/v1/nodes", NodeList)
	auth.POST("/api/v1/nodes", NodeAdd)
	auth.PUT("/api/v1/nodes/:node_id", NodeUpdate)
	auth.DELETE("/api/v1/nodes/:node_id", NodeDelete)
	auth.POST("/api/v1/tasks/:task_id/triggers/:trigger_id/runs", RunCreate) // Start to run a task by trigger
	auth.GET("/api/v1/tasks/:task_id/runs", RunList)
	auth.GET("/api/v1/tasks/:task_id/runs/:run_id", RunDetail)
	auth.GET("/api/v1/tasks", TaskList)
	auth.POST("/api/v1/tasks", TaskCreate)
	auth.PUT("/api/v1/tasks/:task_id", TaskUpdate)
	auth.GET("/api/v1/tasks/:task_id", TaskDetail)
	auth.DELETE("/api/v1/tasks/:task_id", TaskDelete)
	auth.GET("/api/v1/tokens", TokenList)
	auth.POST("/api/v1/tokens", TokenAdd)
	auth.PUT("/api/v1/tokens/:token_id", TokenUpdate)
	auth.DELETE("/api/v1/tokens/:token_id", TokenDelete)
	auth.GET("/api/v1/tasks/:task_id/triggers", TriggerList)
	auth.POST("/api/v1/tasks/:task_id/triggers", TriggerAdd)
	auth.PUT("/api/v1/tasks/:task_id/triggers/:trigger_id", TriggerUpdate)
	auth.DELETE("/api/v1/tasks/:task_id/triggers/:trigger_id", TriggerDelete)
	auth.GET("/api/v1/my/profile", UserMyProfile) // get login user profile
}

func Run() {
	log.Info("main", log.String("info", "started to listening"), log.String("bind", config.ROOT.Bind))
	if err := eng.Run(config.ROOT.Bind); err != nil {
		panic(err)
	}
}
