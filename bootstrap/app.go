package bootstrap

import (
	"gin-laravel/app/Http/Middleware"
	Config "gin-laravel/config"
	Routers "gin-laravel/routes"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	ENVIRONMENT = "dev" // 全局环境 dev开发环境 pro部署环境
)

//创建框架实例,并赋值全局变量R
var Engine = gin.New()

//日志文件writter
var F *os.File

func Run() {
	//初始化框架配置
	Config.InitConfig(ENVIRONMENT)
	//日志记录中间件
	Engine.Use(Middleware.LoggerToFile())
	//错误异常恢复中间件
	Engine.Use(gin.Recovery())
	//终止前端options请求,直接放回
	Engine.Use(Middleware.Options)

	Routers.InitRouters(Engine)

	//server start
	initServer()
	//注册crontab服务
	//initCron()
}
