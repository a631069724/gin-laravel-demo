package v1

import (
	"github.com/gin-gonic/gin"
)

func C(c *gin.Context) {
	c.JSON(200,"插入成功")
	return
}

func R(c *gin.Context) {

	c.JSON(200,"r")
	return
}

func U(c *gin.Context)  {

	c.JSON(200,"u")
	return
}

func D(c *gin.Context)  {

	c.JSON(200,"删除成功")
	return
}
