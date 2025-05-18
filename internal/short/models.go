package short

import "time"

type Short struct {
	ID          string
	Alias       *string
	OriginalURL string
	ExpiresAt   *time.Time
	CreatedAt   time.Time
}
