package main

import "fmt"

func main() {
	// 创建map（映射）,有点像python的字典  map[key]value
	ages := map[string]int{
		"alice":   31,
		"bob":     32,
		"charlie": 33,
	}
	fmt.Println(ages, len(ages))

	// 创建空映射
	var scores map[string]int
	fmt.Println(scores, len(scores))

	scores = make(map[string]int)
	scores["alice"] = 100
	fmt.Println(scores, len(scores))

	// 删除键值对
	users := map[string]string{
		"u1": "Mixsu",
		"u2": "Bob",
	}
	delete(users, "u2")
	fmt.Println(users)

	// value ok 检查是否存在该键值对
	points := map[string]int{
		"a": 10,
		"b": 0, // 0 是一个有效值
	}
	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"]) // 但是我们打印不存在的键，也会得出 0

	value, ok := points["c"] // 检查“c”的value是否存在，存在则ok返回true
	fmt.Println("c", value, ok)

	// for 循环
	prices := map[string]int{
		"apple":  10,
		"banana": 20,
	}
	total := 0
	for item, price := range prices {
		fmt.Println(item, price)
		total += price
	}
}
