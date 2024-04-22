package handlers

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/dizars1776/library-lite/internal/templates"
	"github.com/dizars1776/library-lite/internal/templates/components/librarian"
)

func Lpatrons(c echo.Context) error {
	patronsComp := librarian.Patrons()
	dashboardComp := templates.LibrarianDashboard(patronsComp)
	librarianIndex := templates.Index(dashboardComp, "Librarian Dashboard - lIlI")
	librarianIndex.Render(context.Background(), os.Stdout)
	return librarianIndex.Render(context.Background(), c.Response().Writer)
}

func Lbooks(c echo.Context) error {
	booksComp := librarian.Books()
	dashboardComp := templates.LibrarianDashboard(booksComp)
	librarianIndex := templates.Index(dashboardComp, "Librarian Dashboard - lIlI")
	librarianIndex.Render(context.Background(), os.Stdout)
	return librarianIndex.Render(context.Background(), c.Response().Writer)
}
