package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("TCP","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			log.Print(err)
			continue 
		}
		go handleConn(conn)
	}
}
func handleConn(m net.Conn){
	defer m.Close()
	for{
		_,err:=io.WriteString(m,time.Now().Format("15:02:55\n"))
		if err!=nil{
			return 
		}

		time.Sleep(1*time.Second)
	}
}