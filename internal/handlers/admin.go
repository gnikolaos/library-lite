package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/dizars1776/library-lite/internal/store"
	"github.com/dizars1776/library-lite/internal/templates"
	"github.com/dizars1776/library-lite/internal/templates/admin"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	store *store.Store
}

func NewAdminHandler(store *store.Store) *AdminHandler {
	return &AdminHandler{store: store}
}

func (h *AdminHandler) Dashboard(c echo.Context) error {
	adminDash := admin.AdminDashboard()
	authIndex := templates.AuthIndex(adminDash, "Dashboard - Admin")
	authIndex.Render(context.Background(), os.Stdout)
	return authIndex.Render(context.Background(), c.Response().Writer)
}

func (h *AdminHandler) Books(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Books admin test page.")
}

func (h *AdminHandler) Users(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Users admin test page.")
}
