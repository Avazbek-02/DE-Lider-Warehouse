package main

import (
	"log"

	"github.com/Avazbek-02/DE-Lider-Warehouse/config"
	"github.com/Avazbek-02/DE-Lider-Warehouse/internal/app"
)

func main() {
	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run database migrations
	if err := app.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %s", err)
	}

	// Run app
	app.Run(cfg)
}
