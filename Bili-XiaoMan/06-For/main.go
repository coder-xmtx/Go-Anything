package main

import (
	"fmt"
)

func main() {

	games := []string{"PVZ", "GTA5", "It takes Two"}

	scores := map[string]int{
		"Chinese": 90,
		"Math":    100,
		"English": 80,
	}

	str := "你好世界"

	// for 循环

	for i := 0; i < len(games); i++ {
		fmt.Println(games[i])
	}

	fmt.Println()

	// go 语言没有 while 循环，全部用 for 来搞定

	index := 0

	for {
		fmt.Println(games[index])

		index++

		if index >= len(games) {
			break // 跳出循环
		}
	}

	fmt.Println()

	// range 遍历

	for index, value := range games {
		fmt.Println(index, value)
	}

	fmt.Println()

	for key, value := range scores {
		fmt.Println(key, value)
	}

	fmt.Println()

	// 遍历字符串，返回的是字节索引和对应unicode编码
	for byteIndex, uniCode := range str {
		fmt.Println(byteIndex, uniCode)
	}

	fmt.Println()

	// 为了正常遍历字符串，需要先把字符串转成切片
	for index, uniCode := range []rune(str) {
		fmt.Println(index, string(uniCode))
	}

}
