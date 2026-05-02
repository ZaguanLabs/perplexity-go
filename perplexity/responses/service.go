package responses

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct {
	client *internalhttp.Client
}

func NewService(httpClient *internalhttp.Client) *Service {
	return &Service{client: httpClient}
}

func (s *Service) Create(ctx context.Context, params *CreateParams, opts ...api.RequestOption) (*CreateResponse, error) {
	result, _, err := s.createWithResponse(ctx, params, opts...)
	return result, err
}

func (s *Service) CreateRaw(ctx context.Context, params *CreateParams, opts ...api.RequestOption) (*api.RawResponse[CreateResponse], error) {
	result, raw, err := s.createWithResponse(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[CreateResponse]{
		Data:       result,
		StatusCode: raw.StatusCode,
		Headers:    raw.Headers,
		Body:       raw.Body,
		RequestID:  raw.RequestID,
	}, nil
}

func (s *Service) createWithResponse(ctx context.Context, params *CreateParams, opts ...api.RequestOption) (*CreateResponse, *internalhttp.Response, error) {
	if params == nil {
		return nil, nil, fmt.Errorf("params cannot be nil")
	}
	if params.Input.Text == nil && len(params.Input.Items) == 0 {
		return nil, nil, fmt.Errorf("input is required")
	}
	if params.Stream != nil && *params.Stream {
		return nil, nil, fmt.Errorf("use CreateStream for streaming responses")
	}

	req := &internalhttp.Request{
		Method:  http.MethodPost,
		Path:    "/v1/responses",
		Body:    params,
		Options: api.ApplyRequestOptions(opts),
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}

	var result CreateResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, resp, nil
}

func (s *Service) CreateStream(ctx context.Context, params *CreateParams, opts ...api.RequestOption) (*Stream, error) {
	if params == nil {
		return nil, fmt.Errorf("params cannot be nil")
	}
	if params.Input.Text == nil && len(params.Input.Items) == 0 {
		return nil, fmt.Errorf("input is required")
	}

	streamEnabled := true
	params.Stream = &streamEnabled

	req := &internalhttp.Request{
		Method:  http.MethodPost,
		Path:    "/v1/responses",
		Body:    params,
		Options: api.ApplyRequestOptions(opts),
	}

	resp, err := s.client.DoStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("streaming request failed: %w", err)
	}

	return newStream(ctx, resp.Response), nil
}
