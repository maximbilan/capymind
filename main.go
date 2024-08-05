package main

import (
	"capymind/cloud"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting capymind...")
	http.HandleFunc("/", cloud.Handler)
	fmt.Println("Starting server on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
