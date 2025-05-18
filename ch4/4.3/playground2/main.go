package main

import "fmt"

func eaual(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for i, j := range x {
		if value, ok := y[i]; !ok || value != j {
			return false
		}
	}
	return true
}
func main() {
	ages := map[string]int{
		"mike": 23,
		"jina": 34,
		"lina": 27,
	}
	bges := map[string]int{
		"mike": 23,
		"jina": 34,
		"lina": 27,
	}
	b := eaual(ages, bges)
	fmt.Println(b)
}