package main

import (
	"fmt"
	"time"
)

// 子协程 1
func worker1(ch chan string) {
	// 假设延迟50毫秒
	time.Sleep(50 * time.Millisecond)
	ch <- "1st Data"
}

// 子协程 2
func worker2(ch chan string) {
	// 假设延迟30毫秒
	time.Sleep(30 * time.Millisecond)
	ch <- "2nd Data"
}

// 主协程
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go worker1(ch1)
	go worker2(ch2)

	// 现在有两个通道的数据，select只会执行延迟低的通道
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(20 * time.Millisecond): // 超时判断：这里定义超过20毫秒为超时状态
		fmt.Println("超时")
		// 如果不想等通道就写default
		// default:
		// 	fmt.Println("非阻塞测试")
	}
}
