package limiter

import (
	"fmt"
	"sync"
	"time"
)

type FixedWindowConfig struct {
	MaxRequests int
	WindowSize  time.Duration
}

type FixedWindowLimiter struct {
	mu                sync.Mutex
	maxRequests       int
	remainingRequests int
	windowSize        time.Duration
	windowStart       time.Time
}

func NewFixedWindowLimiter(cfg FixedWindowConfig) *FixedWindowLimiter {
	now := time.Now()
	return &FixedWindowLimiter{
		maxRequests:       cfg.MaxRequests,
		remainingRequests: cfg.MaxRequests,
		windowSize:        cfg.WindowSize,
		windowStart:       now,
	}
}

func (l *FixedWindowLimiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(l.windowStart)
	if elapsed >= l.windowSize {
		// Reset the max limit
		l.remainingRequests = l.maxRequests

		// Reset the window start time for the current window
		// This will keep the window truly fixed: [0-10] [11-20] [21-30] etc.
		windowsPassed := int(elapsed / l.windowSize)
		l.windowStart = l.windowStart.Add(l.windowSize * time.Duration(windowsPassed))

	}

	fmt.Println("Remaining requests:", l.remainingRequests)

	if l.remainingRequests <= 0 {
		return false
	}

	l.remainingRequests--
	return true
}
