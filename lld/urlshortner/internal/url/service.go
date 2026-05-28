package url

import (
	"errors"
	"fmt"
	"lld-urlshortner/internal/encoder"
	"lld-urlshortner/internal/generator"
	"lld-urlshortner/internal/observer"
	"time"
)

type Service interface {
	Shorten(url string, config *URLConfig) (string, error)
	Resolve(code string) (string, error)
}

type URLConfig struct {
	Alias     string
	ExpiresAt time.Time
}

const maxRetries = 3

type service struct {
	bus       *observer.EventBus
	encoder   encoder.Encoder
	generator generator.Generator
	repo      Repository
}

func NewService(bus *observer.EventBus, encoder encoder.Encoder, generator generator.Generator, repo Repository) *service {
	return &service{
		bus:       bus,
		repo:      repo,
		encoder:   encoder,
		generator: generator,
	}
}

func (s *service) Shorten(url string, config *URLConfig) (string, error) {
	now := time.Now()
	expiresAt := time.Time{}

	if config != nil && config.Alias != "" {
		expiresAt = config.ExpiresAt
		u := &URL{
			Code:      config.Alias,
			Original:  url,
			ExpiresAt: expiresAt,
			CreatedAt: now,
			UpdatedAt: now,
		}

		if err := s.repo.Save(u); err != nil {
			if errors.Is(err, ErrAlreadyExists) {
				return "", fmt.Errorf("alias %s is already taken: %w", config.Alias, err)
			}
			return "", fmt.Errorf("unable to store the short url: %w", err)
		}

		s.bus.Notify("url.created", config.Alias)
		return config.Alias, nil
	}

	if config != nil {
		expiresAt = config.ExpiresAt
	}

	var lastErr error
	for range maxRetries {
		id := s.generator.Generate()
		code := s.encoder.Encode(id)

		u := &URL{
			Code:      code,
			Original:  url,
			ExpiresAt: expiresAt,
			CreatedAt: now,
			UpdatedAt: now,
		}

		lastErr = s.repo.Save(u)
		if lastErr == nil {
			s.bus.Notify("url.created", code)
			return code, nil
		}

		if !errors.Is(lastErr, ErrAlreadyExists) {
			return "", fmt.Errorf("unable to store the short url with code %s: %w", code, lastErr)
		}
	}

	return "", fmt.Errorf("failed to generate unique code after %d attempts: %w", maxRetries, lastErr)
}

func (s *service) Resolve(code string) (string, error) {
	url, err := s.repo.FindByCode(code)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return "", fmt.Errorf("no URL found for code %s: %w", code, err)
		}
		return "", fmt.Errorf("error resolving code %s: %w", code, err)
	}

	s.bus.Notify("url.accessed", code)
	return url.Original, nil
}
