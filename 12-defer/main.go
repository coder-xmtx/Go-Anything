package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("--Case 1: success--")
	if err := doWork(true); err != nil {
		fmt.Println("errors:", err)
	}
	fmt.Println("--Case 2: error--")
	if err := doWork(false); err != nil {
		fmt.Println("errors:", err)
	}
}

func doWork(success bool) error {
	fmt.Println("Start: resource acquired")

	// defer 会在函数执行结束时运行,不管函数是否出错,类似try-finally
	defer fmt.Println("Cleanup: resource released")

	if !success {
		return errors.New("Something went wrong. I am return early.")
	}

	fmt.Println("Work: doing somthing imp")
	fmt.Println("Work: this work is done")

	return nil
}
