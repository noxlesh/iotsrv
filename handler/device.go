package handler

import (
	"fmt"
	"iotsrv/dto"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// DeviceStorager interface for a storage of device info
type DeviceStorager interface {
	All() ([]dto.Device, error)
	Create(dto.Device) (int64, error)
	Remove(deviceID int64) error
}

// Device handles http requests to manipulate registered devices
type Device struct {
	storage DeviceStorager
}

// NewDevice creates new Device handler
func NewDevice(storage DeviceStorager) *Device {
	return &Device{storage}
}

// All returns all registered devices
func (h *Device) All(ctx echo.Context) error {
	devices, err := h.storage.All()
	if err != nil {
		return echo.ErrBadRequest
	}
	return ctx.JSON(http.StatusOK, devices)
}

// Create registers new device
func (h *Device) Create(ctx echo.Context) error {
	var device dto.Device
	err := ctx.Bind(&device)
	if err != nil {
		return echo.ErrBadRequest
	}
	deviceID, err := h.storage.Create(device)
	if err != nil {
		fmt.Print(err)
		return echo.ErrBadRequest
	}
	data := map[string]int64{"ID": deviceID}
	return ctx.JSON(http.StatusOK, data)
}

// Remove removes a registered device
func (h *Device) Remove(ctx echo.Context) error {
	deviceIDStr := ctx.Param("deviceID")
	deviceID, err := strconv.ParseInt(deviceIDStr, 10, 64)
	if err != nil {
		return echo.ErrBadRequest
	}
	h.storage.Remove(deviceID)
	return ctx.NoContent(http.StatusOK)
}
