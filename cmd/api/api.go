package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/kauanpecanha/odsquiz-auth/internal/migrations"
	"github.com/kauanpecanha/odsquiz-auth/internal/routes"
	"github.com/kauanpecanha/odsquiz-auth/pkg/config"
	"github.com/kauanpecanha/odsquiz-auth/pkg/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	
	startupCtx, cancel := context.WithTimeout(context.Background(), database.StartupConnectionTimeout)
	defer cancel()

	err = migrations.RunMigrations(startupCtx, cfg)
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}

	app := fiber.New(fiber.Config{
		AppName: "ODS Quiz Auth Microservice",
	})

	app.Use(cors.New())

	routes.Setup(app)
	
	log.Fatal(app.Listen(":" + cfg.Port))
}
