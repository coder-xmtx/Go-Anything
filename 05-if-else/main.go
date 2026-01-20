package main

import "fmt"

func main() {

	// 例子：成绩判断
	score := 95

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 60 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}

	// 例子：短路逻辑判断
	if score := 95; score >= 90 {
		fmt.Println("A")
	}

	perPrice := 30
	num := 3

	if totalPrice := perPrice * num; totalPrice > 100 {
		fmt.Println("大于100")
	}
}
