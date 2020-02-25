package main

import (
	"context"
	"fmt"
	Config "gin-laravel/config"
	Routers "gin-laravel/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	ENVIRONMENT = "dev" // 全局环境 dev开发环境 pro部署环境
)

//创建框架实例,并赋值全局变量R
var router = gin.New()

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
	Routers.InitRouters(router)

	//server start
	initServer()
	//注册crontab服务
	//initCron()
}

func initServer() {
	fmt.Printf("运行gin 模式%s http://%s \r\n", Config.App.GetString("server.GinMode"), Config.App.GetString("server.Address"))

	//设置gin模式
	gin.SetMode(Config.App.GetString("server.GinMode"))

	srv := &http.Server{
		Addr:    Config.App.GetString("server.Address"),
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
