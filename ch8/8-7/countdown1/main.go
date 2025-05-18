package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始计数：")
	tick:=time.Tick(time.Second)
	for k:=10;k>0;k--{
		fmt.Println(k)
		<-tick
	}
}