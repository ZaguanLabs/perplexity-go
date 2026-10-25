package batches

import (
	"context"
	"fmt"
	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct{ client *internalhttp.Client }

func NewService(client *internalhttp.Client) *Service { return &Service{client: client} }
func (s *Service) Cancel(ctx context.Context, id string, opts ...api.RequestOption) (*Batch, error) {
	r, err := s.CancelRaw(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) CancelRaw(ctx context.Context, id string, opts ...api.RequestOption) (*api.RawResponse[Batch], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	request := &internalhttp.Request{Method: "POST", Path: "/v1/batches/" + segment + "/cancel", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[Batch](ctx, s.client, request)
}
func (s *Service) Create(ctx context.Context, params *BatchCreateRequest, opts ...api.RequestOption) (*Batch, error) {
	r, err := s.CreateRaw(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) CreateRaw(ctx context.Context, params *BatchCreateRequest, opts ...api.RequestOption) (*api.RawResponse[Batch], error) {
	if params == nil {
		return nil, fmt.Errorf("params are required")
	}
	request := &internalhttp.Request{Method: "POST", Path: "/v1/batches", Options: api.ApplyRequestOptions(opts), Body: params}
	return internalhttp.ExecuteJSON[Batch](ctx, s.client, request)
}
func (s *Service) List(ctx context.Context, opts ...api.RequestOption) (*BatchList, error) {
	r, err := s.ListRaw(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) ListRaw(ctx context.Context, opts ...api.RequestOption) (*api.RawResponse[BatchList], error) {
	request := &internalhttp.Request{Method: "GET", Path: "/v1/batches", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[BatchList](ctx, s.client, request)
}
func (s *Service) Retrieve(ctx context.Context, id string, opts ...api.RequestOption) (*Batch, error) {
	r, err := s.RetrieveRaw(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) RetrieveRaw(ctx context.Context, id string, opts ...api.RequestOption) (*api.RawResponse[Batch], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	request := &internalhttp.Request{Method: "GET", Path: "/v1/batches/" + segment + "", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[Batch](ctx, s.client, request)
}
