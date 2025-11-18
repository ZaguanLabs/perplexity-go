package types

// UsageInfo contains token usage and cost information for a completion.
type UsageInfo struct {
	// CompletionTokens is the number of tokens in the completion.
	CompletionTokens int `json:"completion_tokens"`

	// Cost contains detailed cost breakdown.
	Cost Cost `json:"cost"`

	// PromptTokens is the number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens"`

	// TotalTokens is the total number of tokens used.
	TotalTokens int `json:"total_tokens"`

	// CitationTokens is the number of tokens used for citations (optional).
	CitationTokens *int `json:"citation_tokens,omitempty"`

	// NumSearchQueries is the number of search queries performed (optional).
	NumSearchQueries *int `json:"num_search_queries,omitempty"`

	// ReasoningTokens is the number of tokens used for reasoning (optional).
	ReasoningTokens *int `json:"reasoning_tokens,omitempty"`

	// SearchContextSize is the size of the search context (optional).
	SearchContextSize *string `json:"search_context_size,omitempty"`
}

// Cost contains detailed cost breakdown for a completion.
type Cost struct {
	// InputTokensCost is the cost of input tokens.
	InputTokensCost float64 `json:"input_tokens_cost"`

	// OutputTokensCost is the cost of output tokens.
	OutputTokensCost float64 `json:"output_tokens_cost"`

	// TotalCost is the total cost.
	TotalCost float64 `json:"total_cost"`

	// CitationTokensCost is the cost of citation tokens (optional).
	CitationTokensCost *float64 `json:"citation_tokens_cost,omitempty"`

	// ReasoningTokensCost is the cost of reasoning tokens (optional).
	ReasoningTokensCost *float64 `json:"reasoning_tokens_cost,omitempty"`

	// RequestCost is the base request cost (optional).
	RequestCost *float64 `json:"request_cost,omitempty"`

	// SearchQueriesCost is the cost of search queries (optional).
	SearchQueriesCost *float64 `json:"search_queries_cost,omitempty"`
}
