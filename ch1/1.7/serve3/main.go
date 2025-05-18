// handler echoes the entire HTTP request information back to the client.
package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))

}
func handler(w http.ResponseWriter, r *http.Request) {
    // 打印请求的方法 (GET, POST, etc.), URL 以及 HTTP 协议版本 (HTTP/1.1, etc.)
    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

    // 遍历请求头 (Headers) 并打印每个头部字段的键值对
    for k, v := range r.Header {
        // k 是头部字段的名称 (例如 "Content-Type", "User-Agent")
        // v 是一个字符串切片，因为一个头部字段可以有多个值 (虽然通常只有一个)
        fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
    }

    // 打印请求的 Host 头部字段，它指示请求的目标主机名
    fmt.Fprintf(w, "Host = %q\n", r.Host)

    // 打印客户端的远程地址 (IP 地址和端口)，即发送请求的客户端的地址
    fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

    // 解析请求的表单数据。
    // 对于 POST 请求，如果 Content-Type 是 'application/x-www-form-urlencoded' 或 'multipart/form-data'，
    // 或者对于 GET 请求，URL 中包含了查询参数，
    // r.ParseForm() 会将这些数据解析到 r.Form 字段中。
    if err := r.ParseForm(); err != nil {
        // 如果解析表单数据出错，例如请求格式不正确，则记录错误日志
        log.Print(err)
    }

    // 遍历解析后的表单数据 (Form data) 并打印每个表单字段的键值对。
    // Form 数据包含了 URL 查询参数以及 POST 请求体中的表单数据。
    for k, v := range r.Form {
        // k 是表单字段的名称
        // v 是一个字符串切片，因为一个表单字段可以有多个值 (例如复选框)
        fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
    }
}
