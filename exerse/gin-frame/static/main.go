package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("static\\templates/*") // 加载模板

    // 提供静态文件
    r.Static("/static", "./static")
    r.StaticFile("/favicon.ico", "./static/favicon.ico")

    // 渲染 HTML 页面
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index_with_static.tmpl", gin.H{
            "title": "带静态文件的页面",
        })
    })

    r.Run() // 默认监听 :8080
}