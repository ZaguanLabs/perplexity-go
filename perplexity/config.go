package perplexity

import (
	"net/http"
	"time"
)

const (
	// DefaultBaseURL is the default base URL for the Perplexity API.
	DefaultBaseURL = "https://api.perplexity.ai"

	// DefaultTimeout is the default timeout for HTTP requests (15 minutes).
	DefaultTimeout = 15 * time.Minute

	// DefaultMaxRetries is the default maximum number of retries for failed requests.
	DefaultMaxRetries = 2

	// DefaultUserAgent is the default User-Agent header.
	DefaultUserAgent = "perplexity-go/0.1.0"
)

// ClientOption is a functional option for configuring the Client.
type ClientOption func(*Client) error

// WithBaseURL sets a custom base URL for the API.
func WithBaseURL(url string) ClientOption {
	return func(c *Client) error {
		c.baseURL = url
		return nil
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) ClientOption {
	return func(c *Client) error {
		c.httpClient = client
		return nil
	}
}

// WithMaxRetries sets the maximum number of retries for failed requests.
func WithMaxRetries(retries int) ClientOption {
	return func(c *Client) error {
		c.maxRetries = retries
		return nil
	}
}

// WithTimeout sets the request timeout.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) error {
		if c.httpClient == nil {
			c.httpClient = &http.Client{
				Timeout: timeout,
			}
		} else {
			c.httpClient.Timeout = timeout
		}
		return nil
	}
}

// WithDefaultHeader adds a default header to all requests.
func WithDefaultHeader(key, value string) ClientOption {
	return func(c *Client) error {
		if c.defaultHeaders == nil {
			c.defaultHeaders = make(map[string]string)
		}
		c.defaultHeaders[key] = value
		return nil
	}
}
