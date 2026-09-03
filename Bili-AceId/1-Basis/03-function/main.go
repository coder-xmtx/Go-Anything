package main

import "fmt"

// 返回单个值
func foo1(a, b int) int {
	return a + b
}

// 返回多个值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("---foo2---")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	return 100, 200
}

// 返回多个值，命名
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---foo3---")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	r1 = 100
	r2 = 200
	return
}

func main() {
	a := foo1(1, 2)
	fmt.Println("a =", a)

	return1, return2 := foo2("hello", 10)
	fmt.Println("return1 =", return1, "return2 =", return2)

	r1, r2 := foo3("hello", 10)
	fmt.Println("r1 =", r1, "r2 =", r2)
}
