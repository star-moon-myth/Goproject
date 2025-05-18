package main

import (
	"fmt"
	"syscall"
)

func main() {
	err := syscall.Errno(2) // 创建一个 syscall.Errno 错误，错误代码为 2 (ENOENT - 文件或目录不存在)
	fmt.Println(err)        // 打印错误信息
	fmt.Printf("err 的类型是: %T\n", err) // 打印 err 的类型
	fmt.Println(err.Error()) // 显式调用 Error() 方法打印错误信息

	var sysError error = syscall.Errno(2) //  把 syscall.Errno 赋值给 error 接口变量 (多态!)
	fmt.Println(sysError)
	fmt.Println(sysError.Error())
}
