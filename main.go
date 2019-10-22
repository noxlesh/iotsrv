package main

import (
	"flag"
	"iotsrv/handler"
	"iotsrv/storage"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// cli
	addr := flag.String("addr", "localhost:8080", "address and port on which the service will run")
	dsn := flag.String("dsn", "iot:iot@/iot?parseTime=true", "DSN used to communicate with a database")
	flag.Parse()

	// db
	db, err := sqlx.Open("mysql", *dsn)
	dieOnErr(err)

	// storages
	temeratureStorage := storage.NewTemperature(db)
	dieOnErr(err)
	humidityStorage := storage.NewHumidity(db)
	dieOnErr(err)
	deviceStorage, err := storage.NewDevice(db)
	dieOnErr(err)

	// handlers
	deviceHandler := handler.NewDevice(deviceStorage)
	dieOnErr(err)
	temperatureHandler := handler.NewTemperature(temeratureStorage)
	dieOnErr(err)
	humidityHandler := handler.NewHumidity(humidityStorage)
	dieOnErr(err)

	// http
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	api := e.Group("/api")
	{
		dev := api.Group("/device")
		dev.GET("/", deviceHandler.All)
		dev.POST("/", deviceHandler.Create)
		dev.DELETE("/:deviceID", deviceHandler.Remove)

		temp := api.Group("/temperature")
		temp.GET("/:deviceID", temperatureHandler.All)
		temp.GET("/:deviceID/:from/:to", temperatureHandler.ByTimePeriod)
		temp.POST("/", temperatureHandler.Create)

		hum := api.Group("/humidity")
		hum.GET("/:deviceID", humidityHandler.All)
		hum.GET("/:deviceID/:from/:to", humidityHandler.ByTimePeriod)
		hum.POST("/", humidityHandler.Create)
	}
	dieOnErr(e.Start(*addr))
}

func dieOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
