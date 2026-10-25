package responses

import (
	"context"
	"fmt"

	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type RetrieveParams struct {
	Format        *string
	StartingAfter *int64
}

func retrieveOptions(params *RetrieveParams, streaming bool, opts []api.RequestOption) (api.RequestOptions, error) {
	options := api.ApplyRequestOptions(opts)
	if options.Query == nil {
		options.Query = map[string]any{}
	}
	for _, key := range []string{"format", "starting_after", "stream"} {
		delete(options.Query, key)
	}
	if params != nil {
		if params.Format != nil {
			options.Query["format"] = *params.Format
		}
		if params.StartingAfter != nil {
			if *params.StartingAfter < 0 || *params.StartingAfter > 9007199254740991 {
				return options, fmt.Errorf("starting_after must be a non-negative safe integer")
			}
			options.Query["starting_after"] = *params.StartingAfter
		}
	}
	if streaming && (params == nil || params.StartingAfter == nil) {
		return options, fmt.Errorf("streaming retrieval requires starting_after")
	}
	options.Query["stream"] = streaming
	return options, nil
}

func (s *Service) RetrieveWithParams(ctx context.Context, id string, params *RetrieveParams, opts ...api.RequestOption) (*RetrieveResponse, error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	options, err := retrieveOptions(params, false, opts)
	if err != nil {
		return nil, err
	}
	r, err := internalhttp.ExecuteJSON[RetrieveResponse](ctx, s.client, &internalhttp.Request{Method: "GET", Path: "/v1/responses/" + segment, Options: options})
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (s *Service) RetrieveStream(ctx context.Context, id string, params *RetrieveParams, opts ...api.RequestOption) (*Stream, error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	options, err := retrieveOptions(params, true, opts)
	if err != nil {
		return nil, err
	}
	r, err := s.client.DoStream(ctx, &internalhttp.Request{Method: "GET", Path: "/v1/responses/" + segment, Options: options})
	if err != nil {
		return nil, err
	}
	return newStream(ctx, r.Response), nil
}
