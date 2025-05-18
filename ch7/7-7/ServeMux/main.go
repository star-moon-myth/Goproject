package main

import (
	"fmt"
	"log"
	"net/http"
)

func listhandler(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w,"商品列表👖，👟，👔")
}
func procehandler(w http.ResponseWriter,r *http.Request){
	price:="未知"
	item:=r.URL.Query().Get("item")
	if item == "shoes" {
		price = "$50"
	} else if item == "socks" {
		price = "$5"
	}
	fmt.Fprintf(w, "%s 的价格是: %s", item, price)

}
func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/list",listhandler)
	mux.HandleFunc("/price",procehandler)
	log.Fatal(http.ListenAndServe("localhost:8080",mux))
}