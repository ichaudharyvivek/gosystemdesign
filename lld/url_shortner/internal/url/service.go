package url

import (
	"fmt"
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

type serviceImpl struct {
	encoder Encoder
	repo    Repository
}

func NewService(encoder Encoder, repo Repository) *serviceImpl {
	return &serviceImpl{
		repo:    repo,
		encoder: encoder,
	}
}

// Todo: add a retry logic here
func (s *serviceImpl) Shorten(original string, config *URLConfig) (string, error) {
	code, err := s.encoder.Generate()
	if err != nil {
		return "", err
	}

	if config != nil && config.CustomCode != "" {
		code = config.CustomCode
	}

	var expiresAt *time.Time
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

	return code, nil
}

func (s *serviceImpl) Resolve(code string) (string, error) {
	url, err := s.repo.FindByCode(code)
	if err != nil {
		return "", fmt.Errorf("cannot finding the url: %w", err)
	}

	// Check if the code has been expired
	if url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt) {
		return "", fmt.Errorf("the short code %s has expired", code)
	}

	return url.Original, nil
}
