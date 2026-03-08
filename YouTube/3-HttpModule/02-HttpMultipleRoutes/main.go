package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("welcome try to /hello?name=mixsu"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}

	_, _ = w.Write([]byte("Hello " + name))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Turn to http://localhost:5000/")

	err := http.ListenAndServe(":5000", nil)

	if err != nil {
		fmt.Println(err)
	}
}
