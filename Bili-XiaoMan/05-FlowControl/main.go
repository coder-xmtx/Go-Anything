package main

import "fmt"

func main() {
	var age int
	fmt.Print("请输入用户年龄：")

	// 下面这行代码，如果仅仅是写age，则无法给原始的变量age赋值，因为函数的传参是复制，两个的内存地址不一样
	fmt.Scan(&age)
	fmt.Println("用户输入的数据：", age)

	// ====== if ======

	if age >= 18 {
		fmt.Println("用户已成年")
	} else {
		fmt.Println("用户未成年")
	}

	// ====== switch ======
	// 注意，go语言的switch 不需要 break，会自动截断

	status := 1

	switch status {
	case 1:
		fmt.Println("在线")
	case 2:
		fmt.Println("离线")
	case 3:
		fmt.Println("忙碌")
	default:
		fmt.Println("状态异常")
	}

	// ====== 逻辑运算符 ======
	// && 与  || 或  ! 非
}
