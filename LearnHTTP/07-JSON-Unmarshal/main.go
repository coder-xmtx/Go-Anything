package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// 接收传来的JSON
type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	// 测试接口
	url := "https://catfact.ninja/fact"

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

	// io读取JSON字节流
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		return
	}

	var data CatFactResponse

	// 将JSON转为struct
	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		fmt.Println("JSON unmarshal failed:", err)
		return
	}

	// 打印输出传来的数据
	fmt.Println(data.Fact, data.Length)

	// 防御性编程，确保关闭
	defer resp.Body.Close()
}
