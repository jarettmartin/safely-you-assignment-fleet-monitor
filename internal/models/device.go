package models

import (
	"math"
	"time"
)

// Device represents one device from devices.csv
type Device struct {
	ID         string `csv:"device_id"`
	Heartbeats []time.Time
	Uploads    []UploadStat
}

type UploadStat struct {
	SentAt     time.Time
	UploadTime int
}

func (d *Device) AddHeartbeat(t time.Time) {
	d.Heartbeats = append(d.Heartbeats, t)
}

func (d *Device) AddUpload(upload UploadStat) {
	d.Uploads = append(d.Uploads, upload)
}

// Uptime calculates the percent of minutes where the device was online.
// (heartbeats divided by total minutes between first and last heartbeat)
func (d *Device) Uptime() float64 {
	n := len(d.Heartbeats)
	if n < 2 {
		return 0 // not enough data to compute uptime
	}

	first := d.Heartbeats[0]
	last := d.Heartbeats[n-1]
	totalMinutes := int(math.Round(last.Sub(first).Minutes()))
	if totalMinutes == 0 {
		return 100 // only one-minute window, assume 100% uptime
	}

	uptime := (float64(n) / float64(totalMinutes)) * 100
	if uptime > 100 {
		uptime = 100 // cap it at 100%
	}
	return uptime
}

// AvgUploadTime computes the average upload time across all upload stats.
func (d *Device) AvgUploadTime() float64 {
	if len(d.Uploads) == 0 {
		return 0
	}

	var total int
	for _, u := range d.Uploads {
		total += u.UploadTime
	}
	return float64(total) / float64(len(d.Uploads))
}
