package handlers

import (
	"net/http"

	"github.com/dizars1776/library-lite/internal/store"
	"github.com/labstack/echo/v4"
)

type PatronHandler struct {
	store *store.Store
}

func NewPatronHandler(store *store.Store) *PatronHandler {
	return &PatronHandler{store: store}
}

func (h *PatronHandler) Dashboard(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Patrons Dashboard.")
}
