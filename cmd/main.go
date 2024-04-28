package main

import (
	"log"
	"os"

	"github.com/dizars1776/library-lite/internal/routes"
	"github.com/dizars1776/library-lite/internal/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	dbFile := "lili.db"
	store, err := store.NewStore(dbFile)
	if err != nil {
		log.Fatal("Error on database store creation\n", err)
	}
	defer store.Close()

	// Check if a database file exists
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Print("Database does not exist, seeding...")
		err = store.CreateTables()
		if err != nil {
			log.Fatal("Error on database table creation\n", err)
		}
		err = store.Seed()
		if err != nil {
			log.Fatal("Error on database seed\n", err)
		}
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Static("./static"))

	app := &routes.App{
		Store: store,
	}

	app.RegisterRoutes(e)

	e.Logger.Fatal(e.Start("localhost:8080"))
}
