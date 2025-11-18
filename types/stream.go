package types

// StreamChunk represents a chunk in a streaming response.
type StreamChunk struct {
	// ID is the unique identifier for the completion.
	ID string `json:"id"`

	// Choices contains the completion choices.
	Choices []Choice `json:"choices"`

	// Created is the Unix timestamp of when the completion was created.
	Created int64 `json:"created"`

	// Model is the model used for the completion.
	Model string `json:"model"`

	// Citations contains citation URLs (optional).
	Citations []string `json:"citations,omitempty"`

	// Object is the object type (optional).
	Object *string `json:"object,omitempty"`

	// SearchResults contains search results (optional).
	SearchResults []SearchResult `json:"search_results,omitempty"`

	// Status indicates the completion status (optional).
	Status *CompletionStatus `json:"status,omitempty"`

	// Type indicates the type of stream chunk (optional).
	Type *StreamChunkType `json:"type,omitempty"`

	// Usage contains token usage information (optional).
	Usage *UsageInfo `json:"usage,omitempty"`
}

// CompletionStatus represents the status of a completion.
type CompletionStatus string

const (
	// CompletionStatusPending means the completion is still in progress.
	CompletionStatusPending CompletionStatus = "PENDING"

	// CompletionStatusCompleted means the completion is finished.
	CompletionStatusCompleted CompletionStatus = "COMPLETED"
)

// StreamChunkType represents the type of stream chunk.
type StreamChunkType string

const (
	// StreamChunkTypeMessage indicates a message chunk.
	StreamChunkTypeMessage StreamChunkType = "message"

	// StreamChunkTypeInfo indicates an info chunk.
	StreamChunkTypeInfo StreamChunkType = "info"

	// StreamChunkTypeEndOfStream indicates the end of the stream.
	StreamChunkTypeEndOfStream StreamChunkType = "end_of_stream"
)
