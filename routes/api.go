package Routers

import (
	"gin-laravel/app/Http/Controllers/v1"
	"gin-laravel/app/Http/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(R *gin.Engine) {
	//加载中间件模块
	//终止前端options请求,直接放回
	R.Use(Middleware.Options)
	R.Use(gin.Recovery())
	R.Use(gin.Logger())

	//404处理
	R.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, http.StatusText(404))
	})

	//路由分组
	r1 := R.Group("api/v1")
	{
		//路由设置示例
		r1.GET("/c", v1.C)
		r1.GET("/r", v1.R)
		r1.GET("/u", v1.U)
		r1.GET("/d", v1.D)
	}

	return
}
