// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// getEnv gets an environment variable or returns a default value.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// helloServer responds to requests with a greeting message.
func helloServer(w http.ResponseWriter, r *http.Request) {
	// Get the response message from an environment variable, with a default.
	// This allows us to change the message for the deployment/rollback exercise.
	message := getEnv("RESPONSE_TEXT", "Hello from Go!")
	fmt.Fprintf(w, "%s\n", message)
	log.Printf("Served a request to %s", r.RemoteAddr)
}

func main() {
	// Set up the handler for the root path.
	http.HandleFunc("/", helloServer)

	// Determine the port to listen on.
	port := getEnv("PORT", "8080")
	log.Printf("Server starting on port %s...", port)

	// Start the server.
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
