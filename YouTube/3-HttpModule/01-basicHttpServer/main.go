package main

import (
	"fmt"
	"net/http"
)

// 处理请求
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 只允许get请求
	if r.Method != http.MethodGet {
		http.Error(w, "only Get is allowed", http.StatusMethodNotAllowed)
		return
	}

	// r是一个request对象，w 是一个response对象，Write返回一个字符串
	_, _ = w.Write([]byte("Hello from Go net/http server"))
}

func main() {
	// 创建一个路由，处理请求
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("try going to 5000 port")

	// 启动服务，监听5000端口
	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
