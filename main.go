package main

import (
	"encoding/json"
//	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type StatusResponse struct {
	Status string `json:"status"`
	Hostname string `json:"hostname"`
	Version string 	`json:"version"`
	Timestamp string `json:"timestamp"`
}

func main() {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "1.0.0" 
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			hostname = "unknown"
		}

		resp := StatusResponse{
			Status: "running",
			Hostname: hostname,
			Version: version,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		}

		log.Printf("[%s] %s %s - 200 OK\n", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy"}`))
	})

	log.Printf("Starting SRE Prodcution App on port %s (Version: %s)...", port, version)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}




