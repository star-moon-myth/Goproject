package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"pong!",
		})
	})

	router.POST("/users",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"用户已创建(假的啦)",
		})
	})

	router.GET("/hello/:name",func(ctx *gin.Context) {
		name:=ctx.Param("name")
		ctx.String(200,"你好,%s",name)
	})

	router.GET("/welcome",func(ctx *gin.Context) {
		firstname:=ctx.DefaultQuery("firstname","Mike")
		lastname:=ctx.Query("lastname")
		ctx.String(200,"welcome you %s %s",firstname,lastname)
	})

	v1:=router.Group("/api/v1")
	{
		v1.GET("/users",func(ctx *gin.Context) {
			ctx.JSON(200,gin.H{"users":[]string{"Alice","Bob"}})
		})
		v1.POST("/products",func(ctx *gin.Context) {
			ctx.JSON(200,gin.H{"status":"产品已经添加"})
		})
	}

	router.Run(":8080")
}