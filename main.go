package main

import (
	"log"
	"net/http"
	"receipt-service/handlers"
)

func main() {
	// Initialize Memcached
	handlers.InitMemcached()

	// Set up routes
	http.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler)
	http.HandleFunc("/receipts/", handlers.GetPointsHandler)

	// Log that the server is starting
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
