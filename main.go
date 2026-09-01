package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	log.Printf("[%s] %s %s from %s", timestamp, r.Method, r.RequestURI, r.RemoteAddr)

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("X-Server-Name", "STUPHHP")
	w.Header().Set("X-Timestamp", timestamp)

	fmt.Fprintf(w, "Server: STUPHHP\n")
	fmt.Fprintf(w, "Timestamp: %s\n", timestamp)
	fmt.Fprintf(w, "Request: %s %s\n", r.Method, r.RequestURI)
	fmt.Fprintf(w, "Headers:\n")

	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "%s: %s\n", name, value)
		}
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("[%s] stupttp server starting on port %s", time.Now().Format("2006-01-02 15:04:05"), port)

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
