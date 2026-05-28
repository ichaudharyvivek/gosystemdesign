package main

import (
	"fmt"
	"lld-urlshortner/internal/encoder"
	"lld-urlshortner/internal/generator"
	"lld-urlshortner/internal/observer"
	"lld-urlshortner/internal/stats"
	"lld-urlshortner/internal/url"
	"sync"
)

func main() {
	bus := observer.NewEventBus()
	enc := encoder.NewBase62Encoder()
	gen := generator.NewIncrementalGenerator()
	urlRepo := url.NewInMemoryRepository()
	statsRepo := stats.NewInMemoryRepository()

	us := url.NewService(bus, enc, gen, urlRepo)
	ss := stats.NewService(statsRepo)
	bus.Subscribe(ss)

	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.facebook.com",
		"https://www.github.com",
		"https://www.stackoverflow.com",
	}

	var wg sync.WaitGroup
	codes := make([]string, len(urls))

	for i, u := range urls {
		wg.Add(1)
		go func(idx int, rawURL string) {
			defer wg.Done()
			code, err := us.Shorten(rawURL, nil)
			if err != nil {
				fmt.Printf("error shortening %s: %v\n", rawURL, err)
				return
			}
			codes[idx] = code
			fmt.Printf("[goroutine %d] shortened: %s -> %s\n", idx, rawURL, code)
		}(i, u)
	}
	wg.Wait()

	fmt.Println("\n--- Resolving (3 clicks each) ---")
	for i, code := range codes {
		for click := 0; click < 3; click++ {
			wg.Add(1)
			go func(idx, c int, cd string) {
				defer wg.Done()
				original, err := us.Resolve(cd)
				if err != nil {
					fmt.Printf("error resolving %s: %v\n", cd, err)
					return
				}
				fmt.Printf("[goroutine %d, click %d] resolved: %s -> %s\n", idx, c, cd, original)
			}(i, click, code)
		}
	}
	wg.Wait()

	fmt.Println("\n--- Stats ---")
	for _, code := range codes {
		stat, err := ss.GetStats(code)
		if err != nil {
			fmt.Printf("error getting stats for %s: %v\n", code, err)
			continue
		}
		fmt.Printf("code: %s, clicks: %d\n", stat.Code, stat.Count)
	}
}
