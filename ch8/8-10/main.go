package main

import (
	"bufio"
	
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client) //  ✨  声明并初始化 ente
	leaving  = make(chan client)  //  ✨  声明并初始化 leaving channel
	messages = make(chan string)  //  ✨  声明并初始化 messages channel
)


func main() {
	listen, err := net.Listen("tcp","localhost:8000")
	if err!=nil{
		log.Fatal(err)
	}
	go broadcast()
	for{
		conn,err:=listen.Accept()
		if err!=nil{
			log.Print(err)
			continue
		}
		go handle(conn)
	}
}
func broadcast(){
	clients:=make(map[client]bool)
	for{
		select{
		case msg:=<-messages:
			for cli:=range clients{
				cli <- msg 
			}
		case cli:=<-entering:
			clients[cli]=true
		case cli:=<-leaving:
			delete(clients,cli)
			close(cli)
		}
	}
}

func handle(conn net.Conn){
	newclient:=make(chan string)
	// go clientwriter(conn,newclient)
	who:=conn.RemoteAddr().String()
	newclient<-"you are:"+who
	messages<-who+"have arrived"
	entering<-newclient
	input:=bufio.NewScanner(conn)
	for input.Scan(){
		messages<-who+input.Text()
	}
	leaving<-newclient
	messages<-who+"have leaved"
	conn.Close()

}

// func clientwriter(conn net.Conn,cli <-chan string){
// 	for msg:=range cli{
// 		fmt.Fprintln(conn,msg)
// 	}
// }ing){
// 	for msg:=range cli{
// 		fmt.Fprintln(conn,msg)
// 	}
// }ing){
// 	for msg:=range cli{
// 		fmt.Fprintln(conn,msg)
// 	}
// }ng){
// 	for msg:=range cli{
// 		fmt.Fprintln(conn,msg)
// 	}
// }