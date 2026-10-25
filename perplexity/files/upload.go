package files

import (
	"context"
	"fmt"
	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type CreateParams struct {
	File    api.Upload
	Purpose string
}

func (s *Service) Create(ctx context.Context, params *CreateParams, opts ...api.RequestOption) (*FileObject, error) {
	r, err := s.CreateRaw(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (s *Service) CreateRaw(ctx context.Context, params *CreateParams, opts ...api.RequestOption) (*api.RawResponse[FileObject], error) {
	if params == nil || params.Purpose != "batch" {
		return nil, fmt.Errorf("file purpose must be batch")
	}
	return internalhttp.UploadJSON[FileObject](ctx, s.client, "POST", "/v1/files", params.File, map[string]any{"purpose": params.Purpose}, opts...)
}

func (s *Service) Content(ctx context.Context, id string, opts ...api.RequestOption) ([]byte, error) {
	r, err := s.ContentRaw(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return *r.Data, nil
}

func (s *Service) ContentRaw(ctx context.Context, id string, opts ...api.RequestOption) (*api.RawResponse[[]byte], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	r, err := s.client.Do(ctx, &internalhttp.Request{Method: "GET", Path: "/v1/files/" + segment + "/content", Headers: map[string]string{"Accept": "application/octet-stream"}, Options: api.ApplyRequestOptions(opts)})
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[[]byte]{Data: &r.Body, StatusCode: r.StatusCode, Headers: r.Headers, Body: r.Body, RequestID: r.RequestID}, nil
}
