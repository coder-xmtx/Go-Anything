package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

// type 自定义类型
type Name string

// 给自定义类型添加方法
func (n Name) PrefixName() Name {
	return "Hello, " + n
}

func (n Name) RandomName() Name {
	ran := rand.IntN(100) + 1 // 0~100
	return n + Name(strconv.Itoa(ran))
}

// 类型安全，避免参数传递问题
// 比如下面这个函数，如果传参的时候两个参数写反，但是因为都是int，并不会报错
// 此时就可以给int分别起别名区分开

type Uid int
type Oid int

func CreateOrder(uid Uid, oid Oid) (Uid, Oid) {
	return uid, oid
}

func main() {
	var name Name = "Mixsu"
	// 挂方法
	preName := name.PrefixName()
	fmt.Println(preName)
	ranName := name.RandomName()
	fmt.Println(ranName)

	// 类型安全
	var uid Uid = 10
	var oid Oid = 90
	CreateOrder(uid, oid)
}
