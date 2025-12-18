package asyncchat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct {
	client *http.Client
}

func NewService(httpClient *http.Client) *Service {
	return &Service{client: httpClient}
}

func (s *Service) Create(ctx context.Context, params *CompletionCreateParams) (*CompletionCreateResponse, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.Request == nil {
		return nil, fmt.Errorf("request is required")
	}
	if len(params.Request.Messages) == 0 {
		return nil, fmt.Errorf("messages are required")
	}
	if params.Request.Model == "" {
		return nil, fmt.Errorf("model is required")
	}

	req := &http.Request{
		Method: "POST",
		Path:   "/async/chat/completions",
		Body:   params,
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var result CompletionCreateResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, nil
}

func (s *Service) List(ctx context.Context) (*CompletionListResponse, error) {
	req := &http.Request{
		Method: "GET",
		Path:   "/async/chat/completions",
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var result CompletionListResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, nil
}

func (s *Service) Get(ctx context.Context, apiRequest string, params *CompletionGetParams) (*CompletionGetResponse, error) {
	if apiRequest == "" {
		return nil, fmt.Errorf("apiRequest is required")
	}

	path := fmt.Sprintf("/async/chat/completions/%s", url.PathEscape(apiRequest))
	if params != nil && params.LocalMode != nil {
		q := url.Values{}
		q.Set("local_mode", fmt.Sprintf("%t", *params.LocalMode))
		path = path + "?" + q.Encode()
	}

	headers := map[string]string{}
	if params != nil {
		if params.XClientEnv != nil {
			headers["x-client-env"] = *params.XClientEnv
		}
		if params.XClientName != nil {
			headers["x-client-name"] = *params.XClientName
		}
		if params.XCreatedAtEpochSeconds != nil {
			headers["x-created-at-epoch-seconds"] = *params.XCreatedAtEpochSeconds
		}
		if params.XRequestTime != nil {
			headers["x-request-time"] = *params.XRequestTime
		}
		if params.XUsageTier != nil {
			headers["x-usage-tier"] = *params.XUsageTier
		}
		if params.XUserID != nil {
			headers["x-user-id"] = *params.XUserID
		}
	}

	req := &http.Request{
		Method:  "GET",
		Path:    path,
		Headers: headers,
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var result CompletionGetResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, nil
}
