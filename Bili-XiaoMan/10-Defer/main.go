package main

import (
	"fmt"
)

// 九九乘法表（程序级别的异常处理）
// 程序级别的异常，后面的代码就不走了，什么时候要用这种异常，比如说数据库连接不上，那后面的代码也没啥意义
func table(a int) (bool, error) {

	if a != 9 {
		// 抛出异常
		panic("参数只能是9")
	}

	for i := 1; i <= a; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d * %d = %d  ", i, k, i*k)
		}
		fmt.Println()
	}

	return true, nil
}

// defer
// 用 defer 修饰的代码会推迟到函数末尾去执行
// 多个 defer 的话，遵循后进先出
// 整体顺序：普通代码先执行 -> return 已确定返回值 -> defer -> return 返回

func main() {

	// defer例子，最后执行结果为 1 5 4 3 2

	// fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// defer fmt.Println(4)
	// fmt.Println(5)

	// -------------------------------------------

	// 捕获异常，这样子输出就没有一大堆报错，同时，最后的 “其余代码” 也没有执行
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	status, err := table(8)
	fmt.Println(status, err)
	// --------------------
	fmt.Println("其余代码")
}
