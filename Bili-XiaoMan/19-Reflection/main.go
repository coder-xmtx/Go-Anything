package main

import (
	"fmt"
	"reflect"
)

type Age int

func (n Age) Example(param string, num int) {

}

// 反射 reflect
func main() {

	// 反射可以确定值的类型
	var age Age = 20
	t := reflect.TypeOf(age) // TypeOf 确定类型
	fmt.Println(t.Kind())    // .Name 拿到明面上的类型  .Kind 拿到底层类型
	if t.Kind() == reflect.Int {
		fmt.Println("INT TYPE")
	}

	// 反射可以修改值
	var other Age = 21
	val := reflect.ValueOf(&other).Elem() // 通过 .Elem 拿到值
	// .CanSet 是否允许修改这个值   .IsVailid 是否为有效值
	if val.CanSet() && val.IsValid() && val.Kind() == reflect.Int {
		val.SetInt(100) // 修改int值
	}
	fmt.Println(other, val) // 都修改为100了

	// val挂不了方法，因为它的类型是 reflect.Value
	// 此时需要用interface方法转为any类型，再用类型断言转换
	new_val := val.Interface().(Age)
	new_val.Example("test", 20)
}
