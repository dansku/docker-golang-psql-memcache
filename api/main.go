package main

import (
	db "example/database"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func main() {

	// Start pretty logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db.Start()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
