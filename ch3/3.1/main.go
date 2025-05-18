package main
import "fmt"
func main(){
	o:=06666
	fmt.Printf("%d %[1]o %#[1]X\n",o)
	x:=int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}
