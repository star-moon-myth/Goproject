package main
import "fmt"
// func main() {
// 	x:="hello"
// 	for _,x :=range x{
// 		x:=x-32
// 		fmt.Println(string(x))
// 	}	
// }
func f()int{
	var  input int
	fmt.Println("请输入一个 数字：")
	fmt.Scanln(&input)
	return input
}
func g(e int)int{
	return e
}
func main(){
	if i:=f();i==0{
		
		fmt.Println(i)
	}else if m:=g(i);m==i{
		fmt.Println(m,i)
	}else {
		fmt.Println(m,i)
	}
	
}