package url

import (
	"fmt"
	"lld-urlshortner/internal/encoder"
	"time"
)

const (
	MAX_RETRIES = 3
)

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
}

func NewService(encoder encoder.Encoder, repo Repository) *service {
	return &service{
		repo:    repo,
		encoder: encoder,
	}
}

// Todo: add a retry logic here
func (s *service) Shorten(original string, config *URLConfig) (string, error) {
	var err error
	var code string
	var expiresAt *time.Time

	if config != nil {
		if config.CustomCode != "" {
			code = config.CustomCode

			// Check if the custom code already exists
			_, ok := s.repo.FindByCode(code)
			if ok != nil {
				return "", fmt.Errorf("the custom code %s already exists", code)
			}
		}

		if config.ExpiresAt != nil {
			expiresAt = config.ExpiresAt
		}
	}

	code, err := s.encoder.Generate()
	if err != nil {
		return "", err
	}

	if config != nil && config.CustomCode != "" {
		code = config.CustomCode
	}

	if config != nil && config.ExpiresAt != nil {
		expiresAt = config.ExpiresAt
	}

	url := &URL{
		Code:      code,
		Original:  original,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := s.repo.Save(url); err != nil {
		return "", fmt.Errorf("something went wrong while saving the shortened url: %w", err)
	}

	return "s", nil
}

func (s *service) Resolve(code string) (string, error) {
	url, err := s.repo.FindByCode(code)
	if err != nil {
		return "", fmt.Errorf("error in fetching the url: %w", err)
	}

	// Check if the code has been expired
	if url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt) {
		return "", fmt.Errorf("the short code %s has expired", code)
	}

	return url.Original, nil
}
