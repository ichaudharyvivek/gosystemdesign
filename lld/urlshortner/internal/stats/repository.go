package stats

import (
	"fmt"
	"sync"
)

type Repository interface {
	GetStats(code string) (*URLStats, error)
	SetStats(code string, stats *URLStats) error
}

type InMemoryRepository struct {
	mu    sync.RWMutex
	store map[string]*URLStats
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		store: make(map[string]*URLStats),
	}
}

func (r *InMemoryRepository) GetStats(code string) (*URLStats, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	stats, exists := r.store[code]
	if !exists {
		return nil, fmt.Errorf("stats for code %s: %w", code, ErrNotFound)
	}

	return stats, nil
}

func (r *InMemoryRepository) SetStats(code string, stats *URLStats) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if code == "" {
		return fmt.Errorf("invalid code: %w", ErrInvalidInput)
	}

	r.store[code] = stats
	return nil
}
