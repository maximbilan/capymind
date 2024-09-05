package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/capymind/internal/bot"
	"github.com/capymind/internal/scheduler"
)

const port = 8080

func main() {
	log.Println("Starting capymind...")
	log.Printf("Starting server on localhost:%d", port)

	http.HandleFunc("/handler", bot.Parse)
	http.HandleFunc("/schedule", scheduler.Schedule)
	http.HandleFunc("/sendMessage", scheduler.SendMessage)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
