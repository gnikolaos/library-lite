package handlers

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/dizars1776/library-lite/internal/store"
	"github.com/dizars1776/library-lite/internal/templates"
	"github.com/dizars1776/library-lite/internal/templates/components/librarian"
)

type LibrarianHandler struct {
	store *store.Store
}

func NewLibrarianHandler(store *store.Store) *LibrarianHandler {
	return &LibrarianHandler{store: store}
}

func (h *LibrarianHandler) Dashboard(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Librarians index test page.")
}

func (h *LibrarianHandler) Patrons(c echo.Context) error {
	patronsComp := librarian.Patrons()
	dashboardComp := templates.LibrarianDashboard(patronsComp)
	librarianIndex := templates.Index(dashboardComp, "Librarian Dashboard - lIlI")
	librarianIndex.Render(context.Background(), os.Stdout)
	return librarianIndex.Render(context.Background(), c.Response().Writer)
}

func (h *LibrarianHandler) Books(c echo.Context) error {
	booksComp := librarian.Books()
	dashboardComp := templates.LibrarianDashboard(booksComp)
	librarianIndex := templates.Index(dashboardComp, "Librarian Dashboard - lIlI")
	librarianIndex.Render(context.Background(), os.Stdout)
	return librarianIndex.Render(context.Background(), c.Response().Writer)
}
