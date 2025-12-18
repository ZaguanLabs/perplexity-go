// Package search provides access to Perplexity's web search API.
//
// This package implements web search functionality with support for
// filtering, multiple queries, and specialized search modes (web, academic, SEC filings).
//
// # Basic Usage
//
// Perform a simple web search:
//
//	result, err := client.Search.Create(ctx, &search.SearchParams{
//		Query:      "latest AI developments",
//		MaxResults: types.Int(5),
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, item := range result.Results {
//		fmt.Printf("%s: %s\n", item.Title, item.URL)
//	}
//
// # Search Modes
//
// Use specialized search modes:
//
//	// Academic search
//	result, err := client.Search.Create(ctx, &search.SearchParams{
//		Query:      "machine learning research",
//		SearchMode: types.SearchModePtr(types.SearchModeAcademic),
//		MaxResults: types.Int(10),
//	})
//
//	// SEC filings search
//	result, err := client.Search.Create(ctx, &search.SearchParams{
//		Query:      "Tesla quarterly report",
//		SearchMode: types.SearchModePtr(types.SearchModeSec),
//	})
//
// # Filtering
//
// Apply filters to refine search results:
//
//	result, err := client.Search.Create(ctx, &search.SearchParams{
//		Query:                "golang tutorial",
//		SearchDomainFilter:   []string{"golang.org", "go.dev"},
//		SearchRecencyFilter:  types.SearchRecencyFilterPtr(types.SearchRecencyWeek),
//		SearchLanguageFilter: []string{"en"},
//		MaxResults:           types.Int(10),
//	})
//
// # Date Filtering
//
// Filter by date range:
//
//	result, err := client.Search.Create(ctx, &search.SearchParams{
//		Query:             "AI news",
//		SearchAfterDate:   types.String("2024-01-01"),
//		SearchBeforeDate:  types.String("2024-12-31"),
//	})
//
// # Multiple Queries
//
// Search with multiple queries simultaneously:
//
//	result, err := client.Search.Create(ctx, &search.SearchParams{
//		Query: []string{
//			"golang performance",
//			"golang best practices",
//			"golang concurrency",
//		},
//		MaxResults: types.Int(5),
//	})
package search
