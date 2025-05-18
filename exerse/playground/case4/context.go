package main

import (
	"context"
	"fmt"
	"time"
)
func longoperation(ctx context.Context,ch chan<- string){
	fmt.Println("Are you ready? work is going on💪")
	select{
	case <-time.After(2*time.Millisecond):
		ch<-"congratulation!🎉"
	case <-ctx.Done():
		ch <-fmt.Sprintf("😭there is an error:%v",ctx.Err())

	}
}
func main() {
	ctx, cancle := context.WithTimeout(context.Background(),1*time.Millisecond)
	defer cancle()
	result:=make(chan string,1)
	go longoperation(ctx,result)
	
	res:=<-result
	fmt.Println(res)
}