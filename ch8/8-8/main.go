package main

import (
	"fmt"
	"time"
)

func main() {
	var tick <-chan time.Time
	if true {
		tick = time.Tick(100 * time.Millisecond)
	}
	time.Sleep(1*time.Second)
	for{
		fmt.Println(<-tick)
	}
}