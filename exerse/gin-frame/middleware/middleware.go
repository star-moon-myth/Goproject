package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyloggerMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		//Now返回一个结构体，但是同时time实现了一个方法String，所以这个类型就实现了fmt中的一个接口
		//接口Stringer中就定义了一个方法String，之后，print时其实调用了接口，
		// 当类型同时实现接口时就会输出被String函数包裹的内容了
		startime:=time.Now() 
		requestpath:=ctx.Request.URL.Path
		requestmethod:=ctx.Request.Method
		clientIP:=ctx.ClientIP()
		log.Printf("[Middleware Start]%s %s from %s",requestmethod,requestpath,clientIP)
		ctx.Next()
		endtime:=time.Now()
		latency:=endtime.Sub(startime)
		statuscode:=ctx.Writer.Status()
		log.Printf("[Middleware End]%s %s from %s,latency:%v,statuscode:%v",requestmethod,requestpath,clientIP,latency,statuscode)
		ctx.Header("x-Processed-By", "MyLoggerMiddleware")
	}
}	
func AuthMiddleware() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		token:=ctx.GetHeader("Authoration")
		if token!="123456"{
			log.Printf("[AuthMiddleware] 无效或缺失token")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
				"error":"未授权访问",
			})
			return
		}
		log.Printf("[AuthMiddleware]验证通过")
		ctx.Set("userID","user123")
		ctx.Next()
		log.Printf("[AuthMiddleware]验证结束")
	}
}

func main(){
	r:=gin.New()
	r.Use(MyloggerMiddleware())
	authorized:=r.Group("/admin")
	authorized.Use(AuthMiddleware())
	{
		authorized.GET("/dashboard",func(ctx *gin.Context) {
			userid,exists:=ctx.Get("userID")
			if !exists{
				ctx.JSON(http.StatusInternalServerError,gin.H{
					"error":"未获取到用户ID",
				})
				return 
			}
			ctx.JSON(http.StatusOK,gin.H{
				"message":"欢迎访问dashboard,用户ID为:",
				"user_id":userid,
			})
		})
		authorized.GET("/settings",func(ctx *gin.Context) {
			userID:=ctx.GetString("userID")
			ctx.JSON(http.StatusOK,gin.H{
				"message":"欢迎访问settings,用户ID为:",
				"user_id":userID,
			})
		})
	}
	r.GET("/public/info",MyloggerMiddleware(),func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"欢迎访问公共信息",
		})
	})
	r.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	r.Run()
}