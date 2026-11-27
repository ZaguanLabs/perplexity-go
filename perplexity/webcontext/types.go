// Code generated from the Perplexity Python SDK 0.50.0 schema. DO NOT EDIT.

package webcontext

import "time"

type ContextCost struct {
	Currency  ContextCurrency `json:"currency"`
	TotalCost float64         `json:"total_cost"`
}
type ContextCurrency = string
type ContextDocument struct {
	Highlights *string          `json:"highlights,omitempty"`
	HTML       *string          `json:"html,omitempty"`
	Links      *[]string        `json:"links,omitempty"`
	Markdown   *string          `json:"markdown,omitempty"`
	Metadata   *ContextMetadata `json:"metadata,omitempty"`
	Status     ContextStatus    `json:"status"`
	URL        string           `json:"url"`
}
type ContextHighlights struct {
	MaxTokensPerURL *int   `json:"max_tokens_per_url,omitempty"`
	Query           string `json:"query"`
}
type ContextMetadata struct {
	Author        *string `json:"author,omitempty"`
	ContentType   *string `json:"content_type,omitempty"`
	Description   *string `json:"description,omitempty"`
	FetchedAt     *string `json:"fetched_at,omitempty"`
	Language      *string `json:"language,omitempty"`
	LastModified  *string `json:"last_modified,omitempty"`
	Origin        *string `json:"origin,omitempty"`
	PublishedDate *string `json:"published_date,omitempty"`
	SourceURL     string  `json:"source_url"`
	Title         *string `json:"title,omitempty"`
}
type ContextRequest struct {
	CacheOnly       *bool              `json:"cache_only,omitempty"`
	Formats         *[]string          `json:"formats,omitempty"`
	Highlights      *ContextHighlights `json:"highlights,omitempty"`
	IfModifiedSince *string            `json:"if_modified_since,omitempty"`
	MaxAgeSeconds   *int               `json:"max_age_seconds,omitempty"`
	SnapshotAsOf    *time.Time         `json:"snapshot_as_of,omitempty"`
	TimeoutMS       *int               `json:"timeout_ms,omitempty"`
	URLs            []string           `json:"urls"`
}
type ContextResponse struct {
	RequestID string            `json:"request_id"`
	Results   []ContextDocument `json:"results"`
	Usage     ContextUsage      `json:"usage"`
}
type ContextStatus struct {
	Message    *string `json:"message,omitempty"`
	StatusCode *int    `json:"status_code,omitempty"`
	Type       string  `json:"type"`
}
type ContextUsage struct {
	Cost        ContextCost `json:"cost"`
	FetchedURLs int         `json:"fetched_urls"`
}
