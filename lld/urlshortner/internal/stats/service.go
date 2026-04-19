package stats

import (
	"errors"
	"fmt"
	"lld-urlshortner/internal/url"
	"time"
)

type Service interface {
	GetStats(code string) (*URLStats, error)
	SetStats(code string, stats *URLStats) error
	RecordClick(code string) error
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
	return s.repo.GetStats(code)
}

func (s *service) SetStats(code string, stats *URLStats) error {
	return s.repo.SetStats(code, stats)
}

func (s *service) RecordClick(code string) error {
	stats, err := s.repo.GetStats(code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return fmt.Errorf("cannot record click: %w", err)
		}
	}

	stats.Clicks++
	stats.AccessLogs = append(stats.AccessLogs, time.Now())
	return s.repo.SetStats(code, stats)
}

func (s *service) Update(event string, data any) {
	url := data.(*url.URL)

	switch event {
	case "url.created":
		fmt.Println("Event:", event)
		urlStats := &URLStats{
			Code:       url.Code,
			Clicks:     1,
			AccessLogs: []time.Time{time.Now()},
		}
		s.SetStats(url.Code, urlStats)

	case "url.accessed":
		fmt.Println("Event:", event)
		s.RecordClick(url.Code)
	}
}
