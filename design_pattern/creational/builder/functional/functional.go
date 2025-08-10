/*
This pattern is called "functional options" pattern.
It is used to create an object just like building but instead of method chaining ( .X().Y().Z() ), it uses variable args

NOTE: This is go specific only.
*/
package functional

import (
	"time"
)

// Option function type
// This function takes the pointer and changes the object in place
type Option func(*HTTPClient)

// Option to set base URL
func WithBaseURL(url string) Option {
	return func(c *HTTPClient) {
		c.BaseURL = url
	}
}

// Option to set request timeout
func WithTimeout(d time.Duration) Option {
	return func(c *HTTPClient) {
		c.Timeout = d
		c.client.Timeout = d
	}
}

// Option to set default headers
func WithHeaders(headers map[string]string) Option {
	return func(c *HTTPClient) {
		for k, v := range headers {
			c.Headers[k] = v
		}
	}
}

// Option to set retry count
func WithRetryCount(count int) Option {
	return func(c *HTTPClient) {
		c.RetryCount = count
	}
}
