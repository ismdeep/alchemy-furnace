package api

import "github.com/gin-gonic/gin"

func Fail(c *gin.Context, msg string) {
	resp := make(map[string]interface{})
	resp["code"] = 1
	resp["msg"] = msg
	c.JSON(200, resp)
}

func Success(c *gin.Context, msg string, data interface{}) {
	resp := make(map[string]interface{})
	resp["code"] = 0
	resp["msg"] = msg
	resp["data"] = data
	c.JSON(200, resp)
}
