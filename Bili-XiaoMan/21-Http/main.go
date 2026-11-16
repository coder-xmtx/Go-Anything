package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// http.NewServeMux 创建多路路由复用器 路由实例独立 不会被污染
	// 如果直接采用 http.HandleFunc 创建的路由是公开的 可能被污染
	mu := http.NewServeMux()
	mu.HandleFunc("GET /user/list", func(w http.ResponseWriter, r *http.Request) {
		// w Response => 返回什么资源
		// r Request => 读取什么资源

		// 假设返回一个字符串
		// 这里需要知道一个语法糖：[]byte(类型) => []byte{}
		// w.Write([]byte("hello"))

		// 查询 id 字段
		targetId := r.URL.Query().Get("id")

		// 结构体源数据
		userList := []User{
			{Id: 1, Name: "Mixsu"},
			{Id: 2, Name: "Vince"},
			{Id: 3, Name: "Bryony"},
		}
		// 存放筛选的结果
		result := []User{}
		if targetId == "" {
			result = userList
		} else {
			for _, item := range userList {
				if strconv.Itoa(item.Id) == targetId {
					result = append(result, item)
				}
			}
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(result)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mu,
	}

	// 启动服务
	server.ListenAndServe()
}
