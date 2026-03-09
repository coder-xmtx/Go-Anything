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
    d     = "learn go"  // 类型推断
)
```

3. 常量可以用`iota`实现枚举

```go
const (
    _  = iota             // 忽略0
    KB = 1 << (10 * iota) // 1 << (10*1) = 1024
    MB                    // 1 << (10*2) = 1048576
    GB
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

零值的好处：变量总是被初始化，避免了未定义行为。

```go
var a int                  // 0
var b float64              // 0
var c bool                 // false
var d string               // 空字符串

var ptr *int               // nil
var slice []int            // nil
var m map[string]int       // nil
var ch chan int            // nil
var fn func()              // nil
var iface interface{}      // nil
```

## 自定义类型

有时候我们可能需要更有语义的数据类型，所以给原数据类型起别名

```go
type personNumber int  // 这时候personNumber也是个数据类型，类型为int

var a personNumber = 10
```

需要强调：自定义类型和原类型是不同的类型，不能直接赋值，除非显式转换。例如：

```go
type personNumber int
var a int = 10
var b personNumber = personNumber(a) // 必须转换
```

## 类型转换

和其他语言一样，Go语言也有类型转换

```go
import "fmt"

var a int = 10
var b float64 = 11.1

fmt.Println(a + int(b)) // 输出 21
```

【例子】字符串与字节切片转换

```go
s := "hello"
b := []byte(s)  // 转换为字节切片
s2 := string(b) // 转换回字符串
```

【例子】数字与字符串转换需要使用strconv包

```go
i, _ := strconv.Atoi("123")   // 字符串转int
s := strconv.Itoa(456)        // int转字符串
```

## 数组array

数组用`[num]type`来定义，这里数组指的是定长的数组

```go
var a [3]int = [3]int{1, 2, 3}  // 标准写法
b := [3]int{1, 2, 3}  // 类型推断
```

## 切片slice

切片是没有定长的数组，用`[]type`定义，切片用的会更多一些

```go
var a []int = []int{1,2,3,4,5}  // 标准写法
b := []int{1,2,3,4,5,6} // 类型推断
```

## 映射map

可以简单理解为键值对，写法为`map[KeyType]ValueType{...}`

```go
var map1 = map[string]int{"Go": 96, "python": 95}  // 标准写法，Key为string类型，Value为int类型
map2 := map[string]int{"Go": 96, "python": 95}  // 类型推断
```

## 结构体

可以看作其他语言中的类，不过有很大不同，格式为`type xxx struct{}`

```go
import "fmt"

// 标准写法
type Person struct {
    name string
    age int
}

// 定义该结构体的方法
func (p Person) information() string {
    return fmt.Sprintf("%s %d", p.name, p.age)
}

// 创建实例
person1 := Person{
    name: "Mixsu",
    age: 20,
}

// 调用方法
person1.information()
```
