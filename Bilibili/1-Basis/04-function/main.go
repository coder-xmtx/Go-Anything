package main

import "fmt"

func main() {
	strs := []string{"hello", "world"}
	result := process(strs, func(str string) string {
		return str + "!"
	})
	fmt.Println(result)
}

//
type StringProcessor func(string) string

func process(strs []string, processor StringProcessor) []string {
	var result []string
	for _, str := range strs {
		result = append(result, processor(str))
	}
	return result
}
