package stats

import (
	"fmt"
	"sync"
	"time"
)

type Repository interface {
	Save(stats *URLStats) error
	FindByCode(code string) (*URLStats, error)
	RecordClick(code string) error
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

func (r *InMemoryRepository) Save(stats *URLStats) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if stats == nil || stats.Code == "" {
		return fmt.Errorf("invalid stats: %w", ErrInvalidInput)
	}

	r.store[stats.Code] = stats
	return nil
}

func (r *InMemoryRepository) FindByCode(code string) (*URLStats, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	stats, exists := r.store[code]
	if !exists {
		return nil, fmt.Errorf("stats for code %s: %w", code, ErrNotFound)
	}

	return stats, nil
}

func (r *InMemoryRepository) RecordClick(code string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if code == "" {
		return fmt.Errorf("invalid code: %w", ErrInvalidInput)
	}

	s, exists := r.store[code]
	if !exists {
		s = &URLStats{Code: code}
		r.store[code] = s
	}

	s.Count++
	s.Logs = append(s.Logs, time.Now())
	return nil
}
