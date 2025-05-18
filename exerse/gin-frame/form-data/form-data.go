package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type LoginForm struct {
// 	Username string `form:"username" binding:"required"`
// 	Passward string `form:"password" binging:"required"`
// 	Remember bool   `form:"remember"`
// }

// func main() {
// 	r := gin.Default()
// 	r.POST("/login",func(ctx *gin.Context) {
// 		var form LoginForm
// 		if err:=ctx.ShouldBind(&form);err!=nil{
// 			ctx.JSON(http.StatusBadRequest,gin.H{
// 				"error":err.Error(),
// 				"message":"输入参数有误,请检查吧",
// 			})
// 			return 
// 		}
// 		if form.Username=="admin"&&form.Passward=="passward"{
// 			ctx.JSON(200,gin.H{
// 				"status":"登录成功",
// 				"username":form.Username,
// 				"passward":form.Passward,
// 			})
// 		}else {
// 			ctx.JSON(http.StatusUnauthorized,gin.H{
// 				"status":"用户名或密码错误",
// 			})
// 		}
// 	})
// 	r.Run()
// }

// type Jsondata struct{
// 	Email string `json:"email" binding:"required,email"`
// 	Passward string `json:"password" binding:"required,min=6"`
// 	Age int `json:"age" binding:"omitempty,gte=1,lte=130"`
// }

// func main(){
// 	r:=gin.Default()
// 	r.POST("/json",func(ctx *gin.Context) {
// 		var req Jsondata
// 		if err:=ctx.ShouldBind(&req);err!=nil{
// 			ctx.JSON(http.StatusBadRequest,gin.H{
// 				"err":err.Error(),
// 			})
// 			return 
// 		}
// 		ctx.JSON(http.StatusCreated,gin.H{
// 			"message":"用户创建成功",
// 			"email":req.Email,
// 			"age":req.Age,
// 		})
// 	})
// 	r.Run()
// }

// type Pagenation struct{
// 	Page string `form:"page" binding:"required,gte=1"`
// 	Limit int `form:"limit" binding:"required,gte=1,lte=100"`
// }
// func main(){
// 	r:=gin.Default()
// 	r.GET("/page",func(ctx *gin.Context) {
// 		var pa Pagenation
// 		if err:=ctx.BindQuery(&pa);err!=nil{
// 			ctx.JSON(http.StatusBadRequest,gin.H{
// 				"msg":err.Error(),
// 			})
// 		}
// 		ctx.JSON(200,gin.H{
// 			"page":pa.Page,
// 			"limit":pa.Limit,
// 		})
// 	})
// 	r.Run()
// }

type UriUrl struct{
	ID string `uri:"id" binding:"required,uuid"`
}
func main(){
	r:=gin.Default()
	r.GET("/user/:id",func(ctx *gin.Context) {
		var uri UriUrl
		if err:=ctx.ShouldBindUri(&uri);err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx.JSON(200,gin.H{
			"id":uri.ID,
			"msg":"成功访问!",
		})
	})
	r.Run()
}
