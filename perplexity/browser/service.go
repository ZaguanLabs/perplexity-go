package browser

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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

func (s *SessionsService) Create(ctx context.Context) (*SessionResponse, error) {
	req := &internalhttp.Request{
		Method: http.MethodPost,
		Path:   "/v1/browser/sessions",
		Body:   map[string]any{},
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	var result SessionResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, nil
}

func (s *SessionsService) Delete(ctx context.Context, sessionID string) error {
	if sessionID == "" {
		return fmt.Errorf("sessionID is required")
	}

	req := &internalhttp.Request{
		Method: http.MethodDelete,
		Path:   "/v1/browser/sessions/" + url.PathEscape(sessionID),
		Headers: map[string]string{
			"Accept": "*/*",
		},
	}

	_, err := s.client.Do(ctx, req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	return nil
}
