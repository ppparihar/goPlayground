package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	nodeName := os.Getenv("name")
	if nodeName == "" {
		nodeName = "Default Node"
	}
	log.Printf("Node name %s\n", nodeName)
	// Get the port from environment variable or use default (8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	// Define the HTTP handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Set the response content type
		w.Header().Set("Content-Type", "text/plain")

		// Write the response body
		fmt.Fprintf(w, "Hello, World! %s\n", nodeName)
	}

	// Register the handler function to handle all requests
	http.HandleFunc("/", handler)

	// Start the HTTP server
	addr := ":" + port
	log.Printf("Starting HTTP server on %s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("HTTP server error:", err)
	}
}
