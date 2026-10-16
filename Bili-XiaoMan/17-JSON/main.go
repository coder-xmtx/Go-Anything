package main

import (
	"encoding/json"
	"fmt"
)

// 结构体示例
// address需要默认为空，但是go结构体不能直接写nil，所以用指针来表示
// 注意：在每个字段后面使用结构体标签，使得转成JSON的时候改为小写
type Person struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Native  string  `json:"native"`
	Golang  bool    `json:"golang"`
	Address *string `json:"address"`
}

// 序列化 - Golang数据格式 -> JSON数据格式
func main() {
	person := Person{
		Name: "Harry",
		Age:  20,
	}

	// 将结构体转成JSON格式
	data, _ := json.Marshal(person)

	fmt.Println(string(data))
}
