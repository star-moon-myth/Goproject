package main
import "fmt"
func add(x,y int)int{
	return x+y
}
func sub(x,y int)(z int){
	z=x+y
	return 
}
func first(x int,_ int)int{
	return x
}
func zero(int,int)int{
	return 0
}
func main(){
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"
}

