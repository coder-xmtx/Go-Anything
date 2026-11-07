package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	// 设置状态码
	w.WriteHeader(http.StatusOK)
	// 设置响应体
	res := map[string]any{
		"ok":       true,
		"message":  "JSON encode sucessful",
		"datetime": time.Now().UTC(),
	}
	// 将 map 映射编码为 JSON 格式，写入响应体
	json.NewEncoder(w).Encode(res)
}

func main() {

	// 注册路由
	http.HandleFunc("/ok", successHandler)

	// 监听并启动端口
	err := http.ListenAndServe(":8081", nil)
	fmt.Println(err)
}
