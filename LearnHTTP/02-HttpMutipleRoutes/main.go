package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome, try to /hello?name=mixsu"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	// Query 查询
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	// 写响应体
	w.Write([]byte("hello, " + name))
}

func main() {

	// 注册路由
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	// 在上一节里使用的是 ListenAndServe，监听和服务
	// 这一节，我们分别单独使用 listen 和 serve

	// 1. 尝试监听端口
	ln, err := net.Listen("tcp", ":8081")

	// 监听失败，记录日志
	if err != nil {
		log.Fatalln("[ PORT 8081 FAILED ]", err)
	}

	// 监听成功，打印信息
	fmt.Println("[ PORT 8081 SUCCEED ]")

	// 2. 再启动 HTTP 服务
	err = http.Serve(ln, nil)
	if err != nil {
		log.Fatalln("[ SERVER STOPPED ]", err)
	}
}
