// 读写锁

package main

import (
	"fmt"
	"sync"
	"time"
)

// 子协程 - 读
func read(menu *[]string, wg *sync.WaitGroup, mu *sync.RWMutex) {

	defer wg.Done() // 最后执行

	mu.Lock()

	time.Sleep(50 * time.Millisecond) // 模拟延迟

	fmt.Println(*menu)

	mu.Unlock()
}

// 子协程 - 写
func write(menu *[]string, str string, wg *sync.WaitGroup, mu *sync.RWMutex) {

	// 人为故意延迟 1 毫秒，让读锁先触发
	time.Sleep(1 * time.Millisecond)

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

	var mu sync.RWMutex // 读写锁，读和读共享，读和写互斥

	wg.Add(3)

	menu := []string{"Python", "C++", "JavaScript"}

	go read(&menu, &wg, &mu)
	go read(&menu, &wg, &mu)

	go write(&menu, "Golang", &wg, &mu)

	wg.Wait()

	fmt.Println(menu)
}
