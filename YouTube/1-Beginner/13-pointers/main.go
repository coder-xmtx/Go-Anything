package main

import "fmt"

func main() {

	// 指针是一个内存地址
	// 可以通过 * 指针来取值
	// 可以通过 & 指针来取地址

	score := 10
	fmt.Println("Before:", score)

	addScore(&score)
	fmt.Println("After:", score)
}

func addScore(score *int) {
	*score = *score + 1
}
