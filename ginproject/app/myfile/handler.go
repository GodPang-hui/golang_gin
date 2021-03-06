package myfile

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func formHandler(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("userpassword")
	// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
}

func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
	}
	// c.JSON(200, gin.H{"message": file.Header.Context})
	c.SaveUploadedFile(file, "file/"+file.Filename)
	c.String(http.StatusOK, file.Filename)
}

func uploadsHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
	// 获取所有图片
	files := form.File["files"]
	// 遍历所有图片
	for _, file := range files {
		// 逐个存
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
}
