package handler

import (
	"iotsrv/dto"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// TemperatureStorager interface for a storage
// of temperature measures
type TemperatureStorager interface {
	All(deviceID int64) ([]dto.Temperature, error)
	ByTimePeriod(deviceID int64, from, to time.Time) ([]dto.Temperature, error)
	Create(item dto.Temperature) (int64, error)
}

// Temperature handles a temperature measure requests
type Temperature struct {
	storage TemperatureStorager
}

// NewTemperature is a temperature handler costructor
func NewTemperature(storage TemperatureStorager) *Temperature {
	return &Temperature{storage}
}

// All handles fetch of temperature measures by device ID
func (h *Temperature) All(ctx echo.Context) error {
	deviceID, err := strconv.ParseInt(ctx.Param("deviceID"), 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}

	temperatures, err := h.storage.All(deviceID)
	if err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, temperatures)
}

// ByTimePeriod handles fetch of temperature measures
// by device ID and a time period
func (h *Temperature) ByTimePeriod(ctx echo.Context) error {
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

	temperatures, err := h.storage.ByTimePeriod(deviceID, from, to)
	if err != nil {
		ctx.Logger().Error(err)
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, temperatures)
}

// Create handles creation of a temperature measure
// by returning it's ID
func (h *Temperature) Create(ctx echo.Context) error {
	var t dto.Temperature
	err := ctx.Bind(&t)
	if err != nil {
		ctx.Logger().Debugf("Failed to parse Temperature: %v", err)
		return echo.ErrBadRequest
	}

	temperatureID, err := h.storage.Create(t)
	if err != nil {
		ctx.Logger().Debugf("Failed to create Temperature: %v", err)
		return echo.ErrBadRequest
	}

	data := map[string]int64{"ID": temperatureID}
	return ctx.JSON(http.StatusOK, data)
}
