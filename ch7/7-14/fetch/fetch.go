package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args)!=2{
		fmt.Fprintf(os.Stderr,"err\n")
		os.Exit(1)
	}
	url:=os.Args[1]
	resp,err:=http.Get(url)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"err%v\n",err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode!=http.StatusOK{
		fmt.Fprintf(os.Stderr,"err%v\n",resp.Status)
		os.Exit(1)
	}
	_,err=io.Copy(os.Stdout,resp.Body)
	if err!=nil{
		fmt.Fprintf(os.Stderr,"url%s:%v\n",url,err)
		os.Exit(1)
	}


}