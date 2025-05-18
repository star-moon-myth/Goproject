package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	lister, err := net.Listen("TCP","localhost:8000")
	if err!=nil{
		log.Fatal(err)

	}
	for{
		conn,err:=lister.Accept()
		if err!=nil{
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}

}
func handleConn(c net.Conn){
	text:=bufio.NewScanner(c)
	defer c.Close()
	for text.Scan(){
		go echo(c,text.Text(),1*time.Second)
	}
}
func echo(c net.Conn,shout string,delay time.Duration){
	fmt.Fprintln(c,"\t",strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",strings.ToLower(shout))
	time.Sleep(delay)
}