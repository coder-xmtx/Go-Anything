package main

import "fmt"

func main() {

	//  -- 变量 --
	var str string = "hello1" // 标准定义
	str2 := "hello2"          // 简写（用的比较多）
	var str3 = "hello3"       // 简写

	// 例子：数值交换
	str, str2 = str2, str
	fmt.Println(str, str2, str3)

	// 零值
	testZeroValue()

	// 类型转换
	testTypeConversion()

	// 自定义类型
	testInt()

	// 数组
	var aaa []int = []int{1, 2, 3}
	fmt.Println(aaa)
}

// 零值
func testZeroValue() {
	var a int
	var b float64
	var c bool
	var d string
	fmt.Println(a, b, c, d) // 0 0 false 空格
}

// 类型转换
func testTypeConversion() {
	var a int = 10
	var b float64 = 11.1
	fmt.Println(a + int(b))
}

// 自定义类型
type myInt int // 类型别名

func testInt() {
	var a myInt = 10
	var b int = 11
	fmt.Println(a + myInt(b))
}

// 结构体
type User struct {
	Name string
	Age  int
}

func (u User) GetName() string {
	return u.Name
}

// 常量
const (
	statusOK = iota + 1
	statusError
	statusNotFound
)
