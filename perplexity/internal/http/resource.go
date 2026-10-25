package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
)

// PathSegment validates and escapes identifiers according to the Python SDK.
func PathSegment(id string) (string, error) {
	if id == "" || id == "." || id == ".." {
		return "", fmt.Errorf("invalid path identifier %q", id)
	}
	return url.PathEscape(id), nil
}

// ExecuteJSON is the shared typed/raw response path for the management APIs.
func ExecuteJSON[T any](ctx context.Context, client *Client, request *Request) (*api.RawResponse[T], error) {
	response, err := client.Do(ctx, request)
	if err != nil {
		return nil, err
	}
	var value T
	if len(response.Body) > 0 {
		if err := json.Unmarshal(response.Body, &value); err != nil {
			return nil, fmt.Errorf("decode response: %w", err)
		}
	}
	return &api.RawResponse[T]{Data: &value, StatusCode: response.StatusCode, Headers: response.Headers, Body: response.Body, RequestID: response.RequestID}, nil
}

func UploadJSON[T any](ctx context.Context, client *Client, method, path string, file api.Upload, fields map[string]any, opts ...api.RequestOption) (*api.RawResponse[T], error) {
	options := api.ApplyRequestOptions(opts)
	for key, value := range options.ExtraBody {
		fields[key] = value
	}
	body, contentType, err := api.Multipart(file, fields)
	if err != nil {
		return nil, err
	}
	return ExecuteJSON[T](ctx, client, &Request{Method: method, Path: path, RawBody: body, Headers: map[string]string{"Content-Type": contentType}, Options: options})
}
