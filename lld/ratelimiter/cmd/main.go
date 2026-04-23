package main

import (
	"lld-ratelimiter/internal/limiter"
	"lld-ratelimiter/internal/manager"
	"lld-ratelimiter/internal/user"
	"time"
)

func main() {
	fixedWindowCfg := limiter.FixedWindowConfig{
		MaxRequests: 5,
		WindowSize:  10 * time.Second,
	}

	m := manager.New()
	m.AddUsers([]*user.User{
		user.New(1, "Alice", limiter.NewFixedWindowLimiter(fixedWindowCfg)),
		user.New(2, "Bob", limiter.NewFixedWindowLimiter(fixedWindowCfg)),
		user.New(3, "James", limiter.NewFixedWindowLimiter(fixedWindowCfg)),
		user.New(4, "Patricia", limiter.NewFixedWindowLimiter(fixedWindowCfg)),
	})
	m.Simulate(nil)
}
