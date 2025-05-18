package main

import (
	"fmt"

	_ "github.com/gin-gonic/gin"
)

// func main(){
// 	router:=gin.Default()
// 	router.GET("/",func(c *gin.Context) {
// 		c.String(200,"哈喽,Gin!我来了!🤣CN")
// 	})
// 	router.Run()
// }
func main() {
    defer fmt.Println("defer 执行了")
    panic("boom!")
}
