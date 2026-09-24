package main

import (
	"fmt"
)

// 子协程
func worker(ch chan int) {
	// 子协程发送数据给主协程
	ch <- 100
	ch <- 200
	ch <- 300
	close(ch)
}

// 主协程
func main() {
	// 1. 无缓冲通道：定义channel的时候，make函数不写容量
	// 无缓冲需要发送和接收都准备好，如果只有发送没有接收，代码会阻塞
	ch := make(chan int) // 传输 int
	go worker(ch)

	// 接收子协程传输的数据
	// 数据比较少，可以用变量一个一个接
	// value1 := <-ch
	// value2 := <-ch
	// value3 := <-ch
	// fmt.Println(value1, value2, value3)

	// 数据量很多，可以使用for循环
	// 注意，for循环会一直等待这个通道传输数据，并没有关闭，需要close(ch)
	for value := range ch {
		fmt.Println(value)
	}

	// 2. 有缓冲通道：定义channel的时候，make定义了容量
	// 像上述的worker函数，如果主协程写的是make(chan int, 3)，就不会发生阻塞了
}
