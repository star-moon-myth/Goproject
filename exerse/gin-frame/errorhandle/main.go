package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(ErrorhandlingMiddleware())
	r.Use(gin.Recovery())
	r.GET("/test/error1",func(ctx *gin.Context) {
		err:=errors.New("this is a error")
		ctx.Error(err)
		ctx.Abort()
	})
	r.GET("/test/error2",func(ctx *gin.Context) {
		err:=errors.New("这个错误需要立即中止")
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		
	})

	r.GET("/test/error3",func(ctx *gin.Context) {
		type Query struct{
			Value int `form:"value" binding:"required,gt=10"`
		}
		var q Query
		if err:=ctx.ShouldBindQuery(&q);err!=nil{
			ctx.Error(err)
			ctx.Abort()
		}
		ctx.JSON(http.StatusOK,gin.H{
			"value":q.Value,
		})
	})
	r.GET("/test/error4",func(ctx *gin.Context) {
		panic("!!!oh,no~~panic")
	})
	r.Run()
}

func ErrorhandlingMiddleware()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors)>0{
			lasterror:=ctx.Errors.Last()
			log.Printf("[ErrorHandlingMiddleware] %s",lasterror.Err)
			var statuscode int
			var errorResponse gin.H
			var validationErrors validator.ValidationErrors
			if errors.As(lasterror.Err,&validationErrors){
				statuscode=http.StatusBadRequest
				errorResponse=gin.H{
					"error":formatValidationErrors(validationErrors),
				}
			}else if errors.Is(lasterror.Err,sql.ErrNoRows){
				statuscode=http.StatusNotFound
				errorResponse=gin.H{
					"error":"找不到数据",
				}
			}else{
				statuscode=ctx.Writer.Status()
				if statuscode==http.StatusOK{
					statuscode=http.StatusInternalServerError
				}
				errorResponse=gin.H{
					"error":"服务器内部错误",
				}
			}
		ctx.JSON(statuscode,errorResponse)
		}
	}
}

func formatValidationErrors(errs validator.ValidationErrors) map[string]string{
	errors:=make(map[string]string)
	for _,err:=range errs{
		errors[err.Field()]=err.Tag()
	}
	return errors
}