package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Status: ", response.Status)
}
