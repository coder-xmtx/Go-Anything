package main

import (
	"fmt"
	"strconv"
)

func main() {

	// ====== 声明变量 ======

	// 变量声明（隐式）
	var name = "小明同学"

	// 变量声明（显式）
	var age int = 21

	// 变量简写（用的多一点，但是只能写在函数体内部）
	sex := "Male"

	fmt.Println(name, age, sex)

	// ====== 数据类型 ======

	// string
	text := "Hello "
	text += "Mixsu"
	fmt.Println(text)

	// int 和 unit

	// float32（4字节） 和 float64（8字节）
	// 有效精度分别为 5~7位 和 15~16位
	var pi float32 = 3.1415926535
	fmt.Println(pi) // output: 3.1415927

	// bool

	// ====== 类型转换 ======

	// 整数与小数
	floatNum := 3.14
	changeToInt := int(floatNum)
	fmt.Println(changeToInt)

	// 字符串转数字（使用strconv.Atoi函数）
	strNumber := "1"
	intNumber := 2
	strToInt, _ := strconv.Atoi(strNumber) // 该函数返回一个int数字和error错误，但是我们不需要error错误信息，故采用 _ 占位
	fmt.Println(intNumber + strToInt)      // output: 3

	// 数字转字符串
	intToStr := strconv.Itoa(intNumber)
	fmt.Println(strNumber + intToStr) // output: 12（字符拼接）

	// TIP 格式化打印
	fmt.Printf("类型：%T 值：%#v", intToStr, intToStr)

	// 字符串转布尔值
	// 表示true："1" "t" "T" "true" "True" "TRUE"
	// 表示false："0" "f" "F" "false" "False" "FALSE"
	str := "1"
	boolean, _ := strconv.ParseBool(str)
	fmt.Println(boolean) // output: true

	// 布尔值转字符串
	newStr := strconv.FormatBool(true)
	fmt.Printf("类型：%T 值：%#v", newStr, newStr) // output: 类型：string 值："true"
}
