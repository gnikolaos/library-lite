package handlers

import (
	"context"
	"os"

	"github.com/dizars1776/library-lite/internal/templates"
	"github.com/dizars1776/library-lite/internal/templates/components"
	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	welcomeComp := components.Welcome()
	index := templates.Index(welcomeComp, "Library Lite - lIlI")
	index.Render(context.Background(), os.Stdout)
	return index.Render(context.Background(), c.Response().Writer)
}
