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
		log.Fatal()
	}
	defer conn.Close()
	mustcopy(os.Stdout,conn)
}
func mustcopy(m io.Writer,n io.Reader){
	_,err:=io.Copy(m,n)
	if err!=nil{
		log.Fatal()
	}
}