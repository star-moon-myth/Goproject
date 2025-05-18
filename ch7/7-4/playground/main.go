// tempflag.go
package main

import (
	"flag"
	"fmt"
	"gohello/ch7/7-4/temp" // 假设 tempconv 包的路径，根据你的实际情况修改
)
//和flag.var有异曲同工之妙，对其进行了二次接口包装，方便使用
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
