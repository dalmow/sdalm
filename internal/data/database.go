package data

import (
	"context"
	"fmt"
	"time"

	"github.com/dalmow/sdalm/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/fx"
)

type Database struct {
	Pool *pgxpool.Pool
}

func NewDatabaseConnection(c *config.Config) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	pool, err := pgxpool.New(ctx, c.DatabaseUrl)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to database: %v", err)
	}

	return &Database{
		Pool: pool,
	}, nil
}

func CloseDatabaseConnection(lc fx.Lifecycle, db *Database) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			db.Pool.Close()
			return nil
		},
	})
}
