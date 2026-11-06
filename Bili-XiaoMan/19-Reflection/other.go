package main

import "reflect"

// main.go 中的让reflect的值拥有方法的另一种写法

func Other() {
	var age Age = 90
	val := reflect.ValueOf(age)
	m := val.MethodByName("Example") // 写方法的名字
	if m.IsValid() {
		m.Call([]reflect.Value{ // Call 传参数
			reflect.ValueOf("hhh"),
			reflect.ValueOf(100),
		})
	}
}
