package storage

import (
	"encoding/csv"
	"fmt"
	"os"

	"fleet-monitor/internal/models"
)

// DeviceStore keeps track of devices in memory
type DeviceStore struct {
	Devices map[string]*models.Device
}

// NewDeviceStore reads devices.csv and populates the store
func NewDeviceStore(csvPath string) (*DeviceStore, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv: %w", err)
	}

	store := &DeviceStore{
		Devices: make(map[string]*models.Device),
	}

	// Skip header if present, load rows
	for i, row := range records {
		if i == 0 && row[0] == "device_id" {
			continue
		}

		if len(row) > 0 {
			id := row[0]
			store.Devices[id] = &models.Device{ID: id}
		}
	}

	return store, nil
}
