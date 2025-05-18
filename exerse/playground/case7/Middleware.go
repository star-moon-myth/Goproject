package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hellophandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"hello,haha😄")
}
func loghandler(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start:=time.Now()
		fmt.Printf("start:➡️%s,%s\n",r.Method,r.URL.Path)
		next.ServeHTTP(w,r)
		fmt.Printf("completed:👈%s,%s\n",r.URL.Path,time.Since(start))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey != "secret123" { // 假设密钥是 secret123
			http.Error(w, "未授权", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r) // 认证通过，继续
	})
}

func main(){
	hh:=http.HandlerFunc(hellophandler)
	last:=loghandler(hh)
	nlast:=authMiddleware(last)
	http.Handle("/hello",nlast)
	fmt.Println("serve starting on 8000")
	log.Fatal(http.ListenAndServe(":8080",nil))
}