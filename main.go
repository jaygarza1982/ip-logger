package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var expectedAPIKey = os.Getenv("API_KEY")

func logIPHandler(w http.ResponseWriter, r *http.Request) {
	// Only accept POST requests, return 404 if method is not POST
	if r.Method != http.MethodPost {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// If the API key from env vars is not set, always return 404
	if expectedAPIKey == "" {
		http.Error(w, "Not Found", http.StatusNotFound)
		log.Println("API_KEY environment variable is not set")
		return
	}

	// Check for API key, return 404 if invalid
	apiKey := r.Header.Get("API-Key")
	if apiKey != expectedAPIKey {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Get client IP address without port
	ip := r.RemoteAddr
	if colonPos := strings.LastIndex(ip, ":"); colonPos != -1 {
		ip = ip[:colonPos]
	}

	// Log IP address
	fmt.Printf("Request received at %s from %s\n", time.Now().String(), ip)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", ip)
}

func main() {
	http.HandleFunc("/logip", logIPHandler)
	fmt.Println("Server is listening on port 4321...")
	log.Fatal(http.ListenAndServe(":4321", nil))
}
