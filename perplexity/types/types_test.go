package types

import (
	"encoding/json"
	"testing"
)

func TestUsageInfo_JSON(t *testing.T) {
	usage := UsageInfo{
		CompletionTokens: 100,
		PromptTokens:     50,
		TotalTokens:      150,
		Cost: Cost{
			InputTokensCost:  0.001,
			OutputTokensCost: 0.002,
			TotalCost:        0.003,
		},
	}

	// Marshal
	data, err := json.Marshal(usage)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result UsageInfo
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.CompletionTokens != usage.CompletionTokens {
		t.Errorf("CompletionTokens = %d, want %d", result.CompletionTokens, usage.CompletionTokens)
	}
	if result.TotalTokens != usage.TotalTokens {
		t.Errorf("TotalTokens = %d, want %d", result.TotalTokens, usage.TotalTokens)
	}
	if result.Cost.TotalCost != usage.Cost.TotalCost {
		t.Errorf("Cost.TotalCost = %f, want %f", result.Cost.TotalCost, usage.Cost.TotalCost)
	}
}

func TestSearchResult_JSON(t *testing.T) {
	result := SearchResult{
		Title:   "Test Result",
		URL:     "https://example.com",
		Snippet: String("This is a snippet"),
		Source:  (*SearchResultSource)(String(string(SearchResultSourceWeb))),
	}

	// Marshal
	data, err := json.Marshal(result)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var unmarshaled SearchResult
	if err := json.Unmarshal(data, &unmarshaled); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if unmarshaled.Title != result.Title {
		t.Errorf("Title = %q, want %q", unmarshaled.Title, result.Title)
	}
	if unmarshaled.URL != result.URL {
		t.Errorf("URL = %q, want %q", unmarshaled.URL, result.URL)
	}
	if unmarshaled.Snippet == nil || *unmarshaled.Snippet != *result.Snippet {
		t.Errorf("Snippet mismatch")
	}
}

func TestChoice_JSON(t *testing.T) {
	choice := Choice{
		Index: 0,
		Message: ChatMessage{
			Role:    RoleAssistant,
			Content: TextContent("Hello!"),
		},
		FinishReason: (*FinishReason)(String(string(FinishReasonStop))),
	}

	// Marshal
	data, err := json.Marshal(choice)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result Choice
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.Index != choice.Index {
		t.Errorf("Index = %d, want %d", result.Index, choice.Index)
	}
	if result.Message.Role != choice.Message.Role {
		t.Errorf("Message.Role = %v, want %v", result.Message.Role, choice.Message.Role)
	}
}

func TestStreamChunk_JSON(t *testing.T) {
	chunk := StreamChunk{
		ID:      "test-123",
		Created: 1234567890,
		Model:   "sonar",
		Choices: []Choice{
			{
				Index: 0,
				Delta: ChatMessage{
					Role:    RoleAssistant,
					Content: TextContent("Hello"),
				},
			},
		},
		Usage: &UsageInfo{
			CompletionTokens: 10,
			PromptTokens:     5,
			TotalTokens:      15,
			Cost: Cost{
				TotalCost: 0.001,
			},
		},
	}

	// Marshal
	data, err := json.Marshal(chunk)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result StreamChunk
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.ID != chunk.ID {
		t.Errorf("ID = %q, want %q", result.ID, chunk.ID)
	}
	if result.Model != chunk.Model {
		t.Errorf("Model = %q, want %q", result.Model, chunk.Model)
	}
	if len(result.Choices) != len(chunk.Choices) {
		t.Errorf("Choices length = %d, want %d", len(result.Choices), len(chunk.Choices))
	}
	if result.Usage == nil {
		t.Fatal("Usage is nil")
	}
	if result.Usage.TotalTokens != chunk.Usage.TotalTokens {
		t.Errorf("Usage.TotalTokens = %d, want %d", result.Usage.TotalTokens, chunk.Usage.TotalTokens)
	}
}

func TestToolCall_JSON(t *testing.T) {
	toolCall := ToolCall{
		ID:   String("call-123"),
		Type: (*ToolCallType)(String(string(ToolCallTypeFunction))),
		Function: &ToolCallFunction{
			Name:      String("get_weather"),
			Arguments: String(`{"location": "San Francisco"}`),
		},
	}

	// Marshal
	data, err := json.Marshal(toolCall)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result ToolCall
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.ID == nil || *result.ID != *toolCall.ID {
		t.Errorf("ID mismatch")
	}
	if result.Function == nil {
		t.Fatal("Function is nil")
	}
	if result.Function.Name == nil || *result.Function.Name != *toolCall.Function.Name {
		t.Errorf("Function.Name mismatch")
	}
}

func TestReasoningStep_JSON(t *testing.T) {
	step := ReasoningStep{
		Thought: "I need to search for information",
		WebSearch: &ReasoningStepWebSearch{
			SearchKeywords: []string{"golang", "tutorial"},
			SearchResults: []SearchResult{
				{
					Title: "Go Tutorial",
					URL:   "https://example.com/go",
				},
			},
		},
	}

	// Marshal
	data, err := json.Marshal(step)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result ReasoningStep
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.Thought != step.Thought {
		t.Errorf("Thought = %q, want %q", result.Thought, step.Thought)
	}
	if result.WebSearch == nil {
		t.Fatal("WebSearch is nil")
	}
	if len(result.WebSearch.SearchKeywords) != 2 {
		t.Errorf("SearchKeywords length = %d, want 2", len(result.WebSearch.SearchKeywords))
	}
}

func TestTool_JSON(t *testing.T) {
	tool := Tool{
		Type: ToolTypeFunction,
		Function: ToolFunction{
			Name:        "get_weather",
			Description: String("Get the current weather"),
			Parameters: map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"location": map[string]interface{}{
						"type":        "string",
						"description": "The city and state",
					},
				},
			},
		},
	}

	// Marshal
	data, err := json.Marshal(tool)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result Tool
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.Type != tool.Type {
		t.Errorf("Type = %v, want %v", result.Type, tool.Type)
	}
	if result.Function.Name != tool.Function.Name {
		t.Errorf("Function.Name = %q, want %q", result.Function.Name, tool.Function.Name)
	}
}

func TestSearchResponse_JSON(t *testing.T) {
	response := SearchResponse{
		ID: "search-123",
		Results: []SearchResultItem{
			{
				Title:   "Result 1",
				URL:     "https://example.com/1",
				Snippet: "This is result 1",
			},
			{
				Title:   "Result 2",
				URL:     "https://example.com/2",
				Snippet: "This is result 2",
			},
		},
		ServerTime: String("2024-01-01T00:00:00Z"),
	}

	// Marshal
	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result SearchResponse
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if result.ID != response.ID {
		t.Errorf("ID = %q, want %q", result.ID, response.ID)
	}
	if len(result.Results) != len(response.Results) {
		t.Errorf("Results length = %d, want %d", len(result.Results), len(response.Results))
	}
	if result.ServerTime == nil || *result.ServerTime != *response.ServerTime {
		t.Errorf("ServerTime mismatch")
	}
}
