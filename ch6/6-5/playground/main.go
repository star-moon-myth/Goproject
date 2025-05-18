package main

import "fmt"

type m struct {
	z []uint64
}
func main() {
	n := m{
		z: []uint64{0,43743,3},
	}
	fmt.Println(n)
	fmt.Printf("%x\n", n)
}