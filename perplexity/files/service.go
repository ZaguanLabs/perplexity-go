package files

import (
	"context"
	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct{ client *internalhttp.Client }

func NewService(client *internalhttp.Client) *Service { return &Service{client: client} }
func (s *Service) Delete(ctx context.Context, id string, opts ...api.RequestOption) (*FileDeleted, error) {
	r, err := s.DeleteRaw(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) DeleteRaw(ctx context.Context, id string, opts ...api.RequestOption) (*api.RawResponse[FileDeleted], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	request := &internalhttp.Request{Method: "DELETE", Path: "/v1/files/" + segment + "", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[FileDeleted](ctx, s.client, request)
}
func (s *Service) List(ctx context.Context, opts ...api.RequestOption) (*FileObjectList, error) {
	r, err := s.ListRaw(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) ListRaw(ctx context.Context, opts ...api.RequestOption) (*api.RawResponse[FileObjectList], error) {
	request := &internalhttp.Request{Method: "GET", Path: "/v1/files", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[FileObjectList](ctx, s.client, request)
}
func (s *Service) Retrieve(ctx context.Context, id string, opts ...api.RequestOption) (*FileObject, error) {
	r, err := s.RetrieveRaw(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) RetrieveRaw(ctx context.Context, id string, opts ...api.RequestOption) (*api.RawResponse[FileObject], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	request := &internalhttp.Request{Method: "GET", Path: "/v1/files/" + segment + "", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[FileObject](ctx, s.client, request)
}
