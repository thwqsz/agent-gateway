package main

import (
	"log"
	"net/http"

	"github.com/thwqsz/agent-gateway/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ask", handlers.AskHandler)

	log.Println("Go-шлюз запущен на :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
