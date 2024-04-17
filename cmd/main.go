package main

import (
	"github.com/dizars1776/library-lite/internal/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("./static"))

	e.GET("/", handlers.Index)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
