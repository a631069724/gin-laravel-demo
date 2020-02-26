package Middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//全局中间件 允许跨域
func AUthMiddleware(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	pass := c.DefaultQuery("pass", "")

	fmt.Printf("login %s %s \r\n", name, pass)
	if name == "123" && pass == "321" {
		c.Next()

	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "登录失效",
		})
		c.Abort()
	}

}
