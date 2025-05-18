package main

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
	mapping=make(map[string]string )
)
func lookup(key string)string{
	mu.Lock()
	n:=mapping[key]
	mu.Unlock()
	return n
}
var cache=struct{
	sync.Mutex
	mapping map[string]string 
}{
	mapping: make(map[string]string),
}
func look(key string)string{
	cache.Lock()
	n:=cache.mapping[key]
	cache.Unlock()
	return n
}
func main(){
	mu.Lock()
	mapping["haha"]="hahaaaa"
	mapping["nina"]="ninnnnn"
	mu.Unlock()
	fmt.Println(lookup("haha"))
	fmt.Println(lookup("nina"))
	cache.Lock()
	cache.mapping["haha"]="hahaaaa"
	cache.mapping["nina"]="ninnnnn"
	cache.Unlock()
	fmt.Println(look("haha"))
	fmt.Println(look("nina"))
	fmt.Println(look("java"))
}
