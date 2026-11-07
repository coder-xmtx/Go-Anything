package main

import (
	"fmt"
	"io"
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

	// 响应体状态码要为200
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return
	}

	// 获取响应体 body，io获取字节流
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将字节流转为字符串
	bodyText := string(bodyBytes)

	// 测试接口的内容很多，筛选输出前250个字符，若不满250，则输出全部
	max := min(len(bodyText), 250)
	fmt.Println(bodyText[:max])

	// 防御性编程，确保关闭
	defer resp.Body.Close()
}
