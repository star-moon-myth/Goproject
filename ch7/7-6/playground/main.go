package main

import (
	"fmt"
	"sort"
)

// type stringslice []string

// func (s stringslice) Len() int {
// 	return len(s)
// }
// func (s stringslice) Less(i, j int) bool {
// 	return s[i] > s[j]
// }
// func (s stringslice) Swap(i, j int) {
// 	s[i], s[j] = s[j], s[i]
// }


func main() {
	name := []string{"Bob", "Aliece", "Dally", "Cherry"}

	// 需要显式转换 name 的类型
	sort.Strings(name) 

	fmt.Println(name) // 输出: [Aliece Bob Cherry Dally]
}
