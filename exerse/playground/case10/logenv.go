package main

import (
	"fmt"
	"os"
	"strconv"
)

type Appconfigue struct {
	Serveport int
	Loglevel  string
}

func Configue() Appconfigue {
	port := os.Getenv("PORT")
	if port==""{
		port="8080"
		fmt.Println("configur is not set,use default port 8080")
	}
	nport,err:=strconv.Atoi(port)
	if err!=nil{
		fmt.Printf("use default port 8080,err:%v",err)
		nport=8080
	}

	log:=os.Getenv("Loglevel")
	if log==""{
		log="info"
		fmt.Println("loglevel is not set,use default level info")

	}
	return Appconfigue{Serveport: nport,Loglevel: log}
}
func main(){
	configue:=Configue()
	fmt.Printf("The configue is ready,port:%d,loglevel:%s",configue.Serveport,configue.Loglevel)
}