package main

import "fmt"

func main() {

	// ====== 数组 ======

	// 写法
	hobby := [3]string{"sing", "jump", "rap"}
	fmt.Println(hobby)

	// 索引
	fmt.Println(hobby[0])

	// 修改
	hobby[2] = "basketball"

	// 长度 len
	fmt.Println(hobby, len(hobby))

	// ====== 切片 ======

	// 可看作可变长度的数组

	// 写法
	s := []int{1, 2, 3}

	// 后面插入
	s = append(s, 4, 5, 6)
	fmt.Println(s) // output: [1 2 3 4 5 6]

	// 前面插入
	//（go语言并没有直接提供方法，不过可以利用解包运算符来实现）
	s = append([]int{-1, 0}, s...)
	fmt.Println(s) // output: [-1 0 1 2 3 4 5 6]

	// 中间插入
	// 先知道切片截取的方法： s[开始索引:结束索引] 左闭右开
	// 不写开始索引默认从0开始 不写结束索引默认截取到末尾（包含末尾）
	s = append(s[:4], append([]int{100, 200}, s[4:]...)...)
	fmt.Println(s) // output: [-1 0 1 2 100 200 3 4 5 6]
}
