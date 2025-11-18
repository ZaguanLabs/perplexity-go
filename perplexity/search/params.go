package search

import (
	"github.com/perplexityai/perplexity-go/perplexity/chat"
)

// SearchParams contains parameters for creating a search request.
type SearchParams struct {
	// Query is the search query (required).
	// Can be a single string or multiple queries.
	Query interface{} `json:"query"` // string or []string

	// Country for search localization (optional).
	Country *string `json:"country,omitempty"`

	// DisplayServerTime includes server time in response (optional).
	DisplayServerTime *bool `json:"display_server_time,omitempty"`

	// MaxResults is the maximum number of results to return (optional).
	MaxResults *int `json:"max_results,omitempty"`

	// MaxTokens is the maximum tokens per result (optional).
	MaxTokens *int `json:"max_tokens,omitempty"`

	// MaxTokensPerPage is the maximum tokens per page (optional).
	MaxTokensPerPage *int `json:"max_tokens_per_page,omitempty"`

	// SearchAfterDateFilter filters results after this date (optional).
	SearchAfterDateFilter *string `json:"search_after_date_filter,omitempty"`

	// SearchBeforeDateFilter filters results before this date (optional).
	SearchBeforeDateFilter *string `json:"search_before_date_filter,omitempty"`

	// SearchDomainFilter limits search to specific domains (optional).
	SearchDomainFilter []string `json:"search_domain_filter,omitempty"`

	// SearchLanguageFilter filters by language (optional).
	SearchLanguageFilter []string `json:"search_language_filter,omitempty"`

	// SearchMode specifies the search mode (optional).
	SearchMode *chat.SearchMode `json:"search_mode,omitempty"`

	// SearchRecencyFilter filters by recency (optional).
	SearchRecencyFilter *chat.SearchRecencyFilter `json:"search_recency_filter,omitempty"`
}

// QueryString sets a single query string.
func (p *SearchParams) QueryString(query string) {
	p.Query = query
}

// QueryStrings sets multiple query strings.
func (p *SearchParams) QueryStrings(queries []string) {
	p.Query = queries
}
