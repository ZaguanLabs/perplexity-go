package search

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

// Service provides access to search APIs.
type Service struct {
	client *http.Client
}

// NewService creates a new search service.
func NewService(httpClient *http.Client) *Service {
	return &Service{
		client: httpClient,
	}
}

// Create performs a web search and returns relevant results.
func (s *Service) Create(ctx context.Context, params *SearchParams) (*types.SearchResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	// Validate required fields
	if params.Query == nil {
		return nil, fmt.Errorf("query is required")
	}

	// Validate query type
	switch q := params.Query.(type) {
	case string:
		if q == "" {
			return nil, fmt.Errorf("query cannot be empty")
		}
	case []string:
		if len(q) == 0 {
			return nil, fmt.Errorf("query cannot be empty")
		}
		for i, query := range q {
			if query == "" {
				return nil, fmt.Errorf("query[%d] cannot be empty", i)
			}
		}
	default:
		return nil, fmt.Errorf("query must be string or []string, got %T", params.Query)
	}

	// Make the request
	req := &http.Request{
		Method: "POST",
		Path:   "/search",
		Body:   params,
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Parse response
	var result types.SearchResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}
