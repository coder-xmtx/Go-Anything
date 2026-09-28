// 互斥锁

package main

import (
	"fmt"
	"sync"
)

// 子协程
func worker(menu *[]string, str string, wg *sync.WaitGroup, mu *sync.Mutex) {

	defer wg.Done() // 最后执行

	mu.Lock()

	// === 临界区 ===
	*menu = append(*menu, str)
	// === 临界区 ===

	mu.Unlock()
}

// 主协程
func main() {

	var wg sync.WaitGroup

	var mu sync.Mutex // 互斥锁，同一时间内只能有一个协程去操作临界区的共享数据

	wg.Add(3)

	menu := []string{"Python", "C++", "JavaScript"}

	// 这三个子协程都要操作 menu 切片，没有顺序，可能会同时覆盖，需要用到互斥锁
	go worker(&menu, "Golang", &wg, &mu)
	go worker(&menu, "TypeScript", &wg, &mu)
	go worker(&menu, "HTML", &wg, &mu)

	wg.Wait()

	fmt.Println(menu)
}
