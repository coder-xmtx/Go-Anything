package main

import (
	"errors"
	"fmt"
)

// 九九乘法表（业务级别的异常处理）
func table(a int) (bool, error) {

	if a != 9 {
		return false, errors.New("参数只能是9") // 错误信息
	}

	for i := 1; i <= a; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("%d * %d = %d  ", i, k, i*k)
		}
		fmt.Println()
	}

	return true, nil
}

func main() {
	status, err := table(8)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
}
