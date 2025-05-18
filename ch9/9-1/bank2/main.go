package main

import (
	"fmt"
	"sync"
)

var counter int

func increasement(wg *sync.WaitGroup){
	defer wg.Done()
	for i:=0;i<1000;i++{
		counter++
	}
}
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	for i:=0;i<2;i++{
		go increasement(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}