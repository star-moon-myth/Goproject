package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"strings"
	"sync"
)

// 1️⃣ **初始化全局变量**：
//    - `urlMap` 存储短链接映射
//    - `mu sync.RWMutex` 读写锁保证并发安全
var urlMap = make(map[string]string)
const urlLength = 6
var mu sync.RWMutex
const shortCodeChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
// 2️⃣ **生成短码函数**：
//    - 使用 `rand.IntN()` 生成 `6` 位随机字符串

func generateCode()string{
	var sb strings.Builder
	for i:=0;i<=urlLength;i++{
		sb.WriteByte(shortCodeChars[rand.IntN(len(shortCodeChars))])
	}
	return sb.String()
}
// 3️⃣ **`shortenHandler` 处理短链接创建** (`POST /shorten`)：
//    - 获取 `url` 参数
//    - 生成**唯一短码**（循环检查是否重复）
//    - 存入 `urlMap`
//    - 返回短链接
func shortenHandler(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodPost{
		http.Error(w,"method not found",http.StatusMethodNotAllowed)
		return 
	}
	originalurl:=r.FormValue("url")
	if originalurl==""{
		http.Error(w,"url not found",http.StatusBadRequest)
	}
	var shortCode string
	mu.Lock()
	for{
		shortCode=generateCode()
		if _,exist:=urlMap[shortCode];!exist{
			urlMap[shortCode]=originalurl
			break
		}
	}
	mu.Unlock()
	fullcode:=fmt.Sprintf("http://localhost:8000/%s",shortCode)
	fmt.Fprintf(w,"shorturl:%s",fullcode)
	fmt.Printf("🔗 Shortened '%s' -> %s\n", originalurl, fullcode)
}
// 4️⃣ **`redirectHandler` 处理短链接重定向** (`GET /{shortCode}`)：
//    - 读取 `urlMap` 获取原始 URL
//    - `http.Redirect()` 302 重定向
func redirectHandler(w http.ResponseWriter,r *http.Request){
	if r.Method!=http.MethodGet{
		http.Error(w,"url not found",http.StatusBadRequest)
		return
	}
	short:=strings.TrimPrefix(r.URL.Path,"/")
	mu.RLock()
	originalurl,exist:=urlMap[short]
	mu.RUnlock()
	if !exist{
		http.NotFound(w, r)
		return
	}
	http.Redirect(w,r,originalurl,http.StatusFound)
	fmt.Printf("↩️ Redirecting '%s' to '%s'\n", short, originalurl)

}
// 5️⃣ **`main()` 启动 HTTP 服务器** (`:8080`)：
//    - `http.HandleFunc()` 注册路由
//    - `http.ListenAndServe()` 监听端口
func main(){
	http.HandleFunc("/shorten",shortenHandler)
	http.HandleFunc("/",redirectHandler)
	err:=http.ListenAndServe(":8000",nil)
	if err!=nil{
		fmt.Printf("💥 Server failed to start:%v\n",err)
	}

}