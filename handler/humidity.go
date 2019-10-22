package handler

import (
	"iotsrv/dto"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// HumidityStorager is an interface for a storage of humidity measures
type HumidityStorager interface {
	All(deviceID int64) ([]dto.Humidity, error)
	ByTimePeriod(deviceID int64, from, to time.Time) ([]dto.Humidity, error)
	Create(item dto.Humidity) (int64, error)
}

// Humidity handles a humidity measure requests
type Humidity struct {
	storage HumidityStorager
}

// NewHumidity is a humidity handler costructor
func NewHumidity(storage HumidityStorager) *Humidity {
	return &Humidity{storage}
}

// All handles fetch of humidity measures by device ID
func (h *Humidity) All(ctx echo.Context) error {
	deviceID, err := strconv.ParseInt(ctx.Param("deviceID"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	humidities, err := h.storage.All(deviceID)
	if err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, humidities)
}

// ByTimePeriod handles fetch of humidity measures by device ID and a time period
func (h *Humidity) ByTimePeriod(ctx echo.Context) error {
	deviceID, err := strconv.ParseInt(ctx.Param("deviceID"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	from, err := time.Parse(time.RFC3339, ctx.Param("from"))
	if err != nil {
		return echo.ErrBadRequest
	}
	to, err := time.Parse(time.RFC3339, ctx.Param("to"))
	if err != nil {
		return echo.ErrBadRequest
	}

	humidities, err := h.storage.ByTimePeriod(deviceID, from, to)
	if err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, humidities)
}

// Create handles creation of a humidity measure by returning it's ID
func (h *Humidity) Create(ctx echo.Context) error {
	var humidity dto.Humidity
	err := ctx.Bind(&humidity)
	if err != nil {
		return err
	}

	humidityID, err := h.storage.Create(humidity)
	if err != nil {
		return echo.ErrBadRequest
	}
	data := map[string]int64{"ID": humidityID}
	return ctx.JSON(http.StatusOK, data)
}
