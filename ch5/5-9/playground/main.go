package main

import (
	"fmt"
	"os"
	"runtime"
)

// func main() {
// 	defer func(){
// 		if r:=recover();r!=nil{
// 			fmt.Println("捕获到了panic了",r)
// 			fmt.Println("完了,出事了")
// 		}
// 	}()
// 	checkage(10)
// 	checkage(-4)
// 	fmt.Println("全都没问题的啦🌶")
// }
// func checkage(age int){
// 	if age<0{
// 		panic("年龄不能为负")
// 	}
// 	fmt.Println("٩(•̤̀ᵕ•̤́๑)ᵒᵏᵎᵎᵎᵎ了")
// }
func f(i int){
	fmt.Printf("f(%d)\n",i+0/i)
	defer fmt.Printf("panic %d\n",i)
	// f(i-1)
}
func printstack(){
	var s [4096]byte
	n:=runtime.Stack(s[:],false)
	os.Stdout.Write(s[:n])
}
func main(){
	defer printstack()
	f(3)

}