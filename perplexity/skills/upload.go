package skills

import (
	"context"
	"fmt"
	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
)

func (s *Service) Create(ctx context.Context, file api.Upload, opts ...api.RequestOption) (*ManagedSkill, error) {
	r, err := s.CreateRaw(ctx, file, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (s *Service) CreateRaw(ctx context.Context, file api.Upload, opts ...api.RequestOption) (*api.RawResponse[ManagedSkill], error) {
	return internalhttp.UploadJSON[ManagedSkill](ctx, s.client, "POST", "/v1/skills", file, map[string]any{}, opts...)
}

func (s *Service) Update(ctx context.Context, id, expectedRevision string, file api.Upload, opts ...api.RequestOption) (*ManagedSkill, error) {
	r, err := s.UpdateRaw(ctx, id, expectedRevision, file, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (s *Service) UpdateRaw(ctx context.Context, id, expectedRevision string, file api.Upload, opts ...api.RequestOption) (*api.RawResponse[ManagedSkill], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	if expectedRevision == "" {
		return nil, fmt.Errorf("expected_revision is required")
	}
	options := append([]api.RequestOption{api.WithQuery("expected_revision", expectedRevision)}, opts...)
	return internalhttp.UploadJSON[ManagedSkill](ctx, s.client, "PUT", "/v1/skills/"+segment, file, map[string]any{}, options...)
}

func (s *Service) Delete(ctx context.Context, id, expectedRevision string, opts ...api.RequestOption) error {
	_, err := s.DeleteRaw(ctx, id, expectedRevision, opts...)
	return err
}

func (s *Service) DeleteRaw(ctx context.Context, id, expectedRevision string, opts ...api.RequestOption) (*api.RawResponse[struct{}], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	if expectedRevision == "" {
		return nil, fmt.Errorf("expected_revision is required")
	}
	options := append([]api.RequestOption{api.WithQuery("expected_revision", expectedRevision)}, opts...)
	return internalhttp.ExecuteJSON[struct{}](ctx, s.client, &internalhttp.Request{Method: "DELETE", Path: "/v1/skills/" + segment, Options: api.ApplyRequestOptions(options)})
}

type RevisionsService struct{ client *internalhttp.Client }

func (s *RevisionsService) List(ctx context.Context, id string, opts ...api.RequestOption) (*SkillRevisionList, error) {
	r, err := s.ListRaw(ctx, id, opts...)
	if err != nil {
		return nil, err
	}
	return r.Data, nil
}

func (s *RevisionsService) ListRaw(ctx context.Context, id string, opts ...api.RequestOption) (*api.RawResponse[SkillRevisionList], error) {
	segment, err := internalhttp.PathSegment(id)
	if err != nil {
		return nil, err
	}
	return internalhttp.ExecuteJSON[SkillRevisionList](ctx, s.client, &internalhttp.Request{Method: "GET", Path: "/v1/skills/" + segment + "/revisions", Options: api.ApplyRequestOptions(opts)})
}
