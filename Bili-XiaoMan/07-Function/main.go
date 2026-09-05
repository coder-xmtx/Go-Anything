package main

import "fmt"

// 具名函数
// 注意，具名函数声明不可以再嵌套具名函数，比如下面这个函数不可以嵌套在main函数里声明

// 一般写法
func calc(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

// 泛型
// 在这个场景下，我们想要数字既能支持整数也能支持小数，为此写了一个 T
func add[T int | uint | float64](num1, num2 T) T {
	return num1 + num2
}

// 闭包
// 变量会逃逸到堆里面，如果不断调用函数，下面 n 的值会递增
// 闭包就相当于里面的匿名函数 “绑架” 了外部的变量
func counter() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}

func main() {

	// 具名函数调用
	sum, sub := calc(10, 20)
	fmt.Println(sum, sub)

	// 匿名函数，立即执行
	func() {
		fmt.Println("匿名函数")
	}()

	// 匿名函数，给名字可以手动调用
	fn := func() {
		fmt.Println("hahaha")
	}
	fn()

	// 闭包
	count := counter()   // 创建闭包
	fmt.Println(count()) // 1
	fmt.Println(count()) // 2

}
