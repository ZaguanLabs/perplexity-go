package chat

import (
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

// CompletionParams contains parameters for creating a chat completion.
type CompletionParams struct {
	// Messages is the list of messages in the conversation (required).
	Messages []types.ChatMessage `json:"messages"`

	// Model is the model to use for completion (required).
	Model string `json:"model"`

	// Stream enables streaming responses via Server-Sent Events (optional).
	Stream *bool `json:"stream,omitempty"`

	// MaxTokens is the maximum number of tokens to generate (optional).
	MaxTokens *int `json:"max_tokens,omitempty"`

	// Temperature controls randomness in the output (0.0 to 2.0) (optional).
	Temperature *float64 `json:"temperature,omitempty"`

	// TopP controls nucleus sampling (optional).
	TopP *float64 `json:"top_p,omitempty"`

	// FrequencyPenalty penalizes frequent tokens (optional).
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`

	// PresencePenalty penalizes tokens that have appeared (optional).
	PresencePenalty *float64 `json:"presence_penalty,omitempty"`

	// Stop sequences where the API will stop generating (optional).
	Stop *Stop `json:"stop,omitempty"`

	// Tools available for the model to call (optional).
	Tools []types.Tool `json:"tools,omitempty"`

	// ToolChoice controls how the model chooses tools (optional).
	ToolChoice *types.ToolChoice `json:"tool_choice,omitempty"`

	// WebSearchOptions configures web search behavior (optional).
	WebSearchOptions *WebSearchOptions `json:"web_search_options,omitempty"`

	// SearchDomainFilter limits search to specific domains (optional).
	SearchDomainFilter []string `json:"search_domain_filter,omitempty"`

	// SearchRecencyFilter filters by recency (optional).
	SearchRecencyFilter *SearchRecencyFilter `json:"search_recency_filter,omitempty"`

	// SearchMode specifies the search mode (optional).
	SearchMode *SearchMode `json:"search_mode,omitempty"`

	// ReturnImages includes images in the response (optional).
	ReturnImages *bool `json:"return_images,omitempty"`

	// ReturnRelatedQuestions includes related questions (optional).
	ReturnRelatedQuestions *bool `json:"return_related_questions,omitempty"`

	// ReasoningEffort controls reasoning depth (optional).
	ReasoningEffort *ReasoningEffort `json:"reasoning_effort,omitempty"`

	// ResponseFormat specifies the response format (optional).
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`

	// Country for search localization (optional).
	Country *string `json:"country,omitempty"`

	// DisableSearch disables web search (optional).
	DisableSearch *bool `json:"disable_search,omitempty"`

	// SearchAfterDateFilter filters search results after date (optional).
	SearchAfterDateFilter *string `json:"search_after_date_filter,omitempty"`

	// SearchBeforeDateFilter filters search results before date (optional).
	SearchBeforeDateFilter *string `json:"search_before_date_filter,omitempty"`

	// SearchLanguageFilter filters by language (optional).
	SearchLanguageFilter []string `json:"search_language_filter,omitempty"`

	// NumSearchResults specifies number of search results (optional).
	NumSearchResults *int `json:"num_search_results,omitempty"`

	// SafeSearch enables safe search filtering (optional).
	SafeSearch *bool `json:"safe_search,omitempty"`

	// Latitude for location-based search (optional).
	Latitude *float64 `json:"latitude,omitempty"`

	// Longitude for location-based search (optional).
	Longitude *float64 `json:"longitude,omitempty"`

	// N is the number of completions to generate (optional).
	N *int `json:"n,omitempty"`

	// TopK for top-k sampling (optional).
	TopK *int `json:"top_k,omitempty"`

	// Logprobs returns log probabilities (optional).
	Logprobs *bool `json:"logprobs,omitempty"`

	// TopLogprobs specifies number of top logprobs to return (optional).
	TopLogprobs *int `json:"top_logprobs,omitempty"`

	// ResponseMetadata for additional metadata (optional).
	ResponseMetadata map[string]interface{} `json:"response_metadata,omitempty"`

	// ThreadID for conversation threading (optional).
	ThreadID *string `json:"thread_id,omitempty"`

	// UseThreads enables threading (optional).
	UseThreads *bool `json:"use_threads,omitempty"`

	// ParallelToolCalls enables parallel tool execution (optional).
	ParallelToolCalls *bool `json:"parallel_tool_calls,omitempty"`

	// ImageDomainFilter filters images by domain (optional).
	ImageDomainFilter []string `json:"image_domain_filter,omitempty"`

	// ImageFormatFilter filters images by format (optional).
	ImageFormatFilter []string `json:"image_format_filter,omitempty"`

	// NumImages specifies number of images to return (optional).
	NumImages *int `json:"num_images,omitempty"`

	// LastUpdatedAfterFilter filters by last updated date (optional).
	LastUpdatedAfterFilter *string `json:"last_updated_after_filter,omitempty"`

	// LastUpdatedBeforeFilter filters by last updated date (optional).
	LastUpdatedBeforeFilter *string `json:"last_updated_before_filter,omitempty"`

	// UpdatedAfterTimestamp filters by timestamp (optional).
	UpdatedAfterTimestamp *int64 `json:"updated_after_timestamp,omitempty"`

	// UpdatedBeforeTimestamp filters by timestamp (optional).
	UpdatedBeforeTimestamp *int64 `json:"updated_before_timestamp,omitempty"`

	// RankingModel specifies the ranking model (optional).
	RankingModel *string `json:"ranking_model,omitempty"`

	// SearchTenant for multi-tenant search (optional).
	SearchTenant *string `json:"search_tenant,omitempty"`

	// SearchInternalProperties for internal search config (optional).
	SearchInternalProperties map[string]interface{} `json:"search_internal_properties,omitempty"`

	// FileWorkspaceID for file workspace (optional).
	FileWorkspaceID *string `json:"file_workspace_id,omitempty"`

	// LanguagePreference for response language (optional).
	LanguagePreference *string `json:"language_preference,omitempty"`

	// EnableSearchClassifier enables search classification (optional).
	EnableSearchClassifier *bool `json:"enable_search_classifier,omitempty"`

	// BestOf for best-of-n sampling (optional).
	BestOf *int `json:"best_of,omitempty"`

	// DiverseFirstToken for diverse sampling (optional).
	DiverseFirstToken *bool `json:"diverse_first_token,omitempty"`

	// CumLogprobs for cumulative log probabilities (optional).
	CumLogprobs *bool `json:"cum_logprobs,omitempty"`

	// HasImageURL indicates if message has image URL (optional).
	HasImageURL *bool `json:"has_image_url,omitempty"`

	// StreamMode controls streaming mode (optional).
	StreamMode *StreamMode `json:"stream_mode,omitempty"`

	// ForceNewAgent forces creating a new agent for this request (optional).
	ForceNewAgent *bool `json:"_force_new_agent,omitempty"`

	// UserOriginalQuery stores the user's original query (optional).
	UserOriginalQuery *string `json:"user_original_query,omitempty"`

	// Internal/debug parameters (optional, not recommended for general use)
	DebugProSearch    *bool `json:"_debug_pro_search,omitempty"`
	Inputs            []int `json:"_inputs,omitempty"`
	IsBrowserAgent    *bool `json:"_is_browser_agent,omitempty"`
	PromptTokenLength *int  `json:"_prompt_token_length,omitempty"`
}

// Stop can be a single string or array of strings.
type Stop interface {
	isStop()
}

// StopString is a single stop sequence.
type StopString string

func (StopString) isStop() {}

// StopArray is multiple stop sequences.
type StopArray []string

func (StopArray) isStop() {}

// SearchRecencyFilter specifies how recent search results should be.
type SearchRecencyFilter string

// Search recency filter constants
const (
	// SearchRecencyHour filters results from the last hour
	SearchRecencyHour SearchRecencyFilter = "hour"
	// SearchRecencyDay filters results from the last day
	SearchRecencyDay SearchRecencyFilter = "day"
	// SearchRecencyWeek filters results from the last week
	SearchRecencyWeek SearchRecencyFilter = "week"
	// SearchRecencyMonth filters results from the last month
	SearchRecencyMonth SearchRecencyFilter = "month"
	// SearchRecencyYear filters results from the last year
	SearchRecencyYear SearchRecencyFilter = "year"
)

// SearchMode specifies the type of search to perform.
type SearchMode string

// Search mode constants
const (
	// SearchModeWeb performs a web search
	SearchModeWeb SearchMode = "web"
	// SearchModeAcademic performs an academic search
	SearchModeAcademic SearchMode = "academic"
	// SearchModeSEC performs an SEC filings search
	SearchModeSEC SearchMode = "sec"
)

// ReasoningEffort controls the depth of reasoning.
type ReasoningEffort string

// Reasoning effort constants
const (
	// ReasoningEffortMinimal uses minimal reasoning
	ReasoningEffortMinimal ReasoningEffort = "minimal"
	// ReasoningEffortLow uses low reasoning effort
	ReasoningEffortLow ReasoningEffort = "low"
	// ReasoningEffortMedium uses medium reasoning effort
	ReasoningEffortMedium ReasoningEffort = "medium"
	// ReasoningEffortHigh uses high reasoning effort
	ReasoningEffortHigh ReasoningEffort = "high"
)

// StreamMode controls streaming behavior.
type StreamMode string

// Stream mode constants
const (
	// StreamModeFull returns full streaming output
	StreamModeFull StreamMode = "full"
	// StreamModeConcise returns concise streaming output
	StreamModeConcise StreamMode = "concise"
)

// WebSearchOptions configures web search behavior.
type WebSearchOptions struct {
	UserLocation                  *UserLocation      `json:"user_location,omitempty"`
	ImageResultsEnhancedRelevance *bool              `json:"image_results_enhanced_relevance,omitempty"`
	SearchContextSize             *SearchContextSize `json:"search_context_size,omitempty"`
	SearchType                    *SearchType        `json:"search_type,omitempty"`
}

// UserLocation specifies user's geographic location.
type UserLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// SearchContextSize specifies the size of the search context.
type SearchContextSize string

// Search context size constants
const (
	SearchContextSizeLow    SearchContextSize = "low"
	SearchContextSizeMedium SearchContextSize = "medium"
	SearchContextSizeHigh   SearchContextSize = "high"
)

// SearchType specifies the type of search to perform.
type SearchType string

// Search type constants
const (
	SearchTypeFast SearchType = "fast"
	SearchTypePro  SearchType = "pro"
	SearchTypeAuto SearchType = "auto"
)

// ResponseFormat specifies the format of the response.
type ResponseFormat struct {
	Type       ResponseFormatType `json:"type"`
	JSONSchema *JSONSchema        `json:"json_schema,omitempty"`
	Regex      *RegexFormat       `json:"regex,omitempty"`
}

// ResponseFormatType specifies the type of response format.
type ResponseFormatType string

// Response format type constants
const (
	// ResponseFormatTypeText returns plain text responses
	ResponseFormatTypeText ResponseFormatType = "text"
	// ResponseFormatTypeJSONSchema returns JSON following a schema
	ResponseFormatTypeJSONSchema ResponseFormatType = "json_schema"
	// ResponseFormatTypeRegex returns text matching a regex pattern
	ResponseFormatTypeRegex ResponseFormatType = "regex"
)

// JSONSchema defines a JSON schema for structured output.
type JSONSchema struct {
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Schema      map[string]interface{} `json:"schema"`
	Strict      *bool                  `json:"strict,omitempty"`
}

// RegexFormat defines a regex pattern for output.
type RegexFormat struct {
	Pattern string `json:"pattern"`
}
