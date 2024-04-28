package handlers

import (
	"context"
	"os"

	"github.com/dizars1776/library-lite/internal/templates/components/errors"
	"github.com/labstack/echo/v4"
)

func NotFound(c echo.Context) error {
	notFoundIndex := errors.NotFound()
	notFoundIndex.Render(context.Background(), os.Stdout)
	return notFoundIndex.Render(context.Background(), c.Response().Writer)
}
