package main

import (
	"fmt"
	links "gohello/ch8/8-6/playground"
	"log"
	"os"
)
var tokens=make(chan struct{},20)
func crawl(url string) []string {
	fmt.Println(url)
	tokens<-struct{}{}
	list,err:=links.Extract(url)
	<-tokens
	if err!=nil{
		log.Print(err)
	}
	return list
}
func main(){
	worklist:=make(chan []string)
	unseenlinks:=make(chan string)


	go func(){
		worklist<-os.Args[1:]
	}()
	for i:=0;i<20;i++{
		go func(){
			for link:=range unseenlinks{
				find:=crawl(link)
				go func ()  {
					worklist<-find
				}()
			}
		}()
	}
	unseenurl:=make(map[string]bool)
	for list:=range worklist{
		for _,link:=range list{
			if !unseenurl[link]{
				unseenurl[link]=true
				unseenlinks<-link
			}
		} 
	}
}