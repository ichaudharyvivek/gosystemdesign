package main

import (
	"fmt"
	"lld-urlshortner/internal/encoder"
	"lld-urlshortner/internal/observer"
	"lld-urlshortner/internal/stats"
	"lld-urlshortner/internal/url"
	"time"
)

func main() {
	bus := observer.NewEventBus()

	statsRepo := stats.NewInMemoryRepository()
	statsService := stats.NewService(statsRepo)
	bus.Subscribe(statsService)

	urlRepo := url.NewInMemoryRepository()
	basicEncoder := encoder.NewBasicEncoder()
	urlService := url.NewService(basicEncoder, urlRepo, bus)

	// Shorten the long URL
	code1, _ := urlService.Shorten("https://www.example.com/1/long/url", nil)
	fmt.Println("Shortened URL code:", code1)

	// Shorten the URL with expiration
	expire := time.Now().Add(1 * time.Minute)
	code2, _ := urlService.Shorten("https://www.example.com/2/long/url", &url.URLConfig{CustomCode: "astronaut", ExpiresAt: &expire})
	fmt.Println("Shortened URL code:", code2)

	// Resolve the short codes to original URL
	original1, _ := urlService.Resolve(code1)
	fmt.Println("Original URL for code", code1, "is", original1)

	original2, err := urlService.Resolve(code2)
	if err != nil {
		fmt.Println("Error resolving code", code2, ":", err)
	} else {
		fmt.Println("Original URL for code", code2, "is", original2)
	}

	// Gather stats
	stats1, _ := statsService.GetStats(code1)
	fmt.Println("Stats for code", code1, ":", stats1)

	stats2, _ := statsService.GetStats(code2)
	fmt.Println("Stats for code", code2, ":", stats2)
}
