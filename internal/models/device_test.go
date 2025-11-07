package models

import (
	"testing"
	"time"
)

func TestUptime(t *testing.T) {
	d := &Device{ID: "1"}
	now := time.Now()

	d.AddHeartbeat(now)
	d.AddHeartbeat(now.Add(time.Minute))
	d.AddHeartbeat(now.Add(2 * time.Minute))

	got := d.Uptime()
	want := 100.0

	if got != want {
		t.Errorf("Uptime() = %v, want %v", got, want)
	}
}
