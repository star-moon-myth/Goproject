package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping",func(ctx *gin.Context) {
		log.Println("begin to handle long works:")
		time.Sleep(10*time.Second)
		log.Println("long works completed!")
		ctx.String(http.StatusOK,"long works completed!")
	})

	srv:=&http.Server{
		Addr: ":8000",
		Handler: r,
	}

	go func() {
		log.Println("服务器启动,监听端口:")
		if err:=srv.ListenAndServe();err!=nil && err!=http.ErrServerClosed{
			log.Fatalf("listen:%s",err)
		}
	}()

	quit:=make(chan os.Signal,1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit
	log.Println("receive shutdown signal,starting graceful server shutdown:")
	ctx,cancel:=context.WithTimeout(context.Background(),4*time.Second)
	defer cancel()
	if err:=srv.Shutdown(ctx);err!=nil{
		log.Fatal("failed to gracefully shut down server",err)
	}
	log.Println("Server had already shut down!")
}