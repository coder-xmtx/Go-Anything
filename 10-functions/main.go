package main

import "fmt"

// 创建函数 func function_name(parameter) return_type {}
func add(a int, b int) int {
	return a + b
}

// 多返回值的函数
func SumAndProduct(x, y int) (int, int) {
	return x + y, x * y
}

// 命名返回值
func SumAndProduct2(x, y int) (sum int, product int) {
	sum = x + y
	product = x * y
	return
}

// 可变参数
func sumAll(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func main() {
	sum1 := add(1, 2)
	fmt.Println(sum1)

	a, b := SumAndProduct(1, 2)
	fmt.Println(a, b)

	// 只想要一个返回值可以用_省略
	_, c := SumAndProduct(1, 2)
	fmt.Println(c)

	// 可变参数例子
	sum2 := sumAll(1, 2, 3, 4, 5)
	fmt.Println(sum2)
	values := []int{1, 2, 3, 4, 5}
	sum3 := sumAll(values...)
	fmt.Println(sum3)

	// 匿名函数
	res := func(n int) int {
		return n + 1
	}
	fmt.Println(res(10))

	// 匿名函数直接调用
	res1 := func(n int) int {
		return n * 2
	}(10)
	fmt.Println(res1)
}
