package stats

import (
	"errors"
	"fmt"
)

type Service interface {
	GetStats(code string) (*URLStats, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetStats(code string) (*URLStats, error) {
	stats, err := s.repo.FindByCode(code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return nil, fmt.Errorf("error in getting stats: %w", err)
		}

		return nil, fmt.Errorf("something went wrong: %w", err)
	}

	return stats, nil
}

func (s *service) Update(event string, data any) {
	code, ok := data.(string)
	if !ok {
		fmt.Printf("invalid data type for event %s\n", event)
		return
	}

	switch event {
	case "url.created":
		stats := &URLStats{Code: code}
		if err := s.repo.Save(stats); err != nil {
			fmt.Printf("failed to save stats for code %s: %v\n", code, err)
		}

	case "url.accessed":
		if err := s.repo.RecordClick(code); err != nil {
			fmt.Printf("failed to record click for code %s: %v\n", code, err)
		}
	}
}
