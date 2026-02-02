# Variable and Type

## 1. 变量定义

在 Golang 中，采用 `var xxx type` 的形式来定义变量

```go
package main

import "fmt"

func main() {
    var name string = "Mixso"
	var year int = 2026
	fmt.Println("name:", name, "year:", year)
}
```

## 2. 类型推断

在 Golang 中，不显式写明类型，则用 `:=` 来定义变量，Golang 会进行类型推断

```go
package main

import "fmt"

func main() {

	// 定义变量的一般写法
	var number int = 10

	// 使用类型推断
	number2 := 20

	fmt.Println(number, number2)
}
```
