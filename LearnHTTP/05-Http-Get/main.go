package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 使用著名的测试接口
	url := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取状态和状态码
	fmt.Println("Status:", resp.Status)
	fmt.Println("Status Code:", resp.StatusCode)

	// 防御性编程，确保关闭
	defer resp.Body.Close()
}
