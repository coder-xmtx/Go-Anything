package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type testRequest struct {
	Name string `json:"name"`
}

// 封装一个函数自己写JSON
func writeJSON(w http.ResponseWriter, status int, data any) {
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	// 设置状态码
	w.WriteHeader(status)
	// JSON
	json.NewEncoder(w).Encode(data)
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	// 示例：只允许POST请求
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "only POST is allowed",
		})
		return
	}

	var request testRequest

	// 创建Decoder，数据来源为r.body，即请求体
	dec := json.NewDecoder(r.Body)

	// Decode开始读取解码，并写入request变量
	if err := dec.Decode(&request); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "Invalid JSON format",
		})
		return
	}

	// request为空的时候返回
	request.Name = strings.TrimSpace(request.Name)
	if request.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "name must not be empty",
		})
		return
	}

	// 一切正常，写回响应体
	writeJSON(w, http.StatusOK, map[string]any{
		"ok":       true,
		"data":     request,
		"datetime": time.Now().UTC(),
	})

	// 防御性编程，关闭请求体body，也可不写
	defer r.Body.Close()
}

func main() {

	// 注册路由
	http.HandleFunc("/test", testHandler)

	// 监听并启动端口
	err := http.ListenAndServe(":8081", nil)
	fmt.Println(err)
}
