package user

import "lld-ratelimiter/internal/limiter"

type User struct {
	ID      int
	Name    string
	Limiter limiter.RateLimiter
}
