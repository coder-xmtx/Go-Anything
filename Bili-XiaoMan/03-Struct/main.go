package main

import "fmt"

// ====== 结构体 ======

type Car struct {
	Model string
	Color string
}

// 属性规则：首字母大写，表示公开，允许所有包访问；反之亦然
type Person struct {
	Name      string
	Age       int
	Hobby     []string
	IsStudent bool
	MyCar     Car // 结构体嵌套
	// 结构体嵌套也可直接用结构体的名字，不需要额外写一个名字，访问的时候就可以直接person.Model，而不用person.MyCar.Model
}

// ====== 泛型结构体 ======
// 有些属性可能存在多种数据类型，这时候需要用到泛型
// 在这个例子里，顾客的号码可能是座机（字符串）或者手机号（int类型）
type Customer[PhoneType string | int] struct {
	Name  string
	Age   int
	Phone PhoneType
}

func main() {
	person := Person{
		Name:      "Mixsu",
		Age:       21,
		Hobby:     []string{"play"},
		IsStudent: true,
		MyCar: Car{
			Model: "BMW",
			Color: "Black",
		},
	}
	fmt.Println(person)

	// 确定用户手机号为string类型
	customer := Customer[string]{
		Name:  "Vince",
		Age:   20,
		Phone: "010-10086",
	}
	fmt.Println(customer)
}
