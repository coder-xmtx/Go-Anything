package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CatFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func main() {
	url := "https://catfact.ninja/fact"

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
		fmt.Println("Read body failed", err)
		return
	}

	var data CatFactResponse

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		fmt.Println("Unmarshal failed", err)
		return
	}

	fmt.Println(data.Fact, data.Length)
}
