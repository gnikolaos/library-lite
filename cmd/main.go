package main

import (
	"github.com/dizars1776/library-lite/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("./static"))

	routes.RegisterRoutes(e)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
