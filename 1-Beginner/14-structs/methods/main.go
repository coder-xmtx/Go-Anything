package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	user1 := User{Name: "mixsu", Age: 20}
	fmt.Println(user1.Intro())

	user2 := User{Name: "vince", Age: 19}
	fmt.Println(user2.Age)
	user2.AgePlus()
	fmt.Println(user2.Age)
}

// 值接收者 - 操作对象的副本
// value receiver means this method receives a copy of the User
func (u User) Intro() string {
	return fmt.Sprintf("Hi, I am %s", u.Name)
}

// 指针接收者 - 操作对象本身
// pointer receiver means this method receives a pointer to the User
func (u *User) AgePlus() {
	u.Age++
}
