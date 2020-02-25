package Server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	listener net.Listener = nil
	graceful =  flag.Bool("graceful", false, "listen on fd open 3")
)

//监听服务器
func ListenAndServer(server *http.Server){
	var err error

	//解析参数
	flag.Parse()

	//设置监听的对象(新建或已存在的socket描述符)
	if *graceful {
		//子进程监听父进程传递的 socket描述符
		log.Println("listening on the existing file descriptor 3")
		//子进程的 0 1 2 是预留给 标准输入 标准输出 错误输出
		//因此传递的socket 描述符应该放在子进程的 3
		f := os.NewFile(3,"")
		listener,err = net.FileListener(f)
		log.Printf( "graceful-reborn  %v %v  %#v \n", f.Fd(), f.Name(), listener)
	}else{
		listener,err = net.Listen("tcp",server.Addr)

	}
	if err != nil{
		fmt.Printf("listener error: %v\n",err)
		log.Fatalf("listener error: %v\n",err)
	}
	go func(){
		err = server.Serve(listener)
		fmt.Printf("server.Serve err: %v\n",err)
		log.Printf("server.Serve err: %v\n",err)
		tcp,_ := listener.(*net.TCPListener)
		fd,_ := tcp.File()
		log.Printf( "first-boot  %v %v %#v \n ", fd.Fd(),fd.Name(), listener)
	}()
	//监听信号
	handleSignal(server)
}
//处理信号
func handleSignal(server *http.Server){
	//把信号 赋值给 通道
	ch := make(chan os.Signal, 1)
	//监听信号
	signal.Notify(ch, syscall.SIGINT,syscall.SIGTERM)
	//阻塞主进程， 不停的监听系统信号
	for{
		//通道 赋值给 sig
		sig := <-ch
		log.Printf("signal receive: %v\n", sig)
		ctx,_ := context.WithTimeout(context.Background(),20*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM: //终止进程执行
			log.Println("shutdown")
			signal.Stop(ch)          //停止通道
			_ = server.Shutdown(ctx) //关闭服务器窗口
			log.Println("graceful shutdown")
			return
		}
	}
}


func Run(srv *http.Server)  {
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
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
