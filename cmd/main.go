package main

import (
	"fmt"
	"net/http"

	"github.com/capymind"
)

func main() {
	fmt.Println("Starting capymind...")
	http.HandleFunc("/", capymind.Handler)
	fmt.Println("Starting server on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
