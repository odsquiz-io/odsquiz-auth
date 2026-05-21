// internal/migrations/api.go: Implement automigration usage outside database package to be used whenever requested
package migrations

import (
	"log"

	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/pkg/config"
	"github.com/kauanpecanha/odsquiz-auth/pkg/database"
)

func RunMigrations() error {
	// Load configuration settings
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Establish connection to PostgreSQL database
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Automigrate user schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	return err
}