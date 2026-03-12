package main

import "fmt"

type User struct {
	Id      int
	Name    string
	Balence float64
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
	fmt.Println("user1:", user2)

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
