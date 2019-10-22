package dto

import "time"

// Temperature is a dataset for temperature measure of a device
type Temperature struct {
	ID        int       `db:"id" json:"id,omitempty"`
	DeviceID  int       `db:"device_id" json:"device_id"`
	Value     float32   `db:"value" json:"value"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// Humidity is a dataset for humidity measure of a device
type Humidity struct {
	ID        int       `db:"id" json:"id,omitempty"`
	DeviceID  int       `db:"device_id" json:"device_id"`
	Value     float32   `db:"value" json:"value"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

// Device is a dataset for a device
type Device struct {
	ID        int       `db:"id" json:"id,omitempty"`
	Name      string    `db:"name" json:"name"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
}
