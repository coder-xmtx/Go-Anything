package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Age   int
	Email string
}

func main() {
	user1 := User{
		ID:    1,
		Name:  "mixsu",
		Age:   20,
		Email: "mixsu@example.com",
	}

	fmt.Println(user1, user1.Email)

	// 字段的值是可以修改的
	user1.Name = "mixsu-change"
	fmt.Println(user1, user1.Email)

	// 可以只指定部分字段
	user2 := User{
		Name: "bob",
		Age:  20,
	}
	fmt.Println(user2, user2.Email)
}
