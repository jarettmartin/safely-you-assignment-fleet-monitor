package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"fleet-monitor/internal/models"
	"fleet-monitor/internal/storage"
)

type DeviceHandler struct {
	Store *storage.DeviceStore
}

type HeartbeatPayload struct {
	SentAt time.Time `json:"sent_at"`
}

type StatsPayload struct {
	SentAt     time.Time `json:"sent_at"`
	UploadTime int       `json:"upload_time"`
}

// writeJSONError is a helper for structured error messages per OpenAPI spec.
func writeJSONError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"msg": msg})
}

// PostHeartbeat handles POST /devices/{device_id}/heartbeat
func (h *DeviceHandler) PostHeartbeat(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "device_id")

	device := h.Store.Devices[deviceID]
	if device == nil {
		writeJSONError(w, http.StatusNotFound, "Device not found")
		return
	}

	var payload HeartbeatPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Invalid request body")
		return
	}

	device.AddHeartbeat(payload.SentAt)
	w.WriteHeader(http.StatusNoContent) // 204 per spec
}

// PostStats handles POST /devices/{device_id}/stats
func (h *DeviceHandler) PostStats(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "device_id")

	device := h.Store.Devices[deviceID]
	if device == nil {
		writeJSONError(w, http.StatusNotFound, "Device not found")
		return
	}

	var payload StatsPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Invalid request body")
		return
	}

	upload := models.UploadStat{
		SentAt:     payload.SentAt,
		UploadTime: int(payload.UploadTime),
	}
	device.AddUpload(upload)

	w.WriteHeader(http.StatusNoContent) // 204 per spec
}

// GetStats handles GET /devices/{device_id}/stats
func (h *DeviceHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	deviceID := chi.URLParam(r, "device_id")

	device := h.Store.Devices[deviceID]
	if device == nil {
		writeJSONError(w, http.StatusNotFound, "Device not found")
		return
	}

	// Return 204 if there are no heartbeats or uploads yet
	if len(device.Heartbeats) == 0 && len(device.Uploads) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Convert avg upload time → duration string (e.g. "4.5s")
	avgUploadDuration := time.Duration(device.AvgUploadTime()) * time.Nanosecond

	response := map[string]any{
		"avg_upload_time": avgUploadDuration.String(),
		"uptime":          device.Uptime(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(response)
}
