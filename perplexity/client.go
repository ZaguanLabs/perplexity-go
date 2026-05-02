package perplexity

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/ZaguanLabs/perplexity-go/perplexity/asyncchat"
	"github.com/ZaguanLabs/perplexity-go/perplexity/browser"
	"github.com/ZaguanLabs/perplexity-go/perplexity/chat"
	"github.com/ZaguanLabs/perplexity-go/perplexity/contextualizedembeddings"
	"github.com/ZaguanLabs/perplexity-go/perplexity/embeddings"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
	"github.com/ZaguanLabs/perplexity-go/perplexity/responses"
	"github.com/ZaguanLabs/perplexity-go/perplexity/search"
)

func clientErrorFactory(kind internalhttp.ErrorKind, statusCode int, message string, body []byte, requestID string, cause error) error {
	switch kind {
	case internalhttp.ErrorKindStatus:
		return newError(statusCode, message, body, requestID)
	case internalhttp.ErrorKindTimeout:
		msg := message
		if cause != nil {
			msg = fmt.Sprintf("%s: %v", message, cause)
		}
		return &TimeoutError{Err: &Error{Message: msg, StatusCode: statusCode, Body: body, RequestID: requestID}}
	case internalhttp.ErrorKindConnection:
		msg := message
		if cause != nil {
			msg = fmt.Sprintf("%s: %v", message, cause)
		}
		return &ConnectionError{Err: &Error{Message: msg, StatusCode: statusCode, Body: body, RequestID: requestID}}
	default:
		if cause != nil {
			return fmt.Errorf("perplexity: %s: %w", message, cause)
		}
		return fmt.Errorf("perplexity: %s", message)
	}
}

// Client is the main Perplexity API client.
type Client struct {
	apiKey         string
	baseURL        string
	httpClient     *http.Client
	maxRetries     int
	defaultHeaders map[string]string
	defaultQuery   map[string]any
	userAgent      string

	// Services
	Chat                     *chat.Service
	Search                   *search.Service
	AsyncChat                *asyncchat.Service
	Responses                *responses.Service
	Embeddings               *embeddings.Service
	ContextualizedEmbeddings *contextualizedembeddings.Service
	Browser                  *browser.Service
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
		defaultHeaders: defaultPlatformHeaders(),
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	if baseURL := os.Getenv("PERPLEXITY_BASE_URL"); baseURL != "" {
		client.baseURL = baseURL
	}

	// Apply options
	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, fmt.Errorf("perplexity: failed to apply option: %w", err)
		}
	}

	client.initServices()

	return client, nil
}

func (c *Client) initServices() {
	httpClientWrapper := internalhttp.NewClient(
		c.httpClient,
		c.baseURL,
		c.apiKey,
		c.maxRetries,
		c.defaultHeaders,
		c.userAgent,
		clientErrorFactory,
	)
	httpClientWrapper.SetDefaultQuery(c.defaultQuery)

	c.Chat = chat.NewService(httpClientWrapper)
	c.Search = search.NewService(httpClientWrapper)
	c.AsyncChat = asyncchat.NewService(httpClientWrapper)
	c.Responses = responses.NewService(httpClientWrapper)
	c.Embeddings = embeddings.NewService(httpClientWrapper)
	c.ContextualizedEmbeddings = contextualizedembeddings.NewService(httpClientWrapper)
	c.Browser = browser.NewService(httpClientWrapper)
}

// APIKey returns the API key being used by the client.
func (c *Client) APIKey() string {
	return c.apiKey
}

func (c *Client) Copy(opts ...ClientOption) (*Client, error) {
	if c == nil {
		return nil, errors.New("perplexity: cannot copy nil client")
	}
	copyClient := &Client{
		apiKey:         c.apiKey,
		baseURL:        c.baseURL,
		httpClient:     c.httpClient,
		maxRetries:     c.maxRetries,
		defaultHeaders: cloneStringMap(c.defaultHeaders),
		defaultQuery:   cloneAnyMap(c.defaultQuery),
		userAgent:      c.userAgent,
	}
	for _, opt := range opts {
		if err := opt(copyClient); err != nil {
			return nil, fmt.Errorf("perplexity: failed to apply option: %w", err)
		}
	}
	copyClient.initServices()
	return copyClient, nil
}

func (c *Client) WithOptions(opts ...ClientOption) (*Client, error) {
	return c.Copy(opts...)
}

func cloneStringMap(input map[string]string) map[string]string {
	if input == nil {
		return nil
	}
	output := make(map[string]string, len(input))
	for key, value := range input {
		output[key] = value
	}
	return output
}

func cloneAnyMap(input map[string]any) map[string]any {
	if input == nil {
		return nil
	}
	output := make(map[string]any, len(input))
	for key, value := range input {
		output[key] = value
	}
	return output
}

// BaseURL returns the base URL being used by the client.
func (c *Client) BaseURL() string {
	return c.baseURL
}

// Version returns the SDK version.
func (c *Client) Version() string {
	return Version
}
