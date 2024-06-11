package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/voikin/devan-distribution/internal/application"
)

func waitQuitSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func main() {
	app, err := application.New()
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}

	waitQuitSignal()

	if err := app.Stop(); err != nil {
		log.Fatalf("Failed to stop app gracefully: %v", err)
	}
}
