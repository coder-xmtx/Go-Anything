package main

import "fmt"

// const 来定义枚举类型
// 可以在const()添加一个关键字iota，每行的iota都会自动递增1，第一行iota的值为0
// 注意：iota只能在const内部使用
const (
	x = iota // 0
	y        // 1
	z        // 2
)

const (
	a, b = iota + 1, iota + 2 // 1,2  iota = 0
	c, d                      // 2,3  iota = 1

	e, f = iota * 1, iota * 2 // 2,4  iota = 2
)

func main() {
	const number int = 10
	fmt.Println("number =", number) // number = 10
}
