package main

import (
	"api-marketplace/config"
	"api-marketplace/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"hello": "world",
		})
	})

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	dbGorm.Ping()

	PembeliRoute := e.Group("/pembeli")
	PembeliRoute.POST("/", controller.CreatePembeli)
	PembeliRoute.GET("/:id", controller.GetPembeli)
	PembeliRoute.PUT("/:id", controller.UpdatePembeli)
	PembeliRoute.DELETE("/:id", controller.DeletePembeli)

	e.Logger.Fatal(e.Start(":8080"))
}
