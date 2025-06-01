package repository

import (
	"errors"
	m "lld/stackoverflow/models"
	"sync"
)

type AnswerRepo interface {
	Save(answer *m.Answer) error
	GetByID(id int) (*m.Answer, error)
}

type InMemoryAnswerRepo struct {
	answers map[int]*m.Answer
	mu      sync.Mutex
}

func NewInMemoryAnswerRepo() AnswerRepo {
	return &InMemoryAnswerRepo{answers: make(map[int]*m.Answer)}
}

func (r *InMemoryAnswerRepo) Save(answer *m.Answer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.answers[answer.ID]
	if exists {
		return errors.New("answer already exists")
	}

	r.answers[answer.ID] = answer
	return nil
}

func (r *InMemoryAnswerRepo) GetByID(id int) (*m.Answer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	answer, found := r.answers[id]
	if !found {
		return nil, errors.New("answer not found")
	}

	return answer, nil
}
