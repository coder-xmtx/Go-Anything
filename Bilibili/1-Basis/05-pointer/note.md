# 05-指针

内存地址：每个变量在内存中都有唯一的地址

指针变量：存储另一个变量内存地址的变量

指针操作：

```go
var p *int // 声明指针
p = &variable // 取地址
*p = 100 // 解引用（通过指针修改变量值）
```

指针作为函数参数：

- **值传递**：传递的是变量值的副本。函数内部修改副本，不影响原变量。
- **指针传递**：传递的是指针变量的副本（也就是内存地址的副本）。虽然指针本身是副本，但它指向的内存地址与原指针相同，所以通过这个副本指针修改内存地址上的内容，原变量也会被修改。

例子：

```go
package main

import "fmt"

type User struct {
	Id      int
	Name    string
	Balance float64
}

// 值传递
func updateBalanceValue(u User, amount float64) {
	u.Balence += amount
	fmt.Println("Value更新后的余额是:", u.Balence)
}

// 指针传递
func updateBalancePointer(u *User, amount float64) {
	u.Balence += amount
	fmt.Println("Pointer更新后的余额是:", u.Balence)
}

func main() {
	user1 := User{Id: 1, Name: "小明", Balence: 100.0}
	updateBalanceValue(user1, 200.0)
	fmt.Println("user1:", user1)

	user2 := User{Id: 2, Name: "小红", Balence: 100.0}
	updateBalancePointer(&user2, 200.0)
	fmt.Println("user2:", user2)

	users := []*User{
		&User{Id: 3, Name: "Mixsu", Balence: 100.0},
		&User{Id: 4, Name: "Vince", Balence: 100.0},
	}

	for _, user := range users {
		user.Balence += 100.0
	}

	for _, u := range users {
		fmt.Println("用户", u.Name, "更新后的余额是:", u.Balence)
	}
}
```
