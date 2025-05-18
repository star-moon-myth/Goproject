package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		log.Println("start to process request:")
		time.Sleep(2*time.Second)
		fmt.Fprintf(w,"request processed😄😄！")
	})
	Server:=&http.Server{Addr: ":8080",Handler: mux}
	go func ()  {
		log.Println("listen on 8080-----------")
		if err:=Server.ListenAndServe();err!=nil && err!=http.ErrServerClosed{
			log.Fatalf("could ont listen on 8080:%v",err)
		}
	}()

	quit:=make(chan os.Signal,1)
	signal.Notify(quit,syscall.SIGINT,syscall.SIGTERM)
	<-quit 
	log.Println("shutdown signal has received,initiate graceful shuodown")

	ctx,cancel :=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	if err:=Server.Shutdown(ctx);err!=nil{
		log.Fatalf("serve forced to shutdown uncleanly:%v",err)
	}
	log.Println("Serve exiting gracefully.")

}