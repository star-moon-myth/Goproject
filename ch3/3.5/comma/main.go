package main

import "fmt"

// func comma(s string) string {
// 	if len(s) <= 3 {
// 		return s
// 	}
// 	return comma(s[:len(s)-3]) + "," + s[len(s)-3:]
// }
func main() {
	s := "1233"
	// fmt.Println(comma(s))
	b:=[]byte(s)
	fmt.Println(b)
	fmt.Println(s)
}