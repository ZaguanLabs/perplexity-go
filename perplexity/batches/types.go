// Code generated from Perplexity Python SDK 0.46.1 models. DO NOT EDIT.
package batches

type Batch struct {
	CancelledAt      *int               `json:"cancelled_at,omitempty"`
	CancellingAt     *int               `json:"cancelling_at,omitempty"`
	CompletedAt      *int               `json:"completed_at,omitempty"`
	CompletionWindow string             `json:"completion_window"`
	CreatedAt        int                `json:"created_at"`
	Endpoint         string             `json:"endpoint"`
	ErrorFileID      *string            `json:"error_file_id,omitempty"`
	Errors           *BatchErrors       `json:"errors,omitempty"`
	ExpiredAt        *int               `json:"expired_at,omitempty"`
	ExpiresAt        *int               `json:"expires_at,omitempty"`
	FailedAt         *int               `json:"failed_at,omitempty"`
	FinalizingAt     *int               `json:"finalizing_at,omitempty"`
	ID               string             `json:"id"`
	InProgressAt     *int               `json:"in_progress_at,omitempty"`
	InputFileID      string             `json:"input_file_id"`
	Metadata         map[string]string  `json:"metadata,omitempty"`
	Object           string             `json:"object"`
	OutputFileID     *string            `json:"output_file_id,omitempty"`
	RequestCounts    BatchRequestCounts `json:"request_counts"`
	Status           string             `json:"status"`
	Usage            BatchUsage         `json:"usage"`
}
type BatchCreateRequest struct {
	CompletionWindow *string           `json:"completion_window,omitempty"`
	Endpoint         string            `json:"endpoint"`
	InputFileID      string            `json:"input_file_id"`
	Metadata         map[string]string `json:"metadata,omitempty"`
}
type BatchError struct {
	Code    string `json:"code"`
	Line    *int   `json:"line,omitempty"`
	Message string `json:"message"`
}
type BatchErrors struct {
	Data   []BatchError `json:"data"`
	Object string       `json:"object"`
}
type BatchList struct {
	Data    []Batch `json:"data"`
	FirstID *string `json:"first_id,omitempty"`
	HasMore bool    `json:"has_more"`
	LastID  *string `json:"last_id,omitempty"`
	Object  string  `json:"object"`
}
type BatchRequestCounts struct {
	Completed int `json:"completed"`
	Failed    int `json:"failed"`
	Total     int `json:"total"`
}
type BatchUsage struct {
	InputTokens         int                           `json:"input_tokens"`
	InputTokensDetails  BatchUsageInputTokensDetails  `json:"input_tokens_details"`
	OutputTokens        int                           `json:"output_tokens"`
	OutputTokensDetails BatchUsageOutputTokensDetails `json:"output_tokens_details"`
	TotalTokens         int                           `json:"total_tokens"`
}
type BatchUsageInputTokensDetails struct {
	CachedTokens int `json:"cached_tokens"`
}
type BatchUsageOutputTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}
