package http

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	// InitialRetryDelay is the initial delay for exponential backoff.
	InitialRetryDelay = 500 * time.Millisecond

	// MaxRetryDelay is the maximum delay between retries.
	MaxRetryDelay = 8 * time.Second
)

type ErrorKind int

const (
	ErrorKindStatus ErrorKind = iota
	ErrorKindConnection
	ErrorKindTimeout
)

type ErrorFactory func(kind ErrorKind, statusCode int, message string, body []byte, requestID string, cause error) error

// Client wraps an HTTP client with retry logic and error handling.
type Client struct {
	httpClient     *http.Client
	baseURL        string
	apiKey         string
	maxRetries     int
	defaultHeaders map[string]string
	userAgent      string
	errorFactory   ErrorFactory
}

// NewClient creates a new HTTP client wrapper.
func NewClient(httpClient *http.Client, baseURL, apiKey string, maxRetries int, defaultHeaders map[string]string, userAgent string, errorFactory ErrorFactory) *Client {
	return &Client{
		httpClient:     httpClient,
		baseURL:        baseURL,
		apiKey:         apiKey,
		maxRetries:     maxRetries,
		defaultHeaders: defaultHeaders,
		userAgent:      userAgent,
		errorFactory:   errorFactory,
	}
}

// Request represents an HTTP request.
type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    interface{}
}

// Response represents an HTTP response.
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
	RequestID  string
}

// StreamResponse represents a streaming HTTP response.
type StreamResponse struct {
	StatusCode int
	Headers    http.Header
	Response   *http.Response
	RequestID  string
}

// Do executes an HTTP request with retry logic.
func (c *Client) Do(ctx context.Context, req *Request) (*Response, error) {
	var lastErr error
	var retryDelay time.Duration

	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		if attempt > 0 {
			select {
			case <-time.After(retryDelay):
				// Continue with retry
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}

		resp, err := c.doRequest(ctx, req)
		if err != nil {
			lastErr = c.wrapTransportError(err)
			if !c.shouldRetryError(err) {
				return nil, lastErr
			}
			retryDelay = c.calculateBackoff(attempt+1, nil)
			continue
		}

		if c.shouldRetryResponse(resp) {
			lastErr = c.errorFromResponse(resp)
			retryDelay = c.calculateBackoff(attempt+1, resp.Headers)
			continue
		}

		if resp.StatusCode >= 400 {
			return nil, c.errorFromResponse(resp)
		}

		return resp, nil
	}

	// Max retries exceeded
	if lastErr != nil {
		return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
	}
	return nil, fmt.Errorf("max retries exceeded")
}

// doRequest performs a single HTTP request.
func (c *Client) doRequest(ctx context.Context, req *Request) (*Response, error) {
	// Build URL
	url := c.baseURL + req.Path

	// Marshal body if present
	var bodyReader io.Reader
	if req.Body != nil {
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	httpReq.Header.Set("Accept", "application/json")
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", c.userAgent)
	httpReq.Header.Set("X-Stainless-Async", "false")

	// Add default headers
	for key, value := range c.defaultHeaders {
		httpReq.Header.Set(key, value)
	}

	// Add request-specific headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Execute request
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer httpResp.Body.Close()

	// Read response body
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Extract request ID from headers
	requestID := httpResp.Header.Get("X-Request-Id")
	if requestID == "" {
		requestID = httpResp.Header.Get("X-Request-ID")
	}

	return &Response{
		StatusCode: httpResp.StatusCode,
		Headers:    httpResp.Header,
		Body:       body,
		RequestID:  requestID,
	}, nil
}

// shouldRetryError determines if an error should trigger a retry.
func (c *Client) shouldRetryError(err error) bool {
	if err == nil {
		return false
	}
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	var netErr net.Error
	return errors.As(err, &netErr)
}

// shouldRetryStatus determines if a status code should trigger a retry.

func (c *Client) shouldRetryResponse(resp *Response) bool {
	if resp == nil {
		return false
	}
	switch resp.Headers.Get("x-should-retry") {
	case "true":
		return true
	case "false":
		return false
	}
	return c.shouldRetryStatus(resp.StatusCode)
}

func (c *Client) shouldRetryStatus(statusCode int) bool {
	switch statusCode {
	case http.StatusRequestTimeout, // 408
		http.StatusConflict,        // 409
		http.StatusTooManyRequests: // 429
		return true
	default:
		// Retry on 5xx errors
		return statusCode >= 500
	}
}

// calculateBackoff calculates the backoff delay for a given attempt.

func (c *Client) calculateBackoff(attempt int, headers http.Header) time.Duration {
	if retryAfter, ok := parseRetryAfter(headers); ok && retryAfter > 0 && retryAfter <= 60*time.Second {
		return retryAfter
	}

	delay := InitialRetryDelay * time.Duration(math.Pow(2, float64(attempt-1)))
	if delay > MaxRetryDelay {
		delay = MaxRetryDelay
	}
	var randomBytes [8]byte
	if _, err := rand.Read(randomBytes[:]); err == nil {
		randomValue := binary.BigEndian.Uint64(randomBytes[:])
		randomFloat := float64(randomValue) / float64(^uint64(0))
		jitter := time.Duration(randomFloat * 0.25 * float64(delay))
		if randomBytes[0]&1 == 0 {
			delay += jitter
		} else {
			delay -= jitter
		}
	}

	return delay
}

func parseRetryAfter(headers http.Header) (time.Duration, bool) {
	if headers == nil {
		return 0, false
	}
	if retryAfterMS := headers.Get("retry-after-ms"); retryAfterMS != "" {
		if value, err := strconv.ParseFloat(retryAfterMS, 64); err == nil {
			return time.Duration(value * float64(time.Millisecond)), true
		}
	}
	if retryAfter := headers.Get("retry-after"); retryAfter != "" {
		if value, err := strconv.ParseFloat(retryAfter, 64); err == nil {
			return time.Duration(value * float64(time.Second)), true
		}
		if when, err := http.ParseTime(retryAfter); err == nil {
			return time.Until(when), true
		}
	}
	return 0, false
}

// errorFromResponse creates an error from an HTTP response.
// This will be replaced with proper error handling using the perplexity package errors.
func (c *Client) errorFromResponse(resp *Response) error {
	// Try to parse error message from response body
	var errorBody struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}

	message := fmt.Sprintf("HTTP %d", resp.StatusCode)
	if err := json.Unmarshal(resp.Body, &errorBody); err == nil {
		if errorBody.Message != "" {
			message = errorBody.Message
		} else if errorBody.Error != "" {
			message = errorBody.Error
		}
	}

	if c.errorFactory != nil {
		if err := c.errorFactory(ErrorKindStatus, resp.StatusCode, message, resp.Body, resp.RequestID, nil); err != nil {
			return err
		}
	}
	return fmt.Errorf("%s (status: %d, request_id: %s)", message, resp.StatusCode, resp.RequestID)
}

