package types

// Choice represents a completion choice.
type Choice struct {
	// Delta contains the incremental message update (for streaming).
	Delta ChatMessage `json:"delta"`

	// Index is the index of this choice in the list.
	Index int `json:"index"`

	// Message contains the complete message (for non-streaming).
	Message ChatMessage `json:"message"`

	// FinishReason indicates why the completion finished (optional).
	FinishReason *FinishReason `json:"finish_reason,omitempty"`
}

// FinishReason indicates why a completion finished.
type FinishReason string

const (
	// FinishReasonStop means the model hit a natural stop point or provided stop sequence.
	FinishReasonStop FinishReason = "stop"

	// FinishReasonLength means the maximum token limit was reached.
	FinishReasonLength FinishReason = "length"
)
