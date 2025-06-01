package repository

import (
	"errors"
	m "lld/stackoverflow/models"
	"sync"
)

type QuestionRepo interface {
	Save(question *m.Question) error
	GetByID(id int) (*m.Question, error)
}

type InMemoryQuestionRepo struct {
	questions map[int]*m.Question
	mu        sync.Mutex
}

func NewInMemoryQuestionRepo() QuestionRepo {
	return &InMemoryQuestionRepo{questions: make(map[int]*m.Question)}
}

func (r *InMemoryQuestionRepo) Save(question *m.Question) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.questions[question.ID]
	if exists {
		return errors.New("question already exists")
	}

	r.questions[question.ID] = question
	return nil
}

func (r *InMemoryQuestionRepo) GetByID(id int) (*m.Question, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	question, found := r.questions[id]
	if !found {
		return nil, errors.New("question not found")
	}

	return question, nil
}
