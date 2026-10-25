package models

import (
	"context"
	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct{ client *internalhttp.Client }

func NewService(client *internalhttp.Client) *Service { return &Service{client: client} }
func (s *Service) List(ctx context.Context, opts ...api.RequestOption) (*ModelList, error) {
	r, err := s.ListRaw(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) ListRaw(ctx context.Context, opts ...api.RequestOption) (*api.RawResponse[ModelList], error) {
	request := &internalhttp.Request{Method: "GET", Path: "/v1/models", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[ModelList](ctx, s.client, request)
}
