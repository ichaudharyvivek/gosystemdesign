package main

import (
	"lld-ratelimiter/internal/limiter"
	"lld-ratelimiter/internal/manager"
	"lld-ratelimiter/internal/user"
	"time"
)

func main() {
	cfg := limiter.FixedWindowConfig{
		WindowSize:  10 * time.Second,
		MaxRequests: 10,
	}

	m := manager.New()
	m.AddUsers([]*user.User{
		user.New(1, "Alice", limiter.NewFixedWindowLimiter(cfg)),
		user.New(2, "Bob", limiter.NewFixedWindowLimiter(cfg)),
		user.New(3, "James", limiter.NewFixedWindowLimiter(cfg)),
		user.New(4, "Patricia", limiter.NewFixedWindowLimiter(cfg)),
	})
	m.Simulate(nil)
}
