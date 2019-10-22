package storage

import (
	"iotsrv/dto"
	"time"

	"github.com/jmoiron/sqlx"
)

// Temperature stores temperature measurements of devices
type Temperature struct {
	db *sqlx.DB
}

// NewTemperature is a constructor of the Temperature measurements storage
func NewTemperature(db *sqlx.DB) *Temperature {
	return &Temperature{db}
}

// All gets every temperature measuremests from the storage
func (s *Temperature) All(deviceID int64) ([]dto.Temperature, error) {
	var temperatures []dto.Temperature
	err := s.db.Select(
		&temperatures,
		"select * from temperature where device_id=?",
		deviceID,
	)
	if err != nil {
		return nil, err
	}
	return temperatures, nil
}

// ByTimePeriod gets temperature measuremests from the storage by a time period
func (s *Temperature) ByTimePeriod(deviceID int64, from, to time.Time) ([]dto.Temperature, error) {
	var temperatures []dto.Temperature
	err := s.db.Select(
		&temperatures,
		"select * from temperature where id=? and created_at between ? and ?",
		deviceID,
		from,
		to,
	)
	if err != nil {
		return nil, err
	}
	return temperatures, nil
}

// Create a temperature measurement in the storage
func (s *Temperature) Create(item dto.Temperature) (int64, error) {
	res, err := s.db.Exec(
		"insert into temperature (device_id, value, created_at) values (?, ?, ?)",
		item.DeviceID,
		item.Value,
		item.CreatedAt.Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
