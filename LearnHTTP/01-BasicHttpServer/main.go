package main

import (
	"fmt"
	"net/http"
)

// r 读请求 ， w 写响应
func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only Get is allowed", http.StatusMethodNotAllowed)
		return
	}

	// 写请求成功时的返回的响应体
	w.Write([]byte("Hello from Go Server"))
}

func main() {

	// 注册路由
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Try going to port 8080")

	// 启动服务器，端口 8080
	http.ListenAndServe(":8080", nil)
}
