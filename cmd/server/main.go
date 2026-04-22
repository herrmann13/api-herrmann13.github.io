package main

import (
	"log"
	"time"
	"net/http"
	"api-herrmann13.github.io/internal/system"
)

func main() {
	mux := http.NewServeMux()

	startedAt := time.Now()
	mux.HandleFunc("/health", system.HealthHandler(startedAt))
	mux.HandleFunc("/projects", system.ProjectsHandler())

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
