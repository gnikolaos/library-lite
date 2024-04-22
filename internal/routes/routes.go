package routes

import (
	"github.com/dizars1776/library-lite/internal/handlers"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {
	// Index page accessible to all users
	e.GET("/", handlers.HomeHandler)

	adminGroup := e.Group("/admin")
	//	adminGroup.Use(RoleMiddleware("admin"))
	e.GET("/admin", handlers.Aadmin)

	librarianGroup := e.Group("/librarian")
	//	librarianGroup.Use(RoleMiddleware("admin", "librarian"))
	e.GET("/librarian/books", handlers.Lbooks)
	e.GET("/librarian/patrons", handlers.Lpatrons)

	patron := e.Group("/patron")
	//	patron.Use(RoleMiddleware("admin", "librarian", "patron"))
	e.GET("/patron", handlers.Ppatron)
}
