package main

import (
	"fmt"
	"time"
)

type server struct {
	Addr    string
	Port    int
	Timeout time.Duration
	UserTLS bool 
}
type Option func(*server)

func withAddr(addr string)Option{
	return func(s *server) {
		s.Addr=addr
	}
}
func withPort(port int)Option{
	return func(s *server) {
		s.Port=port
	}
}
func withTimeout(time time.Duration)Option{
	return func(s *server) {
		s.Timeout=time
	}
}
func withTLS()Option{
	return func(s *server) {
		s.UserTLS=true 
	}
}
func NewServe(options ...Option)*server{
	ss:=server{
		Addr: "localhost",
		Port: 8080,
		Timeout: 1*time.Minute,
		UserTLS: false,
	}
	for _,option:=range options{
		option(&ss)
	} 
	return &ss
}

func main(){
	server1:=NewServe(withAddr("haha"),withPort(9000))
	fmt.Printf("sever1:%+v\n",*server1)

	server2:=NewServe(withTLS(),withTimeout(2*time.Microsecond))
	fmt.Printf("server2:%+v\n",*server2)

	server3:=NewServe()
	fmt.Printf("server2:%+v\n",*server3)
}