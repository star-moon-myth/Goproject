package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	r,m,n:=fetch("https://www.runoob.com/")
	fmt.Println(r,m,n)
}
// 对一个URL里的内容进行提取到一个文件中，文件名叫filename，字节数n，如果出错返回err
func fetch(url string)(filename string,n int64,err error){
	response,err:=http.Get(url)//提取URL,包括里面的响应头,响应体
	if err!=nil{
		return "",0,err
	}
	defer response.Body.Close()
	local:=path.Base(response.Request.URL.Path)
	if local=="/"{
		local="index.html"
	}
	f,err:=os.Create(local)
	if err!=nil{
		return "",0,err
	}
	n,err=io.Copy(f,response.Body)
	if closeerr:=f.Close();err==nil{
		err=closeerr
	}
	return local,n,err
}