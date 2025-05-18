// Server2 是一个最小的 "echo" 和计数器服务器。
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync" // 导入 sync 包以使用互斥锁，保证计数器的并发安全
)

var mu sync.Mutex // mu 互斥锁保护 count 变量，防止竞态条件
var count int     // count 记录 handler 函数被调用的次数

func main() {
	// 注册处理器函数 handler 来处理根路径 "/" 的请求。
	http.HandleFunc("/", handler)
	// 注册处理器函数 counter 来处理路径 "/count" 的请求。
	http.HandleFunc("/count", counter)
	// 启动 HTTP 服务器并监听 localhost:8000 端口。
	// 如果启动失败，log.Fatal 会记录错误并退出程序。
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler 函数处理对根路径 "/" 的请求。
// 它会回显请求 URL 的 Path 组件，并增加计数器。
func handler(w http.ResponseWriter, r *http.Request) {
	// 使用互斥锁 mu 来保护 count 变量的并发访问。
	mu.Lock()      // 获取锁，阻止其他 goroutine 同时访问 count
	count++        // 增加请求计数器
	mu.Unlock()    // 释放锁，允许其他 goroutine 访问 count
	// 使用 fmt.Fprintf 将格式化的字符串写入到 http.ResponseWriter 中。
	// r.URL.Path 获取请求 URL 的 Path 组件。
	// %q 用于将字符串用双引号括起来，并转义特殊字符。
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter 函数处理对路径 "/count" 的请求。
// 它会回显 handler 函数被调用的总次数。
func counter(w http.ResponseWriter, r *http.Request) {
	// 使用互斥锁 mu 来保护 count 变量的并发访问。
	mu.Lock()      // 获取锁，保证读取 count 值的原子性
	fmt.Fprintf(w, "Count %d\n", count) // 将当前的 count 值写入到响应中
	mu.Unlock()    // 释放锁
}
