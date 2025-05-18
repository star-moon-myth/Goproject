// gopl.io/ch8/reverb1/reverb1.go  和 gopl.io/ch8/reverb2/reverb2.go  (共用 echo 函数)
package main

import (
	"bufio"       // 导入 bufio 包，提供带缓冲的 I/O
	"fmt"         // 导入 fmt 包，提供格式化 I/O
	"log"         // 导入 log 包，用于日志记录
	"net"         // 导入 net 包，提供网络编程接口
	"strings"      // 导入 strings 包，提供字符串操作
	"time"        // 导入 time 包，提供时间相关功能
)

// echo 函数：模拟带延迟的回声，先大写后小写
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout)) // 发送大写的回声
	time.Sleep(delay)                             // 延迟一段时间
	fmt.Fprintln(c, "\t", strings.ToLower(shout)) // 发送小写的回声
	time.Sleep(delay)                             // 再次延迟
}


// handleConn 函数 for reverb1 (version 2) - 顺序处理
func handleConnReverb1(c net.Conn) {
	input := bufio.NewScanner(c) // 创建 Scanner 读取连接中的数据
	for input.Scan() {         // 循环读取每一行输入
		echo(c, input.Text(), 1*time.Second) // 顺序调用 echo 函数处理每一行
	}
	c.Close() // 关闭连接
}

// handleConn 函数 for reverb2 - 并发处理


func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // 监听本地 8000 端口 TCP 连接
	if err != nil {
		log.Fatal(err) // 如果监听出错，则记录错误并退出
	}
	for { // 无限循环，持续接受新的连接
		conn, err := listener.Accept() // 接受新的连接
		if err != nil {
			log.Print(err) // 如果接受连接出错，则仅记录错误，不退出
			continue      // 继续下一次循环，尝试接受新的连接
		}
		//  根据需要选择 handleConnReverb1 或 handleConnReverb2
		go handleConnReverb1(conn) // 启动 goroutine 并发处理新连接 (这里默认使用 reverb2 的并发版本)
	}
}

