package main

import (
	"fmt"
	"lld-urlshortner/internal/encoder"
	"lld-urlshortner/internal/generator"
	"lld-urlshortner/internal/observer"
	"lld-urlshortner/internal/stats"
	"lld-urlshortner/internal/url"
	"time"
)

func main() {
	bus := observer.NewEventBus()
	encoder := encoder.NewBase62Encoder()
	generator := generator.NewIncrementalGenerator()
	urlRepository := url.NewInMemoryRepository()
	statsRepository := stats.NewInMemoryRepository()

	us := url.NewService(bus, encoder, generator, urlRepository)

	ss := stats.NewService(statsRepository)
	bus.Subscribe(ss)

	// -----
	// Get the long and short URL
	code1, _ := us.Shorten("https://www.google.com?q=how+to+get+a+job", nil)
	code2, _ := us.Shorten("https://www.youtube.com", nil)
	code3, _ := us.Shorten("https://www.facebook.com", &url.URLConfig{Alias: "custom2", ExpiresAt: time.Now().Add(time.Hour)})

	og1, _ := us.Resolve(code1)
	og2, _ := us.Resolve(code2)
	og3, _ := us.Resolve(code3)
	us.Resolve(code3)

	fmt.Printf("Short: %s, Long: %s\n", code1, og1)
	fmt.Printf("Short: %s, Long: %s\n", code2, og2)
	fmt.Printf("Short: %s, Long: %s\n", code3, og3)

	// -----
	// Now Get Stats
	stat1, _ := ss.GetStats(code1)
	stat2, _ := ss.GetStats(code2)
	stat3, _ := ss.GetStats(code3)

	fmt.Println()
	fmt.Println(stat1)
	fmt.Println(stat2)
	fmt.Println(stat3)

}