func (c *Client) wrapTransportError(err error) error {
	if err == nil {
		return nil
	}
	if c.errorFactory == nil {
		return fmt.Errorf("request failed: %w", err)
	}
	kind := ErrorKindConnection
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		kind = ErrorKindTimeout
	}
	if wrapped := c.errorFactory(kind, 0, err.Error(), nil, "", err); wrapped != nil {
		return wrapped
	}
	return fmt.Errorf("request failed: %w", err)
}

// DoStream executes a streaming HTTP request.
// The caller is responsible for closing the response body.
func (c *Client) DoStream(ctx context.Context, req *Request) (*StreamResponse, error) {
	// Build URL
	url := c.baseURL + req.Path

	// Marshal body if present
	var bodyReader io.Reader
	if req.Body != nil {
		bodyBytes, err := json.Marshal(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// Create HTTP request
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	httpReq.Header.Set("Authorization", "Bearer "+c.apiKey)
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("User-Agent", c.userAgent)
	httpReq.Header.Set("Accept", "text/event-stream")
	httpReq.Header.Set("Cache-Control", "no-cache")
	httpReq.Header.Set("Connection", "keep-alive")
	httpReq.Header.Set("X-Stainless-Async", "false")

	// Add default headers
	for key, value := range c.defaultHeaders {
		httpReq.Header.Set(key, value)
	}

	// Add request-specific headers
	for key, value := range req.Headers {
		httpReq.Header.Set(key, value)
	}

	// Execute request
	httpResp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, c.wrapTransportError(err)
	}

	// Check for error status codes
	if httpResp.StatusCode >= 400 {
		// Read error body
		body, _ := io.ReadAll(httpResp.Body)
		_ = httpResp.Body.Close() // Explicitly ignore close error for error response

		var errorBody struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}

		message := fmt.Sprintf("HTTP %d", httpResp.StatusCode)
		if err := json.Unmarshal(body, &errorBody); err == nil {
			if errorBody.Message != "" {
				message = errorBody.Message
			} else if errorBody.Error != "" {
				message = errorBody.Error
			}
		}

		requestID := httpResp.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = httpResp.Header.Get("X-Request-ID")
		}

		if c.errorFactory != nil {
			if err := c.errorFactory(ErrorKindStatus, httpResp.StatusCode, message, body, requestID, nil); err != nil {
				return nil, err
			}
		}
		return nil, fmt.Errorf("%s (status: %d, request_id: %s)", message, httpResp.StatusCode, requestID)
	}

	// Extract request ID from headers
	requestID := httpResp.Header.Get("X-Request-Id")
	if requestID == "" {
		requestID = httpResp.Header.Get("X-Request-ID")
	}

	return &StreamResponse{
		StatusCode: httpResp.StatusCode,
		Headers:    httpResp.Header,
		Response:   httpResp,
		RequestID:  requestID,
	}, nil
}
