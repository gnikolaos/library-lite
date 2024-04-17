package handlers

import (
	"context"
	"os"

	"github.com/dizars1776/library-lite/internal/templates"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	helloComponent := templates.Hello("Guest User")
	indexComponent := templates.Index(helloComponent, "Library Lite - lIlI")
	indexComponent.Render(context.Background(), os.Stdout)
	return indexComponent.Render(context.Background(), c.Response().Writer)
}
