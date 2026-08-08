package migrations

import (
	"context"
	"fmt"

	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/pkg/config"
	"github.com/kauanpecanha/odsquiz-auth/pkg/database"
)

func RunMigrations(ctx context.Context, cfg *config.Config) error {
	db, err := database.NewPostgresConnection(ctx, cfg)
	if err != nil {
		return fmt.Errorf("connect to database: %w", err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		return fmt.Errorf("migrate users: %w", err)
	}

	return nil
}
