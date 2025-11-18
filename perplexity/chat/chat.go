package chat

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

// Service provides access to chat completion APIs.
type Service struct {
	client *http.Client
}

// NewService creates a new chat service.
func NewService(httpClient *http.Client) *Service {
	return &Service{
		client: httpClient,
	}
}

// Create generates a chat completion for the given parameters.
// This method does not support streaming. Use CreateStream for streaming responses.
func (s *Service) Create(ctx context.Context, params *CompletionParams) (*types.StreamChunk, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	// Validate required fields
	if len(params.Messages) == 0 {
		return nil, fmt.Errorf("messages are required")
	}
	if params.Model == "" {
		return nil, fmt.Errorf("model is required")
	}

	// Ensure stream is false or nil for non-streaming
	if params.Stream != nil && *params.Stream {
		return nil, fmt.Errorf("use CreateStream for streaming responses")
	}

	// Make the request
	req := &http.Request{
		Method: "POST",
		Path:   "/chat/completions",
		Body:   params,
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Parse response
	var result types.StreamChunk
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}

// CreateStream generates a streaming chat completion.
// Returns a Stream that yields StreamChunk objects as they arrive.
// The stream must be closed when done to release resources.
func (s *Service) CreateStream(ctx context.Context, params *CompletionParams) (*Stream, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	// Validate required fields
	if len(params.Messages) == 0 {
		return nil, fmt.Errorf("messages are required")
	}
	if params.Model == "" {
		return nil, fmt.Errorf("model is required")
	}

	// Enable streaming
	streamEnabled := true
	params.Stream = &streamEnabled

	// Make the streaming request
	req := &http.Request{
		Method: "POST",
		Path:   "/chat/completions",
		Body:   params,
	}

	resp, err := s.client.DoStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("streaming request failed: %w", err)
	}

	// Create and return stream
	return newStream(ctx, resp.Response), nil
}
