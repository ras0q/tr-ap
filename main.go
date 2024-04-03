package main

import (
	"tr-ap/internal/handler"
	"tr-ap/internal/migration"
	"tr-ap/internal/pkg/config"
	"tr-ap/internal/repository"

	ap "github.com/go-ap/activitypub"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// connect to database
	db, err := sqlx.Connect("mysql", config.MySQL().FormatDSN())
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer db.Close()

	// migrate tables
	if err := migration.MigrateTables(db.DB); err != nil {
		e.Logger.Fatal(err)
	}

	// setup repository
	repo := repository.New(db)

	// setup routes
	h := handler.New(ap.IRI("https://tr-ap.trap.show"), repo)
	h.SetupRoutes(e)

	e.Logger.Fatal(e.Start(config.AppAddr()))
}
