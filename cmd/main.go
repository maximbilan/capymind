package main

import (
	"fmt"
	"net/http"

	"github.com/capymind/internal/bot"
)

func main() {
	fmt.Println("Starting capymind...")
	http.HandleFunc("/", bot.Parse)
	fmt.Println("Starting server on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
