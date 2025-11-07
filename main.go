package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"fleet-monitor/internal/handlers"
	"fleet-monitor/internal/storage"
)

func main() {
	// Strictly require .env to load, exit on failure
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	// Read required environment variables (no defaults)
	apiVersion := mustGetEnv("API_VERSION")
	apiPort := mustGetEnv("API_PORT")

	// Load devices list from CSV
	store, err := storage.NewDeviceStore("devices.csv")
	if err != nil {
		log.Fatalf("error loading devices: %v", err)
	}
	fmt.Printf("Loaded %d devices\n", len(store.Devices))

	// Setup router
	r := chi.NewRouter()
	h := &handlers.DeviceHandler{Store: store}

	// Prefix routes with API version
	r.Route(fmt.Sprintf("/api/%s", apiVersion), func(r chi.Router) {
		r.Route("/devices", func(r chi.Router) {
			r.Post("/{device_id}/heartbeat", h.PostHeartbeat)
			r.Get("/{device_id}/stats", h.GetStats)
			r.Post("/{device_id}/stats", h.PostStats)
		})
	})

	// Start server
	addr := ":" + apiPort
	fmt.Printf("Server listening on %s (version: %s)\n", addr, apiVersion)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

// mustGetEnv reads a required environment variable or exits if missing
func mustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("required environment variable %s not set", key)
	}
	return val
}
