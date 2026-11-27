package webcontext

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct{ client *internalhttp.Client }

func NewService(client *internalhttp.Client) *Service { return &Service{client} }
func (s *Service) Create(ctx context.Context, params *ContextRequest, opts ...api.RequestOption) (*ContextResponse, error) {
	raw, err := s.CreateRaw(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *Service) CreateRaw(ctx context.Context, params *ContextRequest, opts ...api.RequestOption) (*api.RawResponse[ContextResponse], error) {
	if params == nil {
		return nil, fmt.Errorf("params are required")
	}
	req := &internalhttp.Request{Method: http.MethodPost, Path: "/v1/context", Options: api.ApplyRequestOptions(opts), Body: params}
	return internalhttp.ExecuteJSON[ContextResponse](ctx, s.client, req)
}
