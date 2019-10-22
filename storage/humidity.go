package storage

import (
	"iotsrv/dto"
	"time"

	"github.com/jmoiron/sqlx"
)

// Humidity stores humidity measurements of devices
type Humidity struct {
	db *sqlx.DB
}

// NewHumidity creates new instance of Humidity storage
func NewHumidity(db *sqlx.DB) *Humidity {
	return &Humidity{db}
}

// All gets every humidity measuremests from the storage
func (s *Humidity) All(deviceID int64) ([]dto.Humidity, error) {
	var humiditys []dto.Humidity
	err := s.db.Select(
		&humiditys,
		"select * from humidity where device_id=?",
		deviceID,
	)
	if err != nil {
		return nil, err
	}
	return humiditys, nil
}

// ByTimePeriod gets humidity measuremests from the storage by a time period
func (s *Humidity) ByTimePeriod(deviceID int64, from, to time.Time) ([]dto.Humidity, error) {
	var humiditys []dto.Humidity
	err := s.db.Select(
		&humiditys,
		"select * from humidity where id=? and created_at between ? and ?",
		deviceID,
		from,
		to,
	)
	if err != nil {
		return nil, err
	}
	return humiditys, nil
}

// Create a humidity measurement in the storage
func (s *Humidity) Create(item dto.Humidity) (int64, error) {
	res, err := s.db.Exec(
		"insert into humidity (device_id, value, created_at) values (?, ?, ?)",
		item.DeviceID,
		item.Value,
		item.CreatedAt.Format("2006-01-02 15:04:05"),
	)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
