package user

import "lld-ratelimiter/internal/limiter"

func New(id int, name string, limiter limiter.RateLimiter) *User {
	return &User{
		ID:      id,
		Name:    name,
		Limiter: limiter,
	}
}

func (u *User) Allow() bool {
	return u.Limiter.Allow()
}
