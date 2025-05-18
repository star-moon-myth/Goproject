package main

import (
	_ "errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)
func main(){

	if err := Waitforsever("http://localhost:5432"); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}
func Waitforsever(url string) error{
	const timeout=time.Minute
	deadline:=time.Now().Add(timeout)
	for try:=0;time.Now().Before(deadline);try++{
		_,err:=http.Head(url)
		if err==nil{
			return nil
		}
		log.Printf("server not respending%s",err)
		time.Sleep(time.Second << uint(try))
	}
	return fmt.Errorf("serve %s faild to respond after %s",url,timeout)
}