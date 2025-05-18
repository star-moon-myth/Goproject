package main
import _"fmt"
func main() {
	s:="china,love"
	// fmt.Println(len(s))
	// fmt.Println(s[9])
	// fmt.Println(s[len(s)-1])
	// for i, r := range "Hello, 世界" {
	// 	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	// }
	n := 0
	for range s {
		n++
	}

}