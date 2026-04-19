package url

import (
	"errors"
	"fmt"
	"lld-urlshortner/internal/encoder"
	"lld-urlshortner/internal/observer"
	"time"
)

const MAX_RETRIES = 3

type URLConfig struct {
	CustomCode string
	ExpiresAt  *time.Time
}

type Service interface {
	Shorten(original string, config *URLConfig) (string, error)
	Resolve(code string) (string, error)
}

type service struct {
	repo    Repository
	encoder encoder.Encoder
	bus     *observer.EventBus
}

func NewService(encoder encoder.Encoder, repo Repository, bus *observer.EventBus) *service {
	return &service{
		bus:     bus,
		repo:    repo,
		encoder: encoder,
	}
}

func (s *service) Shorten(original string, config *URLConfig) (string, error) {
	var (
		err       error
		code      string
		expiresAt *time.Time
	)

	now := time.Now()
	if config != nil && config.CustomCode != "" {
		code = config.CustomCode

		_, err = s.repo.FindByCode(code)
		if err == nil {
			return "", fmt.Errorf("custom code already exists")
		}
		if !errors.Is(err, ErrNotFound) {
			return "", fmt.Errorf("repo error: %w", err)
		}
	} else {
		found := false
		for i := 0; i < MAX_RETRIES; i++ {
			code, err = s.encoder.Generate()
			if err != nil {
				return "", fmt.Errorf("generate failed: %w", err)
			}

			_, err = s.repo.FindByCode(code)
			if errors.Is(err, ErrNotFound) {
				found = true
				break
			}
			if err != nil {
				return "", fmt.Errorf("error in generating code: %w", err)
			}
		}

		if !found {
			return "", fmt.Errorf("failed to generate unique code after %d retries", MAX_RETRIES)
		}
	}

	if config != nil && config.ExpiresAt != nil {
		expiresAt = config.ExpiresAt
	}

	url := &URL{
		Code:      code,
		Original:  original,
		ExpiresAt: expiresAt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.repo.Save(url); err != nil {
		return "", fmt.Errorf("save failed: %w", err)
	}

	s.bus.Notify("url.created", url)
	return code, nil
}

func (s *service) Resolve(code string) (string, error) {
	url, err := s.repo.FindByCode(code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return "", err
		}
		return "", fmt.Errorf("repo error: %w", err)
	}

	if url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt) {
		return "", fmt.Errorf("short code %s has expired", code)
	}

	s.bus.Notify("url.accessed", url)
	return url.Original, nil
}
