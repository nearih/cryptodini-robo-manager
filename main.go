package main

import (
	"net/http"

	"cryptodini-robo-manager/delivery"
	"cryptodini-robo-manager/service"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// new cryptodiniService service
	cryptodiniService, err := service.NewCryptodiniService()
	if err != nil {
		panic(err)
	}

	// new assetmanger service
	assetManagerService, err := service.NewAssetManagerService(cryptodiniService)
	if err != nil {
		panic(err)
	}

	// new delivery
	handler, err := delivery.NewDelivery(assetManagerService)
	if err != nil {
		panic(err)
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Cryptodini!")
	})
	e.GET("/test", handler.Test())
	e.POST("/deposit", handler.Deposit())
	e.POST("/withdraw", handler.Withdraw())
	e.GET("/port", handler.GetPort())

	e.Logger.Fatal(e.Start(":1323"))
}
