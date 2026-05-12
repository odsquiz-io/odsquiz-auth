// pkg/database/postgres.go: provides functions for postgres database connections and operations.
package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/pkg/config"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

// NewPostgresConnection establishes a connection to the PostgreSQL database using GORM.
func NewPostgresConnection(cfg *config.Config) (*gorm.DB, error) {
	// Build the database connection string from config
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	// successfull message
	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}

	DB = Dbinstance{
		Db: db,
	}

	return db, nil
}
