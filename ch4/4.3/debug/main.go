package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]bool)
	// fmt.Println(seen)
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan(){
		line:=input.Text()
		if !seen[line]{
			seen[line]=true 
			fmt.Println(line)
		}
	}
	if err:=input.Err();err!=nil{
		fmt.Fprintf(os.Stderr,"err:%v\n",err)
		os.Exit(1)
	}


}