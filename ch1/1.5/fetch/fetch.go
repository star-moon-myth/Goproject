package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
func main() {
	for _,url:=range os.Args[1:] {
		res,err:=http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch%v\n",err)
			os.Exit(1)
		}
		content,err:=io.ReadAll(res.Body)
		res.Body.Close()
		if err !=nil {
			fmt.Fprintf(os.Stderr,"fetch:reading%s,%v\n",url,content)
			os.Exit(1)
		}
		fmt.Printf("%s",content)

	}
}