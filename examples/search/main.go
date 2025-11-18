package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ZaguanLabs/perplexity-go/perplexity"
	"github.com/ZaguanLabs/perplexity-go/perplexity/chat"
	"github.com/ZaguanLabs/perplexity-go/perplexity/search"
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

func main() {
	// Create client (reads API key from PERPLEXITY_API_KEY environment variable)
	client, err := perplexity.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Perplexity Go SDK v%s\n\n", client.Version())

	// Example 1: Basic web search
	fmt.Println("=== Example 1: Basic Web Search ===")
	fmt.Println()
	basicSearch(client)

	fmt.Println()
	fmt.Println(string(make([]byte, 50)))
	fmt.Println()

	// Example 2: Advanced search with filters
	fmt.Println("=== Example 2: Advanced Search with Filters ===")
	fmt.Println()
	advancedSearch(client)

	fmt.Println()
	fmt.Println(string(make([]byte, 50)))
	fmt.Println()

	// Example 3: Multiple queries
	fmt.Println("=== Example 3: Multiple Queries ===")
	fmt.Println()
	multipleQueries(client)

	// Check if API key is set
	if os.Getenv("PERPLEXITY_API_KEY") == "" {
		fmt.Println("\n⚠️  Set PERPLEXITY_API_KEY environment variable to run this example")
	}
}

func basicSearch(client *perplexity.Client) {
	params := &search.SearchParams{
		Query:      "latest developments in artificial intelligence",
		MaxResults: types.Int(5),
	}

	fmt.Println("Searching for:", params.Query)

	result, err := client.Search.Create(context.Background(), params)
	if err != nil {
		log.Printf("Search failed: %v", err)
		return
	}

	fmt.Printf("\nSearch ID: %s\n", result.ID)
	fmt.Printf("Found %d results:\n\n", len(result.Results))

	for i, item := range result.Results {
		fmt.Printf("%d. %s\n", i+1, item.Title)
		fmt.Printf("   URL: %s\n", item.URL)
		if item.Date != nil {
			fmt.Printf("   Date: %s\n", *item.Date)
		}
		fmt.Printf("   %s\n", item.Snippet)
		fmt.Println()
	}

	if result.ServerTime != nil {
		fmt.Printf("Server time: %s\n", *result.ServerTime)
	}
}

func advancedSearch(client *perplexity.Client) {
	recencyFilter := chat.SearchRecencyWeek
	searchMode := chat.SearchModeWeb

	params := &search.SearchParams{
		Query:                "Go programming language",
		MaxResults:           types.Int(3),
		SearchRecencyFilter:  &recencyFilter,
		SearchMode:           &searchMode,
		SearchDomainFilter:   []string{"golang.org", "go.dev"},
		SearchLanguageFilter: []string{"en"},
		Country:              types.String("US"),
		DisplayServerTime:    types.Bool(true),
	}

	fmt.Println("Searching with filters:")
	fmt.Printf("  Query: %s\n", params.Query)
	fmt.Printf("  Domains: %v\n", params.SearchDomainFilter)
	fmt.Printf("  Recency: %s\n", *params.SearchRecencyFilter)
	fmt.Println()

	result, err := client.Search.Create(context.Background(), params)
	if err != nil {
		log.Printf("Search failed: %v", err)
		return
	}

	fmt.Printf("Found %d results:\n\n", len(result.Results))

	for i, item := range result.Results {
		fmt.Printf("%d. %s\n", i+1, item.Title)
		fmt.Printf("   %s\n", item.URL)
		if item.LastUpdated != nil {
			fmt.Printf("   Last updated: %s\n", *item.LastUpdated)
		}
		fmt.Println()
	}
}

func multipleQueries(client *perplexity.Client) {
	params := &search.SearchParams{
		MaxResults: types.Int(2),
	}
	params.QueryStrings([]string{
		"Go concurrency patterns",
		"Go best practices",
	})

	fmt.Println("Searching for multiple queries...")

	result, err := client.Search.Create(context.Background(), params)
	if err != nil {
		log.Printf("Search failed: %v", err)
		return
	}

	fmt.Printf("\nFound %d total results:\n\n", len(result.Results))

	for i, item := range result.Results {
		fmt.Printf("%d. %s\n", i+1, item.Title)
		fmt.Printf("   %s\n", item.URL)
		fmt.Println()
	}
}

// Example: Academic search
func academicSearch() {
	client, _ := perplexity.NewClient("")

	searchMode := chat.SearchModeAcademic
	params := &search.SearchParams{
		Query:      "machine learning algorithms",
		MaxResults: types.Int(5),
		SearchMode: &searchMode,
	}

	result, err := client.Search.Create(context.Background(), params)
	if err != nil {
		log.Printf("Search failed: %v", err)
		return
	}

	fmt.Printf("Academic search found %d results\n", len(result.Results))
	for _, item := range result.Results {
		fmt.Printf("- %s (%s)\n", item.Title, item.URL)
	}
}

// Example: SEC filings search
func secSearch() {
	client, _ := perplexity.NewClient("")

	searchMode := chat.SearchModeSEC
	params := &search.SearchParams{
		Query:      "Tesla quarterly earnings",
		MaxResults: types.Int(5),
		SearchMode: &searchMode,
	}

	result, err := client.Search.Create(context.Background(), params)
	if err != nil {
		log.Printf("Search failed: %v", err)
		return
	}

	fmt.Printf("SEC search found %d results\n", len(result.Results))
	for _, item := range result.Results {
		fmt.Printf("- %s\n", item.Title)
		if item.Date != nil {
			fmt.Printf("  Filed: %s\n", *item.Date)
		}
	}
}
