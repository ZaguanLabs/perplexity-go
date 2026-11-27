package monitors

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct {
	client  *internalhttp.Client
	Runs    *RunsService
	Webhook *WebhookService
}
type RunsService struct{ client *internalhttp.Client }
type WebhookService struct{ client *internalhttp.Client }

func NewService(client *internalhttp.Client) *Service {
	return &Service{client: client, Runs: &RunsService{client}, Webhook: &WebhookService{client}}
}

func (s *Service) Create(ctx context.Context, params *CreateMonitorRequest, opts ...api.RequestOption) (*Monitor, error) {
	raw, err := s.CreateRaw(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *Service) CreateRaw(ctx context.Context, params *CreateMonitorRequest, opts ...api.RequestOption) (*api.RawResponse[Monitor], error) {
	if params == nil {
		return nil, fmt.Errorf("params are required")
	}
	req := &internalhttp.Request{Method: http.MethodPost, Path: "/v1/monitors", Options: api.ApplyRequestOptions(opts), Body: params}
	return internalhttp.ExecuteJSON[Monitor](ctx, s.client, req)
}

func (s *Service) Delete(ctx context.Context, monitorID string, opts ...api.RequestOption) (*DeleteMonitorResponse, error) {
	raw, err := s.DeleteRaw(ctx, monitorID, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *Service) DeleteRaw(ctx context.Context, monitorID string, opts ...api.RequestOption) (*api.RawResponse[DeleteMonitorResponse], error) {
	monitorIDPath, err := internalhttp.PathSegment(monitorID)
	if err != nil {
		return nil, err
	}
	req := &internalhttp.Request{Method: http.MethodDelete, Path: "/v1/monitors/" + monitorIDPath, Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[DeleteMonitorResponse](ctx, s.client, req)
}

func (s *Service) List(ctx context.Context, opts ...api.RequestOption) (*MonitorList, error) {
	raw, err := s.ListRaw(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *Service) ListRaw(ctx context.Context, opts ...api.RequestOption) (*api.RawResponse[MonitorList], error) {
	req := &internalhttp.Request{Method: http.MethodGet, Path: "/v1/monitors", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[MonitorList](ctx, s.client, req)
}

func (s *Service) Pause(ctx context.Context, monitorID string, opts ...api.RequestOption) (*Monitor, error) {
	raw, err := s.PauseRaw(ctx, monitorID, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *Service) PauseRaw(ctx context.Context, monitorID string, opts ...api.RequestOption) (*api.RawResponse[Monitor], error) {
	monitorIDPath, err := internalhttp.PathSegment(monitorID)
	if err != nil {
		return nil, err
	}
	req := &internalhttp.Request{Method: http.MethodPost, Path: "/v1/monitors/" + monitorIDPath + "/pause", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[Monitor](ctx, s.client, req)
}

func (s *Service) Resume(ctx context.Context, monitorID string, opts ...api.RequestOption) (*Monitor, error) {
	raw, err := s.ResumeRaw(ctx, monitorID, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *Service) ResumeRaw(ctx context.Context, monitorID string, opts ...api.RequestOption) (*api.RawResponse[Monitor], error) {
	monitorIDPath, err := internalhttp.PathSegment(monitorID)
	if err != nil {
		return nil, err
	}
	req := &internalhttp.Request{Method: http.MethodPost, Path: "/v1/monitors/" + monitorIDPath + "/resume", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[Monitor](ctx, s.client, req)
}

func (s *Service) Retrieve(ctx context.Context, monitorID string, opts ...api.RequestOption) (*Monitor, error) {
	raw, err := s.RetrieveRaw(ctx, monitorID, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *Service) RetrieveRaw(ctx context.Context, monitorID string, opts ...api.RequestOption) (*api.RawResponse[Monitor], error) {
	monitorIDPath, err := internalhttp.PathSegment(monitorID)
	if err != nil {
		return nil, err
	}
	req := &internalhttp.Request{Method: http.MethodGet, Path: "/v1/monitors/" + monitorIDPath, Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[Monitor](ctx, s.client, req)
}

func (s *Service) Run(ctx context.Context, monitorID string, opts ...api.RequestOption) (*MonitorRun, error) {
	raw, err := s.RunRaw(ctx, monitorID, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *Service) RunRaw(ctx context.Context, monitorID string, opts ...api.RequestOption) (*api.RawResponse[MonitorRun], error) {
	monitorIDPath, err := internalhttp.PathSegment(monitorID)
	if err != nil {
		return nil, err
	}
	req := &internalhttp.Request{Method: http.MethodPost, Path: "/v1/monitors/" + monitorIDPath + "/run", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[MonitorRun](ctx, s.client, req)
}

func (s *Service) Update(ctx context.Context, monitorID string, params *UpdateMonitorRequest, opts ...api.RequestOption) (*Monitor, error) {
	raw, err := s.UpdateRaw(ctx, monitorID, params, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *Service) UpdateRaw(ctx context.Context, monitorID string, params *UpdateMonitorRequest, opts ...api.RequestOption) (*api.RawResponse[Monitor], error) {
	monitorIDPath, err := internalhttp.PathSegment(monitorID)
	if err != nil {
		return nil, err
	}
	if params == nil {
		return nil, fmt.Errorf("params are required")
	}
	req := &internalhttp.Request{Method: http.MethodPatch, Path: "/v1/monitors/" + monitorIDPath, Options: api.ApplyRequestOptions(opts), Body: params}
	return internalhttp.ExecuteJSON[Monitor](ctx, s.client, req)
}

func (s *RunsService) List(ctx context.Context, monitorID string, opts ...api.RequestOption) (*MonitorRunList, error) {
	raw, err := s.ListRaw(ctx, monitorID, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *RunsService) ListRaw(ctx context.Context, monitorID string, opts ...api.RequestOption) (*api.RawResponse[MonitorRunList], error) {
	monitorIDPath, err := internalhttp.PathSegment(monitorID)
	if err != nil {
		return nil, err
	}
	req := &internalhttp.Request{Method: http.MethodGet, Path: "/v1/monitors/" + monitorIDPath + "/runs", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[MonitorRunList](ctx, s.client, req)
}

func (s *RunsService) Retrieve(ctx context.Context, monitorID, runID string, opts ...api.RequestOption) (*MonitorRun, error) {
	raw, err := s.RetrieveRaw(ctx, monitorID, runID, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *RunsService) RetrieveRaw(ctx context.Context, monitorID, runID string, opts ...api.RequestOption) (*api.RawResponse[MonitorRun], error) {
	monitorIDPath, err := internalhttp.PathSegment(monitorID)
	if err != nil {
		return nil, err
	}
	runIDPath, err := internalhttp.PathSegment(runID)
	if err != nil {
		return nil, err
	}
	req := &internalhttp.Request{Method: http.MethodGet, Path: "/v1/monitors/" + monitorIDPath + "/runs/" + runIDPath, Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[MonitorRun](ctx, s.client, req)
}

func (s *WebhookService) Test(ctx context.Context, monitorID string, opts ...api.RequestOption) (*MonitorWebhookTest, error) {
	raw, err := s.TestRaw(ctx, monitorID, opts...)
	if err != nil {
		return nil, err
	}
	return raw.Data, nil
}

func (s *WebhookService) TestRaw(ctx context.Context, monitorID string, opts ...api.RequestOption) (*api.RawResponse[MonitorWebhookTest], error) {
	monitorIDPath, err := internalhttp.PathSegment(monitorID)
	if err != nil {
		return nil, err
	}
	req := &internalhttp.Request{Method: http.MethodPost, Path: "/v1/monitors/" + monitorIDPath + "/webhook/test", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[MonitorWebhookTest](ctx, s.client, req)
}
