// Echo4 prints its command-line arguments.
package main

import (
    "flag"
    "fmt"
    "strings"
)

var n=flag.Bool("n",false,"omit trailing newline")
var s=flag.String("s"," ","separater")
func main(){
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(),*s))
	if !*n{
		fmt.Println()
	}
}