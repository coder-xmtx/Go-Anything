package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// << 10 表示 KB，20 表示 MB，30 表示 GB

// 指定文件最大的大小
const MaxSize = 5 << 20

// 指定文件后缀格式
var allowExt = map[string]bool{
	".png": true,
	".jpg": true,
}

func main() {
	// 创建路由
	router := gin.Default()

	router.POST("/upload", func(ctx *gin.Context) {
		// 通过前端上传的字段 读取图片
		file, err := ctx.FormFile("file") // 单文件上传
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 错误校验
			return
		}
		if file.Size > MaxSize {
			ctx.String(http.StatusRequestEntityTooLarge, "文件过大") // 文件大小校验
			return
		}
		if !allowExt[filepath.Ext(file.Filename)] {
			ctx.String(http.StatusBadRequest, "该后缀暂时不支持") // 后缀名校验
			return
		}

		// filepath.Join 拼接路径
		// filepath.Base 过滤路径（只带文件名字，没有文件名前的本机路径）
		dst := filepath.Join("./files/", filepath.Base(file.Filename))

		// 找个文件夹存储图片
		ctx.SaveUploadedFile(file, dst)

		ctx.String(http.StatusOK, "上传成功")
	})

	router.POST("/uploads", func(ctx *gin.Context) {
		// 通过前端上传的字段 读取图片
		form, err := ctx.MultipartForm() // 多文件上传
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // 错误校验
			return
		}

		files := form.File["files"]
		for _, file := range files {

			dst := filepath.Join("./files/", filepath.Base(file.Filename))

			ctx.SaveUploadedFile(file, dst)
		}

		ctx.String(http.StatusOK, "上传成功")
	})

	// 启动服务
	router.Run()
}
