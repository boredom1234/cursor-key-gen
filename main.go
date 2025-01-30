package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

type Telemetry struct {
	MachineID    string `json:"telemetry.machineId"`
	MacMachineID string `json:"telemetry.macMachineId"`
	DevDeviceID  string `json:"telemetry.devDeviceId"`
	SQMID        string `json:"telemetry.sqmId"`
	LastModified string `json:"lastModified"`
	Version      string `json:"version"`
}

func generateCursorKey() Telemetry {
	id := uuid.New().String()
	log.Printf("Generated new UUID: %s", id)
	return Telemetry{
		MachineID:    id,
		MacMachineID: id,
		DevDeviceID:  id,
		SQMID:        id,
		LastModified: time.Now().UTC().Format(time.RFC3339),
		Version:      "1.0.1",
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to generate telemetry key")
	key := generateCursorKey()
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(key); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	log.Println("Successfully sent telemetry key response")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback to 8080 for local development
	}
	log.Printf("Starting server on port %s", port)
	http.HandleFunc("/generate", handler)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
