package main

import "fmt"

// go 语言里是值传递，修改的是副本
// &x 取地址
// *x 解引用

// 普通数据类型
func change1(x *int) {
	*x = 800
}

// 切片
// go语言在设计切片的时候，里面的数据就是指针，但是切片本身还是一个副本
func changeSlice(s []string) {
	s[0] = "Golang"
	s = append(s, "Python") // 不生效，需要用指针的方式
}

// 数组
// 数组是值传递

// Map
func changeMap(m map[string]int) {
	m["Chinese"] = 100 // 生效
	delete(m, "Math")  // 生效
}

// Struct
// 是值传递
type Person struct {
	Name string
	Age  int
	Sex  bool
}

func changeStruct(p Person) {
	p.Age = 800 // 不生效，需要用指针
}

func changeStruct2(p *Person) {
	p.Age = 800 // 简写，变量不用加*，完整写法为 (*p).Age
}

func main() {

	// 普通数据类型
	x := 10
	change1(&x)
	fmt.Println(x) // output：800

	// 切片
	s := []string{"Java", "C++", "Rust"}
	changeSlice(s)
	fmt.Println(s) // output: Java 被改成 Golang

	// Map
	m := map[string]int{
		"Chinese": 80,
		"Math":    90,
		"English": 80,
	}
	changeMap(m)
	fmt.Println(m) // output: map[Chinese:100 English:80]

	// Struct
	person := Person{
		Name: "Mixsu",
		Age:  21,
		Sex:  true,
	}

	changeStruct(person)
	fmt.Println(person)

	changeStruct2(&person)
	fmt.Println(person)
}
