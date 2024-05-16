package routes

import (
	"github.com/dizars1776/library-lite/internal/auth"
	"github.com/dizars1776/library-lite/internal/handlers"
	"github.com/dizars1776/library-lite/internal/store"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type App struct {
	Store *store.Store
}

func (app *App) RegisterRoutes(e *echo.Echo) {

	// Index page accessible to all users
	e.GET("/", handlers.Home)

	// AUTH ROUTES
	authHandler := handlers.NewAuthHandler(app.Store)
	authGroup := e.Group("/auth")
	authGroup.GET("/login", authHandler.GetLogin)
	authGroup.GET("/signup", authHandler.GetSingup)

	authGroup.POST("/login", authHandler.PostLogin)

	// ADMIN ROUTES
	adminHandler := handlers.NewAdminHandler(app.Store)
	adminGroup := e.Group("/admin")
	// Guarded group
	adminGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:   []byte(auth.GetJWTSecret()),
		TokenLookup:  "cookie:access-token", // "<source>:<name>"
		ErrorHandler: auth.JWTErrorHandler,
	}))
	adminGroup.GET("", adminHandler.Dashboard)
	adminGroup.GET("/books", adminHandler.Books)
	adminGroup.GET("/users", adminHandler.Users)

	// LIBRARIAN ROUTES
	librarianHandler := handlers.NewLibrarianHandler(app.Store)
	librarianGroup := e.Group("/librarian")
	//librarianGroup.Use(auth.Middleware())
	librarianGroup.GET("/librarian/books", librarianHandler.Books)
	librarianGroup.GET("/librarian/patrons", librarianHandler.Patrons)

	// PATRON ROUTES
	patronHandler := handlers.NewPatronHandler(app.Store)
	patronGroup := e.Group("/patron")
	//patronGroup.Use(auth.Middleware())
	patronGroup.GET("/patron", patronHandler.Dashboard)

	// MISC
	e.GET("/*", handlers.NotFound)
}
