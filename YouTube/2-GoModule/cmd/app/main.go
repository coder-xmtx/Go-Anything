package main

import (
	"fmt"
	"go_modules/internal/greet"
)

func main() {
	msg1 := greet.Hello("mixsu")
	fmt.Println(msg1)
}
