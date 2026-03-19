package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	array := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(array...))
}

// 可变参数

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
