package main

import "fmt"

// func example() (result int) { // 命名返回值 result
// 	defer func() {
// 		result *= 2 // 在 defer 中修改 result 的值
// 	}()

// 	result = 10 // 给 result 赋值为 10
// 	return // 相当于 return result
// }
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func main() {
	// fmt.Println(example()) // 打印结果是多少？
	
	defer fmt.Println(double(4))
}
