package main

import "fmt"

func main() {

	// 定义固定长度的数组
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println(arr[0], arr[1], arr[2])

	// 通过类型推断定义固定长度的数组
	slice := [5]int{1, 2, 3, 4, 5}
	fmt.Println(slice[0], slice[1], slice[2])

	// 定义动态长度的数组
	arr2 := []string{"hello", "world", "mixsu"}
	fmt.Println(arr2[0], arr2[1], arr2[2])

	// 定义空数组及追加元素
	nums := []int{}
	nums = append(nums, 1, 2, 3, 4, 5)
	fmt.Println(nums)

	// make函数创建自定义长度和容量的切片 make(type, length, capacity)
	// 返回值 nums2 是一个初始化后的切片，带有 5 个默认值为 0 的 int 元素
	// 超出容量时会自动扩展，扩展后的容量为原来的两倍
	nums2 := make([]int, 5, 10)
	fmt.Println(nums2, len(nums2), cap(nums2))
	nums2 = append(nums2, 1, 2, 3, 4, 5)
	fmt.Println(nums2, len(nums2), cap(nums2))
	nums2 = append(nums2, 100)
	fmt.Println(nums2, len(nums2), cap(nums2))

	// 将一个切片追加到另一个切片中 append(slice1, slice2...)
	todos := []string{"make a video", "make a blog post", "sleep", "repeat"}
	more := []string{"learn golang"}
	todos = append(todos, more...)

	// for 循环
	views := []int{10, 20, 30, 40, 50}
	total := 0
	for index, value := range views {
		fmt.Println("day", index, "views", value)
		total += value
	}
	fmt.Println("total views", total)
}
