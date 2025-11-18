package search

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

func TestService_Create(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/search" {
			t.Errorf("Expected path /search, got %s", r.URL.Path)
		}

		// Verify headers
		if r.Header.Get("Authorization") == "" {
			t.Error("Missing Authorization header")
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Error("Missing or incorrect Content-Type header")
		}

		// Parse request body
		var params SearchParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}

		// Verify query field
		if params.Query == nil {
			t.Error("Query is nil")
		}

		// Return mock response
		response := types.SearchResponse{
			ID: "search-123",
			Results: []types.SearchResultItem{
				{
					Title:   "Go Programming Language",
					URL:     "https://golang.org",
					Snippet: "The Go programming language is an open source project...",
					Date:    types.String("2024-01-01"),
				},
				{
					Title:   "Effective Go",
					URL:     "https://golang.org/doc/effective_go",
					Snippet: "A document that gives tips for writing clear, idiomatic Go code.",
				},
			},
			ServerTime: types.String("2024-01-15T10:00:00Z"),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	// Create service with mock server
	httpClient := internalhttp.NewClient(
		server.Client(),
		server.URL,
		"test-api-key",
		2,
		nil,
		"test-agent",
	)
	service := NewService(httpClient)

	// Test successful request with string query
	t.Run("successful request with string query", func(t *testing.T) {
		params := &SearchParams{
			Query:      "golang programming",
			MaxResults: types.Int(10),
		}

		result, err := service.Create(context.Background(), params)
		if err != nil {
			t.Fatalf("Create failed: %v", err)
		}

		if result == nil {
			t.Fatal("Result is nil")
		}

		if result.ID != "search-123" {
			t.Errorf("ID = %s, want search-123", result.ID)
		}

		if len(result.Results) != 2 {
			t.Fatalf("Results length = %d, want 2", len(result.Results))
		}

		if result.Results[0].Title != "Go Programming Language" {
			t.Errorf("First result title = %s", result.Results[0].Title)
		}

		if result.ServerTime == nil {
			t.Error("ServerTime is nil")
		}
	})

	// Test with array query
	t.Run("successful request with array query", func(t *testing.T) {
		params := &SearchParams{
			Query: []string{"golang", "programming"},
		}

		result, err := service.Create(context.Background(), params)
		if err != nil {
			t.Fatalf("Create failed: %v", err)
		}

		if result == nil {
			t.Fatal("Result is nil")
		}

		if len(result.Results) != 2 {
			t.Errorf("Results length = %d, want 2", len(result.Results))
		}
	})

	// Test validation errors
	t.Run("nil params", func(t *testing.T) {
		_, err := service.Create(context.Background(), nil)
		if err == nil {
			t.Error("Expected error for nil params")
		}
	})

	t.Run("nil query", func(t *testing.T) {
		params := &SearchParams{}
		_, err := service.Create(context.Background(), params)
		if err == nil {
			t.Error("Expected error for nil query")
		}
	})

	t.Run("empty string query", func(t *testing.T) {
		params := &SearchParams{
			Query: "",
		}
		_, err := service.Create(context.Background(), params)
		if err == nil {
			t.Error("Expected error for empty string query")
		}
	})

	t.Run("empty array query", func(t *testing.T) {
		params := &SearchParams{
			Query: []string{},
		}
		_, err := service.Create(context.Background(), params)
		if err == nil {
			t.Error("Expected error for empty array query")
		}
	})

	t.Run("array with empty string", func(t *testing.T) {
		params := &SearchParams{
			Query: []string{"valid", ""},
		}
		_, err := service.Create(context.Background(), params)
		if err == nil {
			t.Error("Expected error for array with empty string")
		}
	})

	t.Run("invalid query type", func(t *testing.T) {
		params := &SearchParams{
			Query: 123, // Invalid type
		}
		_, err := service.Create(context.Background(), params)
		if err == nil {
			t.Error("Expected error for invalid query type")
		}
	})
}

func TestSearchParams_QueryHelpers(t *testing.T) {
	t.Run("QueryString", func(t *testing.T) {
		params := &SearchParams{}
		params.QueryString("test query")

		if params.Query != "test query" {
			t.Errorf("Query = %v, want 'test query'", params.Query)
		}
	})

	t.Run("QueryStrings", func(t *testing.T) {
		params := &SearchParams{}
		queries := []string{"query1", "query2"}
		params.QueryStrings(queries)

		queryArray, ok := params.Query.([]string)
		if !ok {
			t.Fatalf("Query is not []string, got %T", params.Query)
		}

		if len(queryArray) != 2 {
			t.Errorf("Query length = %d, want 2", len(queryArray))
		}

		if queryArray[0] != "query1" {
			t.Errorf("Query[0] = %s, want query1", queryArray[0])
		}
	})
}

func TestSearchParams_JSON(t *testing.T) {
	params := &SearchParams{
		Query:                "golang",
		MaxResults:           types.Int(10),
		Country:              types.String("US"),
		DisplayServerTime:    types.Bool(true),
		SearchDomainFilter:   []string{"golang.org", "go.dev"},
		SearchLanguageFilter: []string{"en"},
	}

	// Marshal
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result SearchParams
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Verify
	if result.Query != params.Query {
		t.Errorf("Query = %v, want %v", result.Query, params.Query)
	}
	if result.MaxResults == nil || *result.MaxResults != 10 {
		t.Error("MaxResults mismatch")
	}
	if len(result.SearchDomainFilter) != 2 {
		t.Errorf("SearchDomainFilter length = %d, want 2", len(result.SearchDomainFilter))
	}
}
