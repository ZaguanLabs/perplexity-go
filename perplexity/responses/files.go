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

type FilesService struct {
	client *internalhttp.Client
}

func NewFilesService(httpClient *internalhttp.Client) *FilesService {
	return &FilesService{client: httpClient}
}

func (s *FilesService) List(ctx context.Context, responseID string, opts ...api.RequestOption) (*ResponseFileList, error) {
	result, _, err := s.listWithResponse(ctx, responseID, opts...)
	return result, err
}

func (s *FilesService) ListRaw(ctx context.Context, responseID string, opts ...api.RequestOption) (*api.RawResponse[ResponseFileList], error) {
	result, raw, err := s.listWithResponse(ctx, responseID, opts...)
	if err != nil {
		return nil, err
	}
	return &api.RawResponse[ResponseFileList]{Data: result, StatusCode: raw.StatusCode, Headers: raw.Headers, Body: raw.Body, RequestID: raw.RequestID}, nil
}

func (s *FilesService) listWithResponse(ctx context.Context, responseID string, opts ...api.RequestOption) (*ResponseFileList, *internalhttp.Response, error) {
	if responseID == "" || responseID == "." || responseID == ".." {
		return nil, nil, fmt.Errorf("responseID is required")
	}
	req := &internalhttp.Request{Method: http.MethodGet, Path: "/v1/responses/" + url.PathEscape(responseID) + "/files", Options: api.ApplyRequestOptions(opts)}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("request failed: %w", err)
	}
	var result ResponseFileList
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, nil, fmt.Errorf("failed to parse response: %w", err)
	}
	return &result, resp, nil
}

func (s *FilesService) Content(ctx context.Context, responseID, fileID string, opts ...api.RequestOption) ([]byte, error) {
	raw, err := s.ContentRaw(ctx, responseID, fileID, opts...)
	if err != nil {
		return nil, err
	}
	return *raw.Data, nil
}

func (s *FilesService) ContentRaw(ctx context.Context, responseID, fileID string, opts ...api.RequestOption) (*api.RawResponse[[]byte], error) {
	if responseID == "" || responseID == "." || responseID == ".." {
		return nil, fmt.Errorf("responseID is required")
	}
	if fileID == "" || fileID == "." || fileID == ".." {
		return nil, fmt.Errorf("fileID is required")
	}
	allOptions := append([]api.RequestOption{api.WithHeader("Accept", "application/octet-stream")}, opts...)
	req := &internalhttp.Request{
		Method:  http.MethodGet,
		Path:    "/v1/responses/" + url.PathEscape(responseID) + "/files/" + url.PathEscape(fileID) + "/content",
		Options: api.ApplyRequestOptions(allOptions),
	}
	resp, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	body := append([]byte(nil), resp.Body...)
	return &api.RawResponse[[]byte]{Data: &body, StatusCode: resp.StatusCode, Headers: resp.Headers, Body: body, RequestID: resp.RequestID}, nil
}
