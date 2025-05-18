package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	count:=make(map[string]int)
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan(){
		line:=input.Text()
		count[line]++


	}
	for m,n:=range count {
		if n>1 {
			fmt.Println(m,n)
		}
	}
}