package main

import (
	"fmt"
	"time"
)

func main() {
	// if else 语句
	score := 80
	if score >= 80 {
		fmt.Println("优秀")
	} else if score >= 60 {
		fmt.Println("良好")
	} else {
		fmt.Println("不及格")
	}

	// for 循环
	orders := []string{"A", "B", "C"}
	for index, value := range orders {
		fmt.Println("处理的订单是:", value, "订单号是:", index)
	}

	// for 循环的while语法糖
	i := 0
	for i < 5 {
		i++
		fmt.Println("处理的订单是:", i)
	}

	// switch case 语句
	statusCode := 200
	switch statusCode {
	case 200, 201:
		fmt.Println("OK")
	case 404:
		fmt.Println("Not Found")
	case 500:
		fmt.Println("Internal Server Error")
	default:
		fmt.Println("Unknown")
	}

	// switch case 重写上面的 if else
	score2 := 80
	switch {
	case score2 >= 80:
		fmt.Println("优秀")
	case score2 >= 60:
		fmt.Println("良好")
	default:
		fmt.Println("不及格")
	}

	// select case 语句
	// 与 channel 一起使用，到后面再学习
	ch := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "finish"
	}()
	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
