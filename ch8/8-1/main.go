package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n =45
	fib:=fibi(n)
	fmt.Printf("\r%d",fib)
}
func spinner(delay time.Duration){
	for{
		for _,r:=range  `-\|/`{
			fmt.Printf("\r%c",r )
			time.Sleep(delay)
		} 
	}
}
func fibi(x int)int{
	if x<2{
		return x
	}
	return fibi(x-1)+fibi(x-2)
}