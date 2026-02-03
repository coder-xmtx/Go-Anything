package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter, r *http.Request) {
	// 写入响应头
	w.Header().Set("Content-Type", "application/json")
	// 写入状态码
	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok":       true,
		"message":  "JSON encode sucessful",
		"datetime": time.Now().UTC(),
	}

	// json编码
	_ = json.NewEncoder(w).Encode(res)
}

func main() {

	http.HandleFunc("/ok", successHandler)

	err := http.ListenAndServe(":5000", nil)

	fmt.Println(err)
}
