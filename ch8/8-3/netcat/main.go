package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("TCP","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	defer conn.Close()
	go mustcopy(os.Stdout,conn)
	mustcopy(conn,os.Stdin)//发送给服务器进行处理，然后再进行输出
}
func mustcopy(m io.Writer,n io.Reader){
	_,err:=io.Copy(m,n)
	if err!=nil{
		log.Fatal()
	}
}
