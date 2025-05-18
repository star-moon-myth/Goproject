package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for{
		t,err:=dec.Token()
		if err==io.EOF{
			break
		}else if err!=nil {
			fmt.Fprintf(os.Stderr,"err:%v",err)
			os.Exit(1)
		}
		switch tok:=t.(type){
		case xml.StartElement:
			stack=append(stack, tok.Name.Local)
		case xml.EndElement:
			stack=stack[:len(stack)-1]
		case xml.CharData:
			if containall(stack,os.Args[1:]){
				fmt.Printf("%s:%s",strings.Join(stack," "),string(tok))
			}

		}
	}
}
func containall(m,n []string)bool{
	for len(n)<len(m){
		if len(n)==0{
			return true 
		}
		if n[0]==m[0]{
			m=m[1:]
		}
		n=n[1:]
	}
	return false
}