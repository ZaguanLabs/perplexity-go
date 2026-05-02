package search

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
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
func (s *Service) Create(ctx context.Context, params *SearchParams, opts ...api.RequestOption) (*types.SearchResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}

	if params.Query.Text == nil && len(params.Query.Texts) == 0 {
		return nil, fmt.Errorf("query is required")
	}

	if params.Query.Text != nil {
		if *params.Query.Text == "" {
			return nil, fmt.Errorf("query cannot be empty")
		}
	} else {
		if len(params.Query.Texts) == 0 {
			return nil, fmt.Errorf("query cannot be empty")
		}
		for i, query := range params.Query.Texts {
			if query == "" {
				return nil, fmt.Errorf("query[%d] cannot be empty", i)
			}
		}
	}

	req := &http.Request{
		Method:  "POST",
		Path:    "/search",
		Body:    params,
		Options: api.ApplyRequestOptions(opts),
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

func (s *Service) CreateRaw(ctx context.Context, params *SearchParams, opts ...api.RequestOption) (*api.RawResponse[types.SearchResponse], error) {
	result, raw, err := s.createWithResponse(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[types.SearchResponse]{
		Data:       result,
		StatusCode: raw.StatusCode,
		Headers:    raw.Headers,
		Body:       raw.Body,
		RequestID:  raw.RequestID,
	}, nil
}

func (s *Service) createWithResponse(ctx context.Context, params *SearchParams, opts ...api.RequestOption) (*types.SearchResponse, *http.Response, error) {
	if params == nil {
		return nil, nil, fmt.Errorf("params cannot be nil")
	}
	if params.Query.Text == nil && len(params.Query.Texts) == 0 {
		return nil, nil, fmt.Errorf("query is required")
	}
	if params.Query.Text != nil {
		if *params.Query.Text == "" {
			return nil, nil, fmt.Errorf("query cannot be empty")
		}
	} else {
		if len(params.Query.Texts) == 0 {
			return nil, nil, fmt.Errorf("query cannot be empty")
		}
		for i, query := range params.Query.Texts {
			if query == "" {
				return nil, nil, fmt.Errorf("query[%d] cannot be empty", i)
			}
		}
	}
	req := &http.Request{
		Method:  "POST",
		Path:    "/search",
		Body:    params,
		Options: api.ApplyRequestOptions(opts),
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}
	var result types.SearchResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, resp, nil
}
