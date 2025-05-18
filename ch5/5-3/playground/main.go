package main

import (
	"fmt"
	"log"
)

func main() {
	// 使用 fmt.Println
	fmt.Println("Hello from fmt.Println!")

	// 使用 log.Println
	log.Println("Hello from log.Println!")

	// 配置 log 包
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Hello with flags!")

}

