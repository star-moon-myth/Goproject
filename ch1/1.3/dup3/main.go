package main

import (
	"fmt"
	"os"
	"strings"
)
func main() {
	counts:=make(map[string]int)
	for _,filename:=range os.Args[1:] {
		data,err:=os.ReadFile(filename)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"err%v",err)
			continue

		}
		for _,line :=range strings.Split(string(data), "\n"){
			counts[line]++
		}


	}
	for m,n := range counts {
		if n>1 {
			fmt.Println(n,m)
		}
	}
}