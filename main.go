package main

import (
	"log"

	"github.com/hyuti/logger-project/config"
	"github.com/hyuti/logger-project/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
