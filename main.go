package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"fleet-monitor/internal/handlers"
	"fleet-monitor/internal/storage"
)

// separate out to handle dynamic versioning in the future (from env var for ex)
const APIVersion = "v1"
const APIPort = "6733"

func main() {
	// Load devices list from csv
	store, err := storage.NewDeviceStore("devices.csv")
	if err != nil {
		log.Fatalf("error loading devices: %v", err)
	}
	fmt.Printf("Loaded %d devices\n", len(store.Devices))

	// routing
	r := chi.NewRouter()
	h := &handlers.DeviceHandler{
		Store: store,
	}

	// api/version route prefix
	r.Route(fmt.Sprintf("/api/%s", APIVersion), func(r chi.Router) {
		// Group all /devices routes
		r.Route("/devices", func(r chi.Router) {
			r.Post("/{device_id}/heartbeat", h.PostHeartbeat)
			r.Get("/{device_id}/stats", h.GetStats)
			r.Post("/{device_id}/stats", h.PostStats)
		})
	})

	//start listening
	fmt.Printf("Server listening on :%s", APIPort)
	http.ListenAndServe(":"+APIPort, r)
}
