package perplexity

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/perplexityai/perplexity-go/perplexity/chat"
	internalhttp "github.com/perplexityai/perplexity-go/perplexity/internal/http"
)

// Client is the main Perplexity API client.
type Client struct {
	apiKey         string
	baseURL        string
	httpClient     *http.Client
	maxRetries     int
	defaultHeaders map[string]string
	userAgent      string

	// Services
	Chat *chat.Service
	// Search *search.Service  // Will be added in Phase 5
}

// NewClient creates a new Perplexity API client.
// The API key can be provided directly or via the PERPLEXITY_API_KEY environment variable.
func NewClient(apiKey string, opts ...ClientOption) (*Client, error) {
	// If no API key provided, try to get from environment
	if apiKey == "" {
		apiKey = os.Getenv("PERPLEXITY_API_KEY")
	}

	if apiKey == "" {
		return nil, errors.New("perplexity: API key is required (provide via parameter or PERPLEXITY_API_KEY environment variable)")
	}

	client := &Client{
		apiKey:         apiKey,
		baseURL:        DefaultBaseURL,
		maxRetries:     DefaultMaxRetries,
		userAgent:      DefaultUserAgent,
		defaultHeaders: make(map[string]string),
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, fmt.Errorf("perplexity: failed to apply option: %w", err)
		}
	}

	// Initialize internal HTTP client
	httpClientWrapper := internalhttp.NewClient(
		client.httpClient,
		client.baseURL,
		client.apiKey,
		client.maxRetries,
		client.defaultHeaders,
		client.userAgent,
	)

	// Initialize services
	client.Chat = chat.NewService(httpClientWrapper)
	// client.Search = search.NewService(httpClientWrapper)  // Phase 5

	return client, nil
}

// APIKey returns the API key being used by the client.
func (c *Client) APIKey() string {
	return c.apiKey
}

// BaseURL returns the base URL being used by the client.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// Version returns the SDK version.
func (c *Client) Version() string {
	return Version
}
