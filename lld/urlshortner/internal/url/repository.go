package url

import (
	"fmt"
	"sync"
)

type Repository interface {
	Save(url *URL) error
	FindByCode(code string) (*URL, error)
}

type InMemoryRepository struct {
	mu    sync.RWMutex
	store map[string]*URL
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		store: make(map[string]*URL),
	}
}

func (r *InMemoryRepository) Save(url *URL) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if url == nil {
		return ErrInvalidInput
	}

	if _, exists := r.store[url.Code]; exists {
		return fmt.Errorf("code %s: %w", url.Code, ErrAlreadyExists)
	}

	r.store[url.Code] = url
	return nil
}

func (r *InMemoryRepository) FindByCode(code string) (*URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, exists := r.store[code]
	if !exists {
		return nil, fmt.Errorf("code %s: %w", code, ErrNotFound)
	}

	return url, nil
}
