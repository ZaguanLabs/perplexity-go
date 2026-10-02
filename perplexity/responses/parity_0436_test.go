package responses

import (
	"encoding/json"
	"testing"

	"github.com/ZaguanLabs/perplexity-go/perplexity/search"
)

func TestPython0436WireModes(t *testing.T) {
	mode := WebSearchTypeFast
	effort := ReasoningEffortNone
	payload := struct {
		Tools     []Tool    `json:"tools"`
		Reasoning Reasoning `json:"reasoning"`
	}{Tools: []Tool{NewToolFromWebSearch(WebSearchTool{Type: ToolTypeWebSearch, SearchType: &mode})}, Reasoning: Reasoning{Effort: &effort}}
	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}
	var wire struct {
		Tools     []map[string]any `json:"tools"`
		Reasoning map[string]any   `json:"reasoning"`
	}
	if err := json.Unmarshal(data, &wire); err != nil {
		t.Fatal(err)
	}
	if wire.Tools[0]["search_type"] != "fast" || wire.Reasoning["effort"] != "none" {
		t.Fatalf("wire: %s", data)
	}
	var decoded Tool
	if err := json.Unmarshal([]byte(`{"type":"web_search","search_type":"fast"}`), &decoded); err != nil {
		t.Fatal(err)
	}
	tool, ok := decoded.AsWebSearch()
	if !ok || tool.SearchType == nil || *tool.SearchType != WebSearchTypeFast {
		t.Fatalf("decoded: %#v", decoded)
	}
	data, err = json.Marshal(WebSearchTool{Type: ToolTypeWebSearch})
	if err != nil {
		t.Fatal(err)
	}
	var omitted map[string]any
	if err := json.Unmarshal(data, &omitted); err != nil {
		t.Fatal(err)
	}
	if _, ok := omitted["search_type"]; ok {
		t.Fatalf("unset search_type emitted: %s", data)
	}
	fast := search.SearchTypeFast
	data, err = json.Marshal(search.SearchParams{SearchType: &fast})
	if err != nil {
		t.Fatal(err)
	}
	var request map[string]any
	if err := json.Unmarshal(data, &request); err != nil {
		t.Fatal(err)
	}
	if request["search_type"] != "fast" {
		t.Fatalf("search: %s", data)
	}
}
