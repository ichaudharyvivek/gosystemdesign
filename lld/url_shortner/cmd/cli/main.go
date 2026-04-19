package main

import (
	"fmt"
	"lld-urlshortner/internal/encoder"
	"lld-urlshortner/internal/url"
	"time"
)

func main() {
	be := encoder.NewBasicEncoder()
	repo := url.NewInMemoryURLRepository()
	us := url.NewService(be, repo)

	// Shorten the long URL
	code1, _ := us.Shorten("https://www.example.com/1/long/url", nil)
	fmt.Println("Shortened URL code:", code1)

	// Shorten the URL with expiration
	expire := time.Now()
	code2, _ := us.Shorten("https://www.example.com/2/long/url", &url.URLConfig{CustomCode: "astronaut", ExpiresAt: &expire})
	fmt.Println("Shortened URL code:", code2)

	// Resolve the short codes to original URL
	original1, _ := us.Resolve(code1)
	fmt.Println("Original URL for code", code1, "is", original1)

	original2, err := us.Resolve(code2)
	if err != nil {
		fmt.Println("Error resolving code", code2, ":", err)
	} else {
		fmt.Println("Original URL for code", code2, "is", original2)
	}
}
