package main

import (
	"flag"
	"fmt"
	"strings"
)

// 定义一个存储多个字符串的切片类型
type StringSlice []string

// 实现 flag.Value 接口
func (s *StringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *StringSlice) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

func main() {
	var names StringSlice

	// 注册自定义类型
	flag.Var(&names, "names", "输入多个名字，用逗号分隔")

	// 解析命令行参数
	flag.Parse()

	fmt.Println("Parsed Names:", names)
}
