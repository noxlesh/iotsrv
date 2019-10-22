package storage

import (
	"errors"
	"iotsrv/dto"

	"github.com/jmoiron/sqlx"
)

// Device is a storage for registered devices
type Device struct {
	db *sqlx.DB
}

// NewDevice creates new instance of Device storage
func NewDevice(db *sqlx.DB) (*Device, error) {
	if db == nil {
		return nil, errors.New("failed to create Device stirage: db is nil")
	}
	return &Device{db}, nil
}

// All gets all devices from the storage
func (s *Device) All() ([]dto.Device, error) {
	var devices []dto.Device
	err := s.db.Select(&devices, "select * from device")
	return devices, err
}

// Create creates a new device
func (s *Device) Create(device dto.Device) (int64, error) {
	res, err := s.db.Exec(
		"insert into device (name, created_at) values (?, curdate())",
		device.Name,
	)
	if err != nil {
		return -1, err
	}
	return res.LastInsertId()
}

// Remove deletes a registered device
func (s *Device) Remove(deviceID int64) error {
	_, err := s.db.Exec(
		"DELETE FROM device WHERE id = ?",
		deviceID,
	)
	return err
}
