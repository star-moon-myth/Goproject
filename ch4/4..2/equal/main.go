package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5,5}
	b := []int{1,2,3,4,5,6}
	fmt.Println(eqaul(s,b))
}
func eqaul(x,y []int)bool{
	if len(x)!=len(y){
		return false
	}
	for i,j:=range x {
		if x[i]!=y[i]{
			fmt.Printf("索引：%d，错误匹配值：%d,%d\n",i,j,y[i])
			return false
		}
	}
	return true
}