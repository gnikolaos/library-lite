package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
	"os"

	"github.com/dizars1776/library-lite/internal/templates"
)

func Index(c echo.Context) error {
	loginTempl := templates.AuthIndex("Login - lIlI")
	loginTempl.Render(context.Background(), os.Stdout)
	return loginTempl.Render(context.Background(), c.Response().Writer)
}
