package data

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dalmow/sdalm/internal/config"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func RunMigrations(lc fx.Lifecycle, c *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			db, err := sql.Open("postgres", c.DatabaseUrl)
			if err != nil {
				return fmt.Errorf("error while get database connection when migrate: %v", err)
			}

			driver, err := postgres.WithInstance(db, &postgres.Config{})
			if err != nil {
				return fmt.Errorf("error while get driver when migrate: %v", err)
			}
			wd, _ := os.Getwd()
			m, err := migrate.NewWithDatabaseInstance(
				fmt.Sprintf("file://%s", filepath.Join(wd, "migrations")),
				"postgres", driver)
			if err != nil {
				return fmt.Errorf("error while get migrate instance when migrate: %v", err)
			}
			defer m.Close()
			m.Up()

			return nil
		},
	})
}
