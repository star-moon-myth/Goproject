package main

import (
	"fmt"
)

type little struct{}
type big struct{}

func grandparent() {
	fmt.Println("我要调用parent函数了")
	parent()
	fmt.Println("函数执行完了")
}
func parent(){
	defer func(){
		if e:=recover();e!=nil{
			fmt.Println("孩子别怕,我来处理")
		}
	}()
	fmt.Println("我要调用child函数了")
	child()
	fmt.Println("child函数执行完毕")
}
func child(){
	defer func(){
		if p:=recover();p!=nil{
			switch p.(type){
			case little:
				fmt.Println("我自己能解决啦!")
			default:
				fmt.Println("出大事了,快来救我啊")
				panic(p)
			}
		}
	}()
		s:="b"
		switch s{
		case "b":
			panic(big{})
		case "l":
			panic(little{})

		}
}
func main(){
	grandparent()
}