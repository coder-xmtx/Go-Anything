package main

import (
	"fmt"
	"sync"
	"time"
)

// === go 协程 ===

// func say() {
// 	fmt.Println("这是子协程")
// }

// // 主协程
// func main() {
// 	go say() // 需要模拟延迟来运行子协程，不然这个函数太简单，主协程直接退出了
// 	time.Sleep(time.Second)
// 	fmt.Println("OVER")
// }

// === 模拟订单并发 ===

type Order struct {
	Name string        // 订单名称
	Time time.Duration // 模拟延迟网络的请求
}

func notify(order Order, wg *sync.WaitGroup) {

	// 表示子协程执行完毕，要在函数结束之前执行
	defer wg.Done()

	// 模拟延迟
	time.Sleep(order.Time)
	// 打印日志
	fmt.Printf("Order %s Finish.\n", order.Name)
}

// 主协程
func main() {

	var wg sync.WaitGroup

	orders := []Order{
		{Name: "A001", Time: 100 * time.Millisecond},
		{Name: "A002", Time: 1000 * time.Millisecond},
		{Name: "A003", Time: 3000 * time.Millisecond},
	}

	// 告诉 go 语言一共有几个子协程
	wg.Add(len(orders))

	// 通知用户下完单了
	for _, order := range orders {
		go notify(order, &wg)
	}

	// 让主协程等待
	wg.Wait()

	fmt.Println("主协程完毕")
}
