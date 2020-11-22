package main

import (
	"Picrsc/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 17  // 8 MiB
	pic:= router.Group("/api/picture")
	{
		pic.POST("", UploadFile)
		pic.GET("",GetImages)
		pic.GET("/hello",Helloworld)
	}
	router.Run(":8080")
}
func UploadFile(c *gin.Context){
		// single file
	file, err := c.FormFile("file")
	tag := c.PostForm("tag")
	token:=c.PostForm("token")
	if token!=util.Settings.Token{
		c.JSON(http.StatusBadRequest,gin.H{
			"msg":"Token不正确，无法上传图片",
		})
		return
	}
	if file.Size>5*(2 << 20){
			c.JSON(http.StatusBadRequest,gin.H{
				"msg":"文件太大",
			})
			return
		}
		util.Check(err)
		file.Filename =util.ParseFileName(file.Filename)

		err=c.SaveUploadedFile(file, "../Files/"+file.Filename)
		fmt.Println(err.Error())
		image := util.Image{Url:"146.56.199.136/Files/"+file.Filename,IsDelete:false,Tag:tag}
		util.AddImage(image)
		c.JSON(http.StatusOK, gin.H{
			"msg":"成功上传图片",
		})
	return
}
func GetImages(c *gin.Context){
	page,_ :=strconv.Atoi( c.Query("page"))
	pagesize,_ := strconv.Atoi(c.Query("page_size"))
	images := util.GetImages(page,pagesize)
	c.JSON(http.StatusOK,gin.H{
		"images":images,
	})
}
func Helloworld(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"msg":"helloworld",
	})
}