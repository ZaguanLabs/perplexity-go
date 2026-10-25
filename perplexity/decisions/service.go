package decisions

import (
	"context"
	"fmt"
	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct{ client *internalhttp.Client }

func NewService(client *internalhttp.Client) *Service { return &Service{client: client} }
func (s *Service) Create(ctx context.Context, params *DecisionsRequest, opts ...api.RequestOption) (*DecisionsResponse, error) {
	r, err := s.CreateRaw(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) CreateRaw(ctx context.Context, params *DecisionsRequest, opts ...api.RequestOption) (*api.RawResponse[DecisionsResponse], error) {
	if params == nil {
		return nil, fmt.Errorf("params are required")
	}
	request := &internalhttp.Request{Method: "POST", Path: "/v1/decisions", Options: api.ApplyRequestOptions(opts), Body: params}
	return internalhttp.ExecuteJSON[DecisionsResponse](ctx, s.client, request)
}
