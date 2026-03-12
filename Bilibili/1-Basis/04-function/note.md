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

闭包是一个函数“捕获”了它外部作用域的变量。即使外部函数已经返回，这些变量依然存在（不会被销毁），因为内部函数仍然持有它们的引用。

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

讲解该例子：

1. `createCounter()`被调用，初始化局部变量`count := 0`
2. 该函数返回一个匿名函数，这个匿名函数**引用**了变量`count`
3. 匿名函数连同它引用的环境（即`count`）一起构成了一个闭包，即使`createCounter()`执行完毕，`count`也不会消失，因为它被闭包持有
4. 每次调用`counter()`，实际上都是在操作同一个`count`变量，所以它会递增

## 函数作为参数

这一块的代码完全看不懂

```go
package main

import "fmt"

func main() {
	strs := []string{"hello", "world"}
	result := process(strs, func(str string) string {
		return str + "!"
	})
	fmt.Println(result)
}

// 定义了一个名为 StringProcessor 的类型
// 它表示所有接受一个 string 参数并返回一个 string 的函数
// 之后就可以用这个类型来声明变量或参数
type StringProcessor func(string) string

func process(strs []string, processor StringProcessor) []string {
	var result []string
	for _, str := range strs {
		result = append(result, processor(str))
	}
	return result
}
```

## 可变参数

```go
package main

import "fmt"

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))

	array := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(array...))
}

// 可变参数

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
```
