package url

import (
	"fmt"
	"time"
)

type URL struct {
	Code      string
	Original  string
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u URL) String() string {
	expires := "never"
	if !u.ExpiresAt.IsZero() {
		expires = u.ExpiresAt.Format(time.RFC3339)
	}
	return fmt.Sprintf("URL{code=%s, original=%s, expires=%s}", u.Code, u.Original, expires)
}
