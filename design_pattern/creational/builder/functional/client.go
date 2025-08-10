package functional

import (
	"fmt"
	"net/http"
	"time"
)

// Imagine u are building a http client object with variable configurations
type HTTPClient struct {
	BaseURL    string
	Timeout    time.Duration
	Headers    map[string]string
	RetryCount int
	client     *http.Client
}

// Constructor
func NewHTTPClient(opts ...Option) *HTTPClient {
	// Default values
	c := &HTTPClient{
		BaseURL:    "https://api.example.com",
		Timeout:    5 * time.Second,
		Headers:    map[string]string{"Content-Type": "application/json"},
		RetryCount: 3,
		client:     &http.Client{Timeout: 5 * time.Second},
	}

	// Override default values with whatever options provided
	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Sample function
func (c *HTTPClient) Ping() {
	fmt.Println("Making request to:", c.BaseURL)
	fmt.Println("Timeout:", c.Timeout)
	fmt.Println("Headers:", c.Headers)
	fmt.Println("Retries:", c.RetryCount)
	fmt.Println()
}
