package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRouter() *gin.Engine{
	r:=gin.Default()
	r.GET("/ping",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	return r
}

func TestPingRouter(t *testing.T) {
	router:=SetupRouter()
	req,_:=http.NewRequest("GET","/ping",nil)
	w:=httptest.NewRecorder()
	router.ServeHTTP(w,req)
	assert.Equal(t,http.StatusOK,w.Code)
	expectedBodY:=`{"message":"pong"}`
	assert.JSONEq(t,expectedBodY,w.Body.String())
}