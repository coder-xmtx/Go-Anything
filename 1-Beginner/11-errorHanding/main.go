package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// Golang 不使用异常来处理错误，使用 error 进行错误处理

	if err := run(); err != nil {
		log.Fatal(err) // log.Fatal 会打印错误信息，然后退出
	}
}

func run() error {
	input := "hu"
	level, err := parseLevel(input)
	if err != nil {
		return err
	}
	fmt.Println(level)
	return nil
}

func parseLevel(s string) (int, error) {
	// (value, error)
	// nil 意味着没有错误
	// error != nil 意味着有错误

	n, err := strconv.Atoi(s) // string to int
	if err != nil {
		return 0, fmt.Errorf("level must be a number")
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("level must be between 1 and 5")
	}
	return n, nil
}
