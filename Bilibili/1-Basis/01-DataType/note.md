# 01-数据类型粗讲

跟其他语言一样，Go语言中也有许多数据类型，下面来对它们有个初步认识。

## 常量

常量使用`const`关键字修饰，格式为`const name type`

1. 可以对常量单一定义

```go
const a int = 10  // 直接定义
const b = "hello"  // 类型推断
```

2. 也可以组合定义

```go
const (
    c int = 10  // 直接定义
    d = "learn go"  // 类型推断
)
```

## 变量

变量使用`var`定义，格式为`var name type`

【注意】`:=`只能在函数内使用，包级别变量需要用`var`

1. 可以对变量单一定义

```go
var a int = 10 // 直接定义
var b = 20  // 类型推断
c := 100  // 类型推断
```

2. 也可以组合定义

```go
var (
    a string = "hello"  // 直接定义
    b = 10  // 类型推断
)
```

【例子】交换两个变量的值

```go
var str1 = "hello 1"
var str2 = "hello 2"

str1, str2 = str2, str1 // 数值交换
```

## 零值

零值在go中算是比较特殊的地方，未具体赋值的变量也会有一个默认值，下面用代码来演示。

```go
import "fmt"

var a int
var b float64
var c bool
var d string

fmt.Println(a, b, c, d)  // 输出 0 0 false 空格
```

## 自定义类型

有时候我们可能需要更有语义的数据类型，所以给原数据类型起别名

```go
type personNumber int  // 这时候personNumber也是个数据类型，类型为int

var a personNumber = 10
```

## 类型转换

和其他语言一样，Go语言也有类型转换

```go
import "fmt"

var a int = 10
var b float64 = 11.1

fmt.Println(a + int(b)) // 输出 21
```

## 数组

数组用`[]type`来定义

## 切片

## 映射map

## 结构体
