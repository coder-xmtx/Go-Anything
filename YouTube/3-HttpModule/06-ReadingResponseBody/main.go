package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// 读取响应体
	url := "https://jsonplaceholder.typicode.com/todos"

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println(response.Status)
		return
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	bodyText := string(bodyBytes)

	max := 250

	if len(bodyText) > max {
		bodyText = bodyText[:max] + "..."
	}

	fmt.Println(bodyText)

}
