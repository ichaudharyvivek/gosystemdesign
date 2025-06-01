package repository

import (
	"errors"
	m "lld/stackoverflow/models"
	"sync"
)

type UserRepo interface {
	Save(user *m.User) error
	GetByID(id int) (*m.User, error)
}

type InMemoryUserRepo struct {
	users map[int]*m.User
	mu    sync.Mutex
}

func NewInMemoryUserRepo() UserRepo {
	return &InMemoryUserRepo{users: make(map[int]*m.User)}
}

func (r *InMemoryUserRepo) Save(user *m.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.users[user.ID]
	r.users[user.ID] = user
	if exists {
		return errors.New("user already exists")
	}

	return nil
}

func (r *InMemoryUserRepo) GetByID(id int) (*m.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user, found := r.users[id]
	if !found {
		return nil, errors.New("user not found")
	}

	return user, nil
}
