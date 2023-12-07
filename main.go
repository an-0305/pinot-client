package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/an-0305/pinot-client/packages/db"
	"github.com/an-0305/pinot-client/routers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()

	db, err := db.Init()
	if err != nil {
		log.Fatalf("failed to initialize a new database: %v", err)
	}

	routers.Init(e, db)

	go func() {
		if err := e.Start(":8000"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
