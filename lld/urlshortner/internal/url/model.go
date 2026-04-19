package url

import "time"

type URL struct {
	Code      string
	Original  string
	ExpiresAt *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
