package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/*")
	r.GET("/upload",func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK,"upload.tmpl",nil)
	})
	r.POST("/upload/single",func(ctx *gin.Context) {
		file,err:=ctx.FormFile("file")
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":"文件上传失败:"+err.Error(),
			})
			return
		}
		log.Printf("Having received a file %s(filesize:%d bytes)",file.Filename,file.Size)
		dst:="upload/"+file.Filename
		if err:=ctx.SaveUploadedFile(file,dst);err!=nil{
			ctx.JSON(http.StatusInternalServerError,gin.H{
				"error":"文件保存失败:"+err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK,gin.H{
			"message":"文件上传成功",
			"filename":file.Filename,
			"size":file.Size,
			"saved_path":dst,
		})
	})

	r.POST("upload/multi",func(ctx *gin.Context) {
		form,err:=ctx.MultipartForm()
		if err!=nil{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":"文件上传失败:"+err.Error(),
			})
			return
		}
		files:=form.File["files"]
		if len(files)==0{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"error":"没有上传文件",
			})
			return
		}
		savesfiles:=[]gin.H{}
		for _,file:=range files{
			log.Printf("Having received a file %s(filesize:%d bytes)",file.Filename,file.Size)
			dst:="./uploads"+file.Filename
			if err:=ctx.SaveUploadedFile(file,dst);err!=nil{
				log.Printf("保存文件%s失败:%v",file.Filename,err)
				continue 
			}
			savesfiles=append(savesfiles, gin.H{
				"filename":file.Filename,
				"size":file.Size,
				"saved_path":dst,
			})
		}
		ctx.JSON(http.StatusOK,gin.H{
			"message":"文件上传成功",
			"files":savesfiles,
		})
	})
	r.Run()
}