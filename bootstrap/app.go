package main

import (
	"fmt"
	"gin-laravel/bootstrap/Server"
	Config "gin-laravel/config"
	Routers "gin-laravel/routes"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var (
	ENVIRONMENT = "dev" // 全局环境 dev开发环境 pro部署环境
)

//创建框架实例,并赋值全局变量R
var R = gin.New()

//日志文件writter
var F *os.File

func main() {
	//初始化框架配置
	Config.InitConfig(ENVIRONMENT)
	//initLog()
	//注册自定义验证
	//initValidator()
	fmt.Printf("Env %s \r\n", Config.App.GetString("server.Env"))

	//Database.InitDb(F)
	//Database.InitRedis()
	Routers.InitRouters(R)

	//server start
	initServer()
	//注册crontab服务
	//initCron()
}

func initServer() {
	//设置gin模式
	gin.SetMode("debug")

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: R,
	}

	Server.ListenAndServer(server)
	//_ = R.Run(Config.Server)/**/
}
