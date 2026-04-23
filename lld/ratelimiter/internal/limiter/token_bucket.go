package limiter

import (
	"sync"
	"time"
)

type TokenBucketConfig struct {
	Capacity   int
	RefillRate float64
}

type TokenBucketLimiter struct {
	mu             sync.Mutex
	tokens         float64
	capacity       float64
	refilRate      float64
	lastRefillTime time.Time
}

func NewTokenBucketLimiter(cfg TokenBucketConfig) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		tokens:         float64(cfg.Capacity),
		capacity:       float64(cfg.Capacity),
		refilRate:      cfg.RefillRate,
		lastRefillTime: time.Now(),
	}
}

func (l *TokenBucketLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(l.lastRefillTime).Seconds()
	if elapsed > 0 {
		// The tokens are filled in a continuous manner to prevent edge spikes
		l.tokens += elapsed * l.refilRate
		l.tokens = min(l.capacity, l.tokens)
		l.lastRefillTime = now
	}

	if l.tokens <= 0 {
		return false
	}

	l.tokens--
	return true
}
