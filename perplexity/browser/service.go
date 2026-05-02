package browser

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct {
	Sessions *SessionsService
}

type SessionsService struct {
	client *internalhttp.Client
}

func NewService(httpClient *internalhttp.Client) *Service {
	return &Service{
		Sessions: &SessionsService{client: httpClient},
	}
}

func (s *SessionsService) Create(ctx context.Context, opts ...api.RequestOption) (*SessionResponse, error) {
	result, _, err := s.createWithResponse(ctx, opts...)
	return result, err
}

func (s *SessionsService) CreateRaw(ctx context.Context, opts ...api.RequestOption) (*api.RawResponse[SessionResponse], error) {
	result, raw, err := s.createWithResponse(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[SessionResponse]{
		Data:       result,
		StatusCode: raw.StatusCode,
		Headers:    raw.Headers,
		Body:       raw.Body,
		RequestID:  raw.RequestID,
	}, nil
}

func (s *SessionsService) createWithResponse(ctx context.Context, opts ...api.RequestOption) (*SessionResponse, *internalhttp.Response, error) {
	req := &internalhttp.Request{
		Method:  http.MethodPost,
		Path:    "/v1/browser/sessions",
		Body:    map[string]any{},
		Options: api.ApplyRequestOptions(opts),
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}

	var result SessionResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, resp, nil
}

func (s *SessionsService) Delete(ctx context.Context, sessionID string, opts ...api.RequestOption) error {
	_, err := s.deleteWithResponse(ctx, sessionID, opts...)
	return err
}

func (s *SessionsService) DeleteRaw(ctx context.Context, sessionID string, opts ...api.RequestOption) (*api.RawResponse[struct{}], error) {
	raw, err := s.deleteWithResponse(ctx, sessionID, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[struct{}]{
		Data:       &struct{}{},
		StatusCode: raw.StatusCode,
		Headers:    raw.Headers,
		Body:       raw.Body,
		RequestID:  raw.RequestID,
	}, nil
}

func (s *SessionsService) deleteWithResponse(ctx context.Context, sessionID string, opts ...api.RequestOption) (*internalhttp.Response, error) {
	if sessionID == "" {
		return nil, fmt.Errorf("sessionID is required")
	}

	req := &internalhttp.Request{
		Method: http.MethodDelete,
		Path:   "/v1/browser/sessions/" + url.PathEscape(sessionID),
		Headers: map[string]string{
			"Accept": "*/*",
		},
		Options: api.ApplyRequestOptions(opts),
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}
