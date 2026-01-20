package main

import (
	"fmt"
	"strings"
)

func main() {

	// 变量定义 var xxx type
	var name string = "mixso studio"
	var year int = 2026

	// 类型推断 :=
	age := 12

	fmt.Println(strings.ToUpper(name), year, age)
}
