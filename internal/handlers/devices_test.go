package handlers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"

	"fleet-monitor/internal/models"
	"fleet-monitor/internal/storage"
)

func TestPostHeartbeat(t *testing.T) {
	store := &storage.DeviceStore{
		Devices: map[string]*models.Device{
			"abc": {ID: "abc"},
		},
	}
	h := &DeviceHandler{Store: store}

	payload := `{"sent_at": "` + time.Now().Format(time.RFC3339) + `"}`

	req := httptest.NewRequest(http.MethodPost, "/devices/abc/heartbeat", bytes.NewBufferString(payload))

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("device_id", "abc")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	w := httptest.NewRecorder()
	h.PostHeartbeat(w, req)

	res := w.Result()
	if res.StatusCode != http.StatusNoContent {
		t.Errorf("expected 204, got %d", res.StatusCode)
	}
}
