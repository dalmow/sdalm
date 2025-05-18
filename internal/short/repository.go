package short

import (
	"context"

	"github.com/dalmow/sdalm/internal/data"
)

type ShortsRepository interface {
	AlreadyExists(ctx context.Context, id string) (bool, error)
	Create(ctx context.Context, s *Short) error
}

type shortsRepository struct {
	DB *data.Database
}

func NewShortsRepository(db *data.Database) ShortsRepository {
	return &shortsRepository{
		DB: db,
	}
}

func (r *shortsRepository) AlreadyExists(ctx context.Context, id string) (bool, error) {
	const query = `select exists(select 1 from shorts where short_id = $1)`
	var exists bool
	err := r.DB.Pool.QueryRow(ctx, query, id).Scan(&exists)
	return exists, err
}

func (r *shortsRepository) Create(ctx context.Context, s *Short) error {
	const query = `
		insert into shorts(short_id, alias, original_url, expires_at)
		values($1, $2, $3, $4)`

	_, err := r.DB.Pool.Exec(
		ctx,
		query,
		s.ID,
		s.Alias,
		s.OriginalURL,
		s.ExpiresAt)

	return err
}
