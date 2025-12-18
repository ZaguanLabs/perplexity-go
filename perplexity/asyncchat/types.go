package asyncchat

import (
	"github.com/ZaguanLabs/perplexity-go/perplexity/chat"
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

type CompletionStatus string

const (
	CompletionStatusCreated    CompletionStatus = "CREATED"
	CompletionStatusInProgress CompletionStatus = "IN_PROGRESS"
	CompletionStatusCompleted  CompletionStatus = "COMPLETED"
	CompletionStatusFailed     CompletionStatus = "FAILED"
)

type CompletionCreateParams struct {
	Request        *chat.CompletionParams `json:"request"`
	IdempotencyKey *string                `json:"idempotency_key,omitempty"`
}

type CompletionResponse struct {
	ID           string             `json:"id"`
	CreatedAt    int64              `json:"created_at"`
	Model        string             `json:"model"`
	Status       CompletionStatus   `json:"status"`
	CompletedAt  *int64             `json:"completed_at,omitempty"`
	ErrorMessage *string            `json:"error_message,omitempty"`
	FailedAt     *int64             `json:"failed_at,omitempty"`
	Response     *types.StreamChunk `json:"response,omitempty"`
	StartedAt    *int64             `json:"started_at,omitempty"`
}

type CompletionCreateResponse = CompletionResponse

type CompletionGetResponse = CompletionResponse

type CompletionListRequest struct {
	ID          string           `json:"id"`
	CreatedAt   int64            `json:"created_at"`
	Model       string           `json:"model"`
	Status      CompletionStatus `json:"status"`
	CompletedAt *int64           `json:"completed_at,omitempty"`
	FailedAt    *int64           `json:"failed_at,omitempty"`
	StartedAt   *int64           `json:"started_at,omitempty"`
}

type CompletionListResponse struct {
	Requests  []CompletionListRequest `json:"requests"`
	NextToken *string                 `json:"next_token,omitempty"`
}

type CompletionGetParams struct {
	LocalMode *bool

	XClientEnv             *string
	XClientName            *string
	XCreatedAtEpochSeconds *string
	XRequestTime           *string
	XUsageTier             *string
	XUserID                *string
}
