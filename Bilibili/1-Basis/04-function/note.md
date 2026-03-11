# 04-函数

基本用法

```go
func 函数名(参数列表) 返回值列表 {
    // 函数体
    return 返回值
}
```

匿名函数与闭包

```go
// 匿名函数
func(参数列表) 返回值列表 {
    // 函数体
}

// 闭包
// 函数内部定义的函数，可以访问外部函数的变量
```

## 普通写法

```go
package main

import "fmt"

func main() {
	_, err := validateUser("admin", "12345")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("用户注册成功")
	}
}

// 用户注册校验函数
func validateUser(username, password string) (bool, error) {
	if len(username) < 3 {
		return false, fmt.Errorf("用户名必须大于3个字符")
	}
	if len(password) < 6 {
		return false, fmt.Errorf("密码必须大于6个字符")
	}
	return true, nil
}
```

## Go语言的函数可以有多个返回值

```go
package main

import "fmt"

func main() {
	// 调用函数
	name, age, err := QueryUser(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("name:", name, "age:", age)
}

// 模拟查询用户,返回用户信息（名字，年龄）
func QueryUser(userId int) (string, int, error) {
	// 参数校验
	if userId <= 0 {
		return "", 0, fmt.Errorf("invalid user id")
	}

	// 模拟数据库
	users := map[int]struct {
		Name string
		Age  int
	}{
		1: {Name: "Tom", Age: 10},
		2: {Name: "Jerry", Age: 20},
		3: {Name: "Alice", Age: 30},
	}

	// 查询
	if user, ok := users[userId]; ok {
		return user.Name, user.Age, nil
	}

	return "", 0, fmt.Errorf("用户不存在")

}
```

## 匿名函数与闭包

这一块没听懂

```go
package main

import "fmt"

func main() {
	countFunc := createCounter()

	for i := 0; i < 10; i++ {
		fmt.Println(countFunc())
	}
}

// 创建计数器
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
```

## 函数作为参数
