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
var router = gin.New()

//日志文件writter
var F *os.File

func Run() {
	//初始化框架配置
	Config.InitConfig(ENVIRONMENT)

	//日志记录
	router.Use(Middleware.LoggerToFile())

	//注册自定义验证
	//initValidator()

	//Database.InitDb(F)
	//Database.InitRedis()
	Routers.InitRouters(router)

	//server start
	initServer()
	//注册crontab服务
	//initCron()
}
