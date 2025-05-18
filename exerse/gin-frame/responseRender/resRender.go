package main

import (
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/someJSON", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	r.GET("/userJSON", func(ctx *gin.Context) {
		user := User{Name: "Alice", Age: 25}
		ctx.JSON(http.StatusOK, user)
	})
	r.GET("/unicodeJSON", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"lang": "go 语言",
			"tag":  "<br>",
		})
	})
	r.GET("/pureJSON", func(ctx *gin.Context) {
		ctx.PureJSON(http.StatusOK, gin.H{
			"lang": "go语言",
			"tag":  "<br>",
		})
	})

	type Product struct {
		XMLName xml.Name `xml:"product"`
		ID      string   `xml:"id,attr"`
		Name    string   `xml:"name"`
		Price   float64  `xml:"price"`
	}
	r.GET("/someXML", func(ctx *gin.Context) {
		product := Product{ID: "p123", Name: "go go go back!", Price: 29.99}
		ctx.XML(http.StatusOK, product)
	})

	r.GET("/someYAML", func(ctx *gin.Context) {
		ctx.YAML(http.StatusOK, gin.H{
			"a": 1, "b": gin.H{
				"c": "d",
			},
		})
	})

	r.GET("/someString", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "some string:%s,%d", "hello", 123)
	})

	r.LoadHTMLGlob("exerse\\gin-frame\\responseRender\\template/*")
	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "主页",
			"user":  "professor gin",
		})
	})
	r.GET("/local/file", func(ctx *gin.Context) {
		ctx.File("./files/my_document.pdf")
	})
	r.GET("/download/file", func(ctx *gin.Context) {
		ctx.FileAttachment("./files/my_document.pdf", "my_document.pdf")
	})

	r.GET("/stream", func(ctx *gin.Context) {
		reader := strings.NewReader("这是一个数据流...")
		contentlength := int64(reader.Len())
		contentType := "text/plain;charset=utf-8"
		extraHeaders := map[string]string{
			"Content-Disposition": "attachment; filename=file.txt",
		}
		ctx.DataFromReader(http.StatusOK, contentlength, contentType, reader, extraHeaders)
	})

	r.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/index")
	})
	r.Run()
}
