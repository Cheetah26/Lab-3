package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	logger := log.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := "0.0.0.0:" + port

	defaultMessage := os.Getenv("DEFAULT_MESSAGE")
	if defaultMessage == "" {
		defaultMessage = "Hello!"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(defaultMessage))
	})

	http.HandleFunc("/echo/{message}", func(w http.ResponseWriter, r *http.Request) {
		response := "Echo: " + r.PathValue("message")
		w.Write([]byte(response))
	})

	logger.Println("Listening on " + addr)
	logger.Fatal(http.ListenAndServe(addr, nil))
}
