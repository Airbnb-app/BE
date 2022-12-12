package main

import (
	"fmt"
	"net/http"

	"github.com/GP-3-Kelompok-2/airbnb-app-project/config"
	"github.com/GP-3-Kelompok-2/airbnb-app-project/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()

	e := echo.New()
	middlewares.LogMiddlewares(e)
	e.Use(middleware.Logger())
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello world")
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
