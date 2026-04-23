package manager

import (
	"fmt"
	"lld-ratelimiter/internal/user"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	once     sync.Once
	instance *RateLimiterManager
)

type RateLimiterManager struct {
	users map[int]*user.User
}

// Returns a singleton instance of a rate limiter manager
func New() *RateLimiterManager {
	once.Do(func() {
		instance = &RateLimiterManager{
			users: make(map[int]*user.User),
		}
	})

	return instance
}

func (m *RateLimiterManager) AddUsers(users []*user.User) {
	for _, user := range users {
		m.users[user.ID] = user
	}
}

func (m *RateLimiterManager) Simulate(cfg *SimulateConfig) error {
	if cfg == nil {
		cfg = &SimulateConfig{
			RPS:      10,
			Duration: 10 * time.Second,
		}
	}

	interval := time.Second / time.Duration(cfg.RPS)
	totalRequests := int(cfg.Duration * time.Duration(cfg.RPS))

	g := new(errgroup.Group)
	for _, user := range m.users {
		g.Go(func() error {
			for i := 0; i < totalRequests; i++ {
				if user.Limiter.Allow() {
					fmt.Printf("[%s] Request %d by %s (%d): ALLOWED\n", time.Now().Format("15:04:05"), i+1, user.Name, user.ID)
				} else {
					fmt.Printf("[%s] Request %d by %s (%d): REJECTED\n", time.Now().Format("15:04:05"), i+1, user.Name, user.ID)
				}

				time.Sleep(interval)
			}
			return nil
		})
	}

	return g.Wait()
}
