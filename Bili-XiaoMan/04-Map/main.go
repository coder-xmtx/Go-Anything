package main

import "fmt"

// ======= Map映射 ======
// Map 是无序的

func main() {

	// ====== 基本写法（:=） ======
	scores := map[string]int{
		"Chinese": 90,
		"Math":    100,
		"English": 80,
	}
	fmt.Println(scores)

	// 新增属性
	scores["Physics"] = 80
	fmt.Println(scores)

	// 删除属性
	delete(scores, "Chinese")
	fmt.Println(scores)

	// 读取属性
	// 方式一：scores[键]
	// 方式二：value, ok := scores[键]    若该属性存在，则 ok 返回 true
	value, ok := scores["Math"]
	fmt.Println(value, ok)

	// ====== 其他写法（var） ======

	var ages map[string]int     // 这样写会报 nil 错误，需要使用 make 来创建容量
	ages = make(map[string]int) // 没写 size 参数，具体情况具体分析
	ages["Mike"] = 90
	fmt.Println(ages)

	// 两行的声明写法还是有点啰嗦了
	cups := make(map[string]int)
	fmt.Println(cups)
}
