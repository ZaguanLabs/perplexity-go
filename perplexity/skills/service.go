package skills

import (
	"context"
	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

type Service struct {
	client    *internalhttp.Client
	Revisions *RevisionsService
}

func NewService(client *internalhttp.Client) *Service {
	return &Service{client: client, Revisions: &RevisionsService{client: client}}
}
func (s *Service) List(ctx context.Context, opts ...api.RequestOption) (*SkillList, error) {
	r, err := s.ListRaw(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) ListRaw(ctx context.Context, opts ...api.RequestOption) (*api.RawResponse[SkillList], error) {
	request := &internalhttp.Request{Method: "GET", Path: "/v1/skills", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[SkillList](ctx, s.client, request)
}
func (s *Service) Retrieve(ctx context.Context, id string, opts ...api.RequestOption) (*ManagedSkill, error) {
	r, err := s.RetrieveRaw(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) RetrieveRaw(ctx context.Context, id string, opts ...api.RequestOption) (*api.RawResponse[ManagedSkill], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	request := &internalhttp.Request{Method: "GET", Path: "/v1/skills/" + segment + "", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[ManagedSkill](ctx, s.client, request)
}
func (s *Service) Download(ctx context.Context, id string, opts ...api.RequestOption) (*SkillDownload, error) {
	r, err := s.DownloadRaw(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}
func (s *Service) DownloadRaw(ctx context.Context, id string, opts ...api.RequestOption) (*api.RawResponse[SkillDownload], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	request := &internalhttp.Request{Method: "GET", Path: "/v1/skills/" + segment + "/download", Options: api.ApplyRequestOptions(opts)}
	return internalhttp.ExecuteJSON[SkillDownload](ctx, s.client, request)
}
