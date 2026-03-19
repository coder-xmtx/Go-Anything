package main

import "fmt"

func main() {

	// == 单变量声明 ==

	// 方法一：声明一个变量，默认值为零值
	var number int
	fmt.Println("a =", number)             // 0
	fmt.Printf("Type of a = %T\n", number) // int

	// 方法二：声明一个变量，初始化一个值
	var number2 int = 10
	fmt.Println("a =", number2)             // 10
	fmt.Printf("Type of a = %T\n", number2) // int

	// 方法三：在初始化的时候，省去数据类型，自动推断
	var number3 = 20
	fmt.Println("a =", number3)             // 20
	fmt.Printf("Type of a = %T\n", number3) // int

	// 方法四：在初始化的时候，省去数据类型，自动推断，简写 （常用）
	// 注意：只能在函数内部使用
	number4 := 30
	fmt.Println("a =", number4)             // 30
	fmt.Printf("Type of a = %T\n", number4) // int

	// == 多变量声明 ==

	var x, y, z int = 1, 2, 3
	fmt.Println("x =", x, "y =", y, "z =", z) // x = 1 y = 2 z = 3

	// == 多行的多变量声明 ==

	var (
		a int    = 1
		b bool   = true
		c string = "hello"
	)
	fmt.Println("a =", a, "b =", b, "c =", c) // a = 1 b = 2 c = 3
}
