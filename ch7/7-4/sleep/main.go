// sleep.go
package main

import (
	"flag"
	"fmt"
	"time"
)
var period=flag.Duration("period",1*time.Second,"sleep time")
func main(){
	flag.Parse()
	fmt.Println(*period)
	time.Sleep(*period)
	fmt.Printf("has done")

}



