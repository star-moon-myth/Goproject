package main

import (
	"fmt"
	"net/http"
)



func Serveh(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"hello worlld🤣")
}
func main(){
	
	http.HandleFunc("/hello",Serveh)
	http.ListenAndServe("localhost:8080",nil)
}