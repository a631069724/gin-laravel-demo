package Routers

import (
	"gin-laravel/app/Http/Controllers/v1"
	"gin-laravel/app/Http/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters(Route *gin.Engine) {

	//404处理
	Route.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, http.StatusText(404))
	})

	//路由分组
	r1 := Route.Group("api/v1").Use(Middleware.AUthMiddleware)
	{
		//路由设置示例
		r1.GET("/c", v1.C)
		r1.GET("/r", v1.R)
		r1.GET("/u", v1.U)
		r1.GET("/d", v1.D)
	}

}
