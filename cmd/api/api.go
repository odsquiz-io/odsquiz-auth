// cmd/api/api.go: Entry point for the ODS Quiz Auth API microservice.
package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/kauanpecanha/odsquiz-auth/internal/routes"
	"github.com/kauanpecanha/odsquiz-auth/pkg/config"
	"github.com/kauanpecanha/odsquiz-auth/pkg/database"
)

func main() {
	// Load configuration settings
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Establish connection to PostgreSQL database
	_, err = database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	app := fiber.New(fiber.Config{
		AppName: "ODS Quiz Auth Microservice",
	})

	app.Use(cors.New())

	routes.Setup(app)
	
	log.Fatal(app.Listen(":" + cfg.Port))
}
