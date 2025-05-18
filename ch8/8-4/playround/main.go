package main
import "fmt"
func main(){
	ch:=make(chan int)
	go func(){
		ch<-10
		fmt.Println("发送了数据10")
	}()
	x:=<-ch
	fmt.Println(x)
}