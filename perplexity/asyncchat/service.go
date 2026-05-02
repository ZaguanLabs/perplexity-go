package asyncchat

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	"github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct {
	client *http.Client
}

func NewService(httpClient *http.Client) *Service {
	return &Service{client: httpClient}
}

func (s *Service) Create(ctx context.Context, params *CompletionCreateParams, opts ...api.RequestOption) (*CompletionCreateResponse, error) {
	result, _, err := s.createWithResponse(ctx, params, opts...)
	return result, err
}

func (s *Service) CreateRaw(ctx context.Context, params *CompletionCreateParams, opts ...api.RequestOption) (*api.RawResponse[CompletionCreateResponse], error) {
	result, raw, err := s.createWithResponse(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[CompletionCreateResponse]{
		Data:       result,
		StatusCode: raw.StatusCode,
		Headers:    raw.Headers,
		Body:       raw.Body,
		RequestID:  raw.RequestID,
	}, nil
}

func (s *Service) createWithResponse(ctx context.Context, params *CompletionCreateParams, opts ...api.RequestOption) (*CompletionCreateResponse, *http.Response, error) {
	if params == nil {
		return nil, nil, fmt.Errorf("params cannot be nil")
	}
	if params.Request == nil {
		return nil, nil, fmt.Errorf("request is required")
	}
	if len(params.Request.Messages) == 0 {
		return nil, nil, fmt.Errorf("messages are required")
	}
	if params.Request.Model == "" {
		return nil, nil, fmt.Errorf("model is required")
	}

	req := &http.Request{
		Method:  "POST",
		Path:    "/async/chat/completions",
		Body:    params,
		Options: api.ApplyRequestOptions(opts),
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}

	var result CompletionCreateResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, resp, nil
}

func (s *Service) List(ctx context.Context, opts ...api.RequestOption) (*CompletionListResponse, error) {
	result, _, err := s.listWithResponse(ctx, opts...)
	return result, err
}

func (s *Service) ListRaw(ctx context.Context, opts ...api.RequestOption) (*api.RawResponse[CompletionListResponse], error) {
	result, raw, err := s.listWithResponse(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[CompletionListResponse]{
		Data:       result,
		StatusCode: raw.StatusCode,
		Headers:    raw.Headers,
		Body:       raw.Body,
		RequestID:  raw.RequestID,
	}, nil
}

func (s *Service) listWithResponse(ctx context.Context, opts ...api.RequestOption) (*CompletionListResponse, *http.Response, error) {
	req := &http.Request{
		Method:  "GET",
		Path:    "/async/chat/completions",
		Options: api.ApplyRequestOptions(opts),
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}

	var result CompletionListResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, resp, nil
}

func (s *Service) Get(ctx context.Context, requestID string, params *CompletionGetParams, opts ...api.RequestOption) (*CompletionGetResponse, error) {
	result, _, err := s.getWithResponse(ctx, requestID, params, opts...)
	return result, err
}

func (s *Service) GetRaw(ctx context.Context, requestID string, params *CompletionGetParams, opts ...api.RequestOption) (*api.RawResponse[CompletionGetResponse], error) {
	result, raw, err := s.getWithResponse(ctx, requestID, params, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[CompletionGetResponse]{
		Data:       result,
		StatusCode: raw.StatusCode,
		Headers:    raw.Headers,
		Body:       raw.Body,
		RequestID:  raw.RequestID,
	}, nil
}

func (s *Service) getWithResponse(ctx context.Context, requestID string, params *CompletionGetParams, opts ...api.RequestOption) (*CompletionGetResponse, *http.Response, error) {
	if requestID == "" {
		return nil, nil, fmt.Errorf("requestID is required")
	}

	path := fmt.Sprintf("/async/chat/completions/%s", url.PathEscape(requestID))
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
		Options: api.ApplyRequestOptions(opts),
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}

	var result CompletionGetResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, resp, nil
}
