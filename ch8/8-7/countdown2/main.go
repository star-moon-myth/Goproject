package main

import (
	"fmt"
	"os"
	"time"

)

var abort = make(chan struct{})  

func main() {
	go func ()  {
		os.Stdin.Read(make([]byte,1 ))
		close(abort)
	}()
	fmt.Println("即将发射")
	tick:=time.Tick(time.Second)
	for i:=10;i>0;i--{
		fmt.Println(i)
		select{
		case<-tick:

		case<-abort:
			fmt.Println("好吧，停止了哦")
			return
		}
	}
}