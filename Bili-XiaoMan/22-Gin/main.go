package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 创建路由
	router := gin.Default()

	// --- 读取 URL 参数 ---

	// 搜索参数
	router.GET("/", func(ctx *gin.Context) {
		// ctx 表示上下文，结合了 w 和 r

		// string 返回字符串
		// ctx.String(http.StatusOK, "value: %v", "hello,Gin")

		// 搜索参数 比如：?name=mixsu&age=21
		// ctx.Query() 查询参数，DefaultQuery() 可以设置一个默认值
		name := ctx.DefaultQuery("name", "none")
		age := ctx.DefaultQuery("age", "none")
		ctx.String(http.StatusOK, "value: %v - %v", name, age)
	})

	// 路径参数    比如：detail/123
	router.GET("/detail/:id", func(ctx *gin.Context) {
		// ctx.Param 接收路径参数
		id := ctx.Param("id")
		ctx.String(http.StatusOK, "value: %v", id)
	})

	// --- 表单请求 ---
	router.POST("/create", func(ctx *gin.Context) {
		name := ctx.DefaultPostForm("name", "none")
		age := ctx.DefaultPostForm("age", "none")
		ctx.JSON(http.StatusOK, gin.H{
			"data": map[string]any{
				"name": name,
				"age":  age,
			},
			"success": true,
		})
	})

	// 发送JSON格式的POST请求
	router.POST("/json", func(ctx *gin.Context) {
		var user User
		// 写入上面的 user 变量
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		// response
		ctx.JSON(http.StatusOK, gin.H{
			"data": user,
		})
	})

	// 启动服务，默认 8080 端口
	router.Run()
}
