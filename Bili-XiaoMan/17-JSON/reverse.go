package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 结构体示例
type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Native  string  `json:"native"`
	Golang  bool    `json:"golang"`
	Address *string `json:"address"`
}

// 反序列化 - JSON数据格式 -> Golang数据格式
func main() {

	// 接收JSON
	list := []Person{}

	data, _ := os.ReadFile("data.json")

	// 反序列化
	json.Unmarshal(data, &list)

	fmt.Println(list)
}
