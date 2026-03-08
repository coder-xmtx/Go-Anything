package main

import "fmt"

func main() {

	// 例子：判断此人是否成年
	personAge := 16
	isAdult := personAge >= 18
	fmt.Println(isAdult)

	// 例子：判断用户是否有删除帖子的权限
	isAdmin := true
	isUser := false
	canDelete := isAdmin || isUser
	fmt.Println(canDelete)
}
