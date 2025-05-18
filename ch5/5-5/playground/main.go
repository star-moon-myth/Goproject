package main

import (
	"fmt"
	"strings"
)

func main() {
	s := product	
	fmt.Println(s(3,3))
	b:=strings.Map(add,"VMS")
	fmt.Println(b)
}
// func square(n int) int     { return n * n }
// func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
func add(a rune)rune{return a+1}
