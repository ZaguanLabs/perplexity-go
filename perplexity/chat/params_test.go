package chat

import (
	"encoding/json"
	"testing"
)

func TestWebSearchOptionsJSON(t *testing.T) {
	enhanced := true
	contextSize := SearchContextSizeHigh
	searchType := SearchTypePro

	opts := WebSearchOptions{
		UserLocation: &UserLocation{
			Latitude:  37.7749,
			Longitude: -122.4194,
		},
		ImageResultsEnhancedRelevance: &enhanced,
		SearchContextSize:             &contextSize,
		SearchType:                    &searchType,
	}

	data, err := json.Marshal(opts)
	if err != nil {
		t.Fatalf("Failed to marshal WebSearchOptions: %v", err)
	}

	expected := `{"user_location":{"latitude":37.7749,"longitude":-122.4194},"image_results_enhanced_relevance":true,"search_context_size":"high","search_type":"pro"}`
	if string(data) != expected {
		t.Errorf("JSON marshaling mismatch.\nExpected: %s\nGot:      %s", expected, string(data))
	}
}
