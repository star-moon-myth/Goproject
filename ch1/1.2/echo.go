package main

import (
	"fmt"
	"os"
	// "strings"
)
func main() {
	// var s,c string
	// for i:=1;i<len(os.Args);i++{
	// 	s+=c+os.Args[i]

	// }
	// fmt.Println(s)
	// fmt.Println(strings.Join(os.Args[1:]," "))
	// fmt.Println(os.Args[1:])
	
	for a,i :=range(os.Args[:]){
		
		fmt.Printf("索引%d,值%s\n",a,i)

		
	}
	
	// fmt.Println(strings.Join(os.Args," "))
}