package contextualizedembeddings

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
	if params.Model == "" {
		return nil, nil, fmt.Errorf("model is required")
	}
	if len(params.Input) == 0 {
		return nil, nil, fmt.Errorf("input cannot be empty")
	}
	for i, document := range params.Input {
		if len(document) == 0 {
			return nil, nil, fmt.Errorf("input[%d] cannot be empty", i)
		}
		for j, chunk := range document {
			if chunk == "" {
				return nil, nil, fmt.Errorf("input[%d][%d] cannot be empty", i, j)
			}
		}
	}

	req := &internalhttp.Request{
		Method:  http.MethodPost,
		Path:    "/v1/contextualizedembeddings",
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
