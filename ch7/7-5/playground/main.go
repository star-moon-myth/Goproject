package main

import (
	"fmt"
	"time"
)

func main() {
	var x interface{} = time.Now()
	fmt.Println(x) // 输出当前时间

	// 类型断言
	if t, ok := x.(time.Time); ok {
		fmt.Println("Time:", t)
	} else {
		fmt.Println("x is not time.Time")
	}
}
