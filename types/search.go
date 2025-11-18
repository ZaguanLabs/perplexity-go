package types

// SearchResult represents a single search result.
type SearchResult struct {
	// Title is the title of the search result.
	Title string `json:"title"`

	// URL is the URL of the search result.
	URL string `json:"url"`

	// Date is the publication date (optional).
	Date *string `json:"date,omitempty"`

	// LastUpdated is when the content was last updated (optional).
	LastUpdated *string `json:"last_updated,omitempty"`

	// Snippet is a text snippet from the result (optional).
	Snippet *string `json:"snippet,omitempty"`

	// Source indicates where the result came from (optional).
	Source *SearchResultSource `json:"source,omitempty"`
}

// SearchResultSource indicates the source of a search result.
type SearchResultSource string

const (
	// SearchResultSourceWeb indicates the result came from web search.
	SearchResultSourceWeb SearchResultSource = "web"

	// SearchResultSourceAttachment indicates the result came from an attachment.
	SearchResultSourceAttachment SearchResultSource = "attachment"
)

// SearchResponse is the response from a search request.
type SearchResponse struct {
	// ID is the unique identifier for the search.
	ID string `json:"id"`

	// Results contains the search results.
	Results []SearchResultItem `json:"results"`

	// ServerTime is the server timestamp (optional).
	ServerTime *string `json:"server_time,omitempty"`
}

// SearchResultItem represents a single item in search results.
type SearchResultItem struct {
	// Snippet is a text snippet from the result.
	Snippet string `json:"snippet"`

	// Title is the title of the result.
	Title string `json:"title"`

	// URL is the URL of the result.
	URL string `json:"url"`

	// Date is the publication date (optional).
	Date *string `json:"date,omitempty"`

	// LastUpdated is when the content was last updated (optional).
	LastUpdated *string `json:"last_updated,omitempty"`
}
