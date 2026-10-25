package responses

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
	client *internalhttp.Client
	Files  *FilesService
}

func NewService(httpClient *internalhttp.Client) *Service {
	return &Service{client: httpClient, Files: NewFilesService(httpClient)}
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

func (s *Service) Retrieve(ctx context.Context, responseID string, opts ...api.RequestOption) (*RetrieveResponse, error) {
	result, _, err := s.retrieveWithResponse(ctx, responseID, opts...)
	return result, err
}

func (s *Service) RetrieveRaw(ctx context.Context, responseID string, opts ...api.RequestOption) (*api.RawResponse[RetrieveResponse], error) {
	result, raw, err := s.retrieveWithResponse(ctx, responseID, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[RetrieveResponse]{
		Data:       result,
		StatusCode: raw.StatusCode,
		Headers:    raw.Headers,
		Body:       raw.Body,
		RequestID:  raw.RequestID,
	}, nil
}

func (s *Service) retrieveWithResponse(ctx context.Context, responseID string, opts ...api.RequestOption) (*RetrieveResponse, *internalhttp.Response, error) {
	if responseID == "" || responseID == "." || responseID == ".." {
		return nil, nil, fmt.Errorf("responseID is required")
	}

	req := &internalhttp.Request{
		Method:  http.MethodGet,
		Path:    "/v1/responses/" + url.PathEscape(responseID),
		Options: api.ApplyRequestOptions(opts),
	}

	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}

	var result RetrieveResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, resp, nil
}

func (s *Service) Cancel(ctx context.Context, responseID string, opts ...api.RequestOption) (*CancelResponse, error) {
	result, _, err := s.cancelWithResponse(ctx, responseID, opts...)
	return result, err
}

func (s *Service) CancelRaw(ctx context.Context, responseID string, opts ...api.RequestOption) (*api.RawResponse[CancelResponse], error) {
	result, raw, err := s.cancelWithResponse(ctx, responseID, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[CancelResponse]{
		Data: result, StatusCode: raw.StatusCode, Headers: raw.Headers, Body: raw.Body, RequestID: raw.RequestID,
	}, nil
}

func (s *Service) cancelWithResponse(ctx context.Context, responseID string, opts ...api.RequestOption) (*CancelResponse, *internalhttp.Response, error) {
	if responseID == "" || responseID == "." || responseID == ".." {
		return nil, nil, fmt.Errorf("responseID is required")
	}
	req := &internalhttp.Request{
		Method:  http.MethodPost,
		Path:    "/v1/responses/" + url.PathEscape(responseID) + "/cancel",
		Options: api.ApplyRequestOptions(opts),
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}
	var result CancelResponse
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, resp, nil
}
