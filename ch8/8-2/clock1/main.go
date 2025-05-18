package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000") // 监听 TCP 连接，地址是 localhost:8000
	if err != nil {
		log.Fatal(err) // 如果监听失败，直接报错退出
	}
	for { // 无限循环，一直监听新的连接
		conn, err := listener.Accept() // 接受客户端的连接，程序会在这里阻塞，直到有新的连接到来
		if err != nil {
			log.Print(err) // e.g., connection aborted  如果接受连接出错，打印错误信息，继续循环
			continue
		}
		handleConn(conn) // 处理一个客户端连接，一次处理一个
	}
}

func handleConn(c net.Conn) {
	defer c.Close() // 函数退出前关闭连接，避免资源泄露
	for { // 无限循环，不停地向客户端发送时间
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) // 获取当前时间并格式化，发送给客户端
		if err != nil {
			return // e.g., client disconnected  如果发送失败，说明客户端可能断开了连接，结束函数
		}
		time.Sleep(1 * time.Second) // 等待 1 秒钟，再发送下一次时间
	}
}
