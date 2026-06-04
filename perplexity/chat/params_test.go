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
			City:      stringPtr("San Francisco"),
			Country:   stringPtr("US"),
			Latitude:  float64Ptr(37.7749),
			Longitude: float64Ptr(-122.4194),
			Region:    stringPtr("CA"),
		},
		ImageResultsEnhancedRelevance: &enhanced,
		SearchContextSize:             &contextSize,
		SearchType:                    &searchType,
	}

	data, err := json.Marshal(opts)
	if err != nil {
		t.Fatalf("Failed to marshal WebSearchOptions: %v", err)
	}

	expected := `{"user_location":{"city":"San Francisco","country":"US","latitude":37.7749,"longitude":-122.4194,"region":"CA"},"image_results_enhanced_relevance":true,"search_context_size":"high","search_type":"pro"}`
	if string(data) != expected {
		t.Errorf("JSON marshaling mismatch.\nExpected: %s\nGot:      %s", expected, string(data))
	}
}

func stringPtr(v string) *string {
	return &v
}

func float64Ptr(v float64) *float64 {
	return &v
}
