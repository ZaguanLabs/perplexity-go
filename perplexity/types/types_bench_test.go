package types

import (
	"encoding/json"
	"testing"
)

// Benchmark JSON marshaling for common types

func BenchmarkUsageInfo_Marshal(b *testing.B) {
	usage := UsageInfo{
		CompletionTokens: 100,
		PromptTokens:     50,
		TotalTokens:      150,
		CitationTokens:   Int(10),
		ReasoningTokens:  Int(5),
		Cost: Cost{
			InputTokensCost:     0.001,
			OutputTokensCost:    0.002,
			TotalCost:           0.003,
			CitationTokensCost:  Float64(0.0001),
			ReasoningTokensCost: Float64(0.0002),
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(usage)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUsageInfo_Unmarshal(b *testing.B) {
	data := []byte(`{
		"completion_tokens": 100,
		"prompt_tokens": 50,
		"total_tokens": 150,
		"citation_tokens": 10,
		"reasoning_tokens": 5,
		"cost": {
			"input_tokens_cost": 0.001,
			"output_tokens_cost": 0.002,
			"total_cost": 0.003,
			"citation_tokens_cost": 0.0001,
			"reasoning_tokens_cost": 0.0002
		}
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var usage UsageInfo
		err := json.Unmarshal(data, &usage)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStreamChunk_Marshal(b *testing.B) {
	chunk := StreamChunk{
		ID:      "test-123",
		Created: 1234567890,
		Model:   "sonar",
		Choices: []Choice{
			{
				Index: 0,
				Delta: ChatMessage{
					Role:    RoleAssistant,
					Content: TextContent("Hello, this is a test response from the assistant."),
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
		Citations: []string{"https://example.com/1", "https://example.com/2"},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(chunk)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkStreamChunk_Unmarshal(b *testing.B) {
	data := []byte(`{
		"id": "test-123",
		"created": 1234567890,
		"model": "sonar",
		"choices": [{
			"index": 0,
			"delta": {
				"role": "assistant",
				"content": "Hello, this is a test response from the assistant."
			}
		}],
		"usage": {
			"completion_tokens": 10,
			"prompt_tokens": 5,
			"total_tokens": 15,
			"cost": {
				"total_cost": 0.001
			}
		},
		"citations": ["https://example.com/1", "https://example.com/2"]
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var chunk StreamChunk
		err := json.Unmarshal(data, &chunk)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkChatMessage_Marshal_SimpleText(b *testing.B) {
	msg := ChatMessage{
		Role:    RoleUser,
		Content: TextContent("What is the weather like today?"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkChatMessage_Marshal_StructuredContent(b *testing.B) {
	msg := ChatMessage{
		Role: RoleUser,
		Content: StructuredContent{
			TextChunk{Type: "text", Text: "What's in this image?"},
			ImageChunk{Type: "image_url", ImageURL: ImageURLString("https://example.com/image.jpg")},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkChatMessage_Unmarshal_SimpleText(b *testing.B) {
	data := []byte(`{
		"role": "user",
		"content": "What is the weather like today?"
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var msg ChatMessage
		err := json.Unmarshal(data, &msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkChatMessage_Unmarshal_StructuredContent(b *testing.B) {
	data := []byte(`{
		"role": "user",
		"content": [
			{"type": "text", "text": "What's in this image?"},
			{"type": "image_url", "image_url": "https://example.com/image.jpg"}
		]
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var msg ChatMessage
		err := json.Unmarshal(data, &msg)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkToolCall_Marshal(b *testing.B) {
	toolCall := ToolCall{
		ID:   String("call-123"),
		Type: (*ToolCallType)(String(string(ToolCallTypeFunction))),
		Function: &ToolCallFunction{
			Name:      String("get_weather"),
			Arguments: String(`{"location": "San Francisco", "unit": "celsius"}`),
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(toolCall)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkToolCall_Unmarshal(b *testing.B) {
	data := []byte(`{
		"id": "call-123",
		"type": "function",
		"function": {
			"name": "get_weather",
			"arguments": "{\"location\": \"San Francisco\", \"unit\": \"celsius\"}"
		}
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var toolCall ToolCall
		err := json.Unmarshal(data, &toolCall)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReasoningStep_Marshal(b *testing.B) {
	step := ReasoningStep{
		Thought: "I need to search for information about Go programming",
		Type:    String("web_search"),
		WebSearch: &ReasoningStepWebSearch{
			SearchKeywords: []string{"golang", "tutorial", "best practices"},
			SearchResults: []SearchResult{
				{
					Title:   "Go Tutorial",
					URL:     "https://example.com/go",
					Snippet: String("Learn Go programming"),
				},
				{
					Title:   "Go Best Practices",
					URL:     "https://example.com/best-practices",
					Snippet: String("Best practices for Go development"),
				},
			},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(step)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkReasoningStep_Unmarshal(b *testing.B) {
	data := []byte(`{
		"thought": "I need to search for information about Go programming",
		"type": "web_search",
		"web_search": {
			"search_keywords": ["golang", "tutorial", "best practices"],
			"search_results": [
				{
					"title": "Go Tutorial",
					"url": "https://example.com/go",
					"snippet": "Learn Go programming"
				},
				{
					"title": "Go Best Practices",
					"url": "https://example.com/best-practices",
					"snippet": "Best practices for Go development"
				}
			]
		}
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var step ReasoningStep
		err := json.Unmarshal(data, &step)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSearchResponse_Marshal(b *testing.B) {
	response := SearchResponse{
		ID: "search-123",
		Results: []SearchResultItem{
			{
				Title:       "Result 1",
				URL:         "https://example.com/1",
				Snippet:     "This is result 1",
				Date:        String("2024-01-01"),
				LastUpdated: String("2024-01-15"),
			},
			{
				Title:       "Result 2",
				URL:         "https://example.com/2",
				Snippet:     "This is result 2",
				Date:        String("2024-01-02"),
				LastUpdated: String("2024-01-16"),
			},
			{
				Title:       "Result 3",
				URL:         "https://example.com/3",
				Snippet:     "This is result 3",
				Date:        String("2024-01-03"),
				LastUpdated: String("2024-01-17"),
			},
		},
		ServerTime: String("2024-01-01T00:00:00Z"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(response)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSearchResponse_Unmarshal(b *testing.B) {
	data := []byte(`{
		"id": "search-123",
		"results": [
			{
				"title": "Result 1",
				"url": "https://example.com/1",
				"snippet": "This is result 1",
				"date": "2024-01-01",
				"last_updated": "2024-01-15"
			},
			{
				"title": "Result 2",
				"url": "https://example.com/2",
				"snippet": "This is result 2",
				"date": "2024-01-02",
				"last_updated": "2024-01-16"
			},
			{
				"title": "Result 3",
				"url": "https://example.com/3",
				"snippet": "This is result 3",
				"date": "2024-01-03",
				"last_updated": "2024-01-17"
			}
		],
		"server_time": "2024-01-01T00:00:00Z"
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var response SearchResponse
		err := json.Unmarshal(data, &response)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark complex nested structures

func BenchmarkComplexResponse_Marshal(b *testing.B) {
	chunk := StreamChunk{
		ID:      "test-complex-123",
		Created: 1234567890,
		Model:   "sonar-pro",
		Choices: []Choice{
			{
				Index: 0,
				Delta: ChatMessage{
					Role:    RoleAssistant,
					Content: TextContent("Based on my search, here's what I found..."),
					ToolCalls: []ToolCall{
						{
							ID:   String("call-1"),
							Type: (*ToolCallType)(String(string(ToolCallTypeFunction))),
							Function: &ToolCallFunction{
								Name:      String("search_web"),
								Arguments: String(`{"query": "golang performance"}`),
							},
						},
					},
					ReasoningSteps: []ReasoningStep{
						{
							Thought: "Need to search for information",
							Type:    String("web_search"),
							WebSearch: &ReasoningStepWebSearch{
								SearchKeywords: []string{"golang", "performance"},
								SearchResults: []SearchResult{
									{Title: "Go Performance", URL: "https://example.com"},
								},
							},
						},
					},
				},
			},
		},
		Usage: &UsageInfo{
			CompletionTokens: 150,
			PromptTokens:     50,
			TotalTokens:      200,
			CitationTokens:   Int(10),
			ReasoningTokens:  Int(20),
			Cost: Cost{
				InputTokensCost:     0.001,
				OutputTokensCost:    0.003,
				TotalCost:           0.004,
				CitationTokensCost:  Float64(0.0001),
				ReasoningTokensCost: Float64(0.0002),
			},
		},
		Citations: []string{"https://example.com/1", "https://example.com/2"},
		SearchResults: []SearchResult{
			{Title: "Result 1", URL: "https://example.com/1"},
			{Title: "Result 2", URL: "https://example.com/2"},
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(chunk)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkComplexResponse_Unmarshal(b *testing.B) {
	data := []byte(`{
		"id": "test-complex-123",
		"created": 1234567890,
		"model": "sonar-pro",
		"choices": [{
			"index": 0,
			"delta": {
				"role": "assistant",
				"content": "Based on my search, here's what I found...",
				"tool_calls": [{
					"id": "call-1",
					"type": "function",
					"function": {
						"name": "search_web",
						"arguments": "{\"query\": \"golang performance\"}"
					}
				}],
				"reasoning_steps": [{
					"thought": "Need to search for information",
					"type": "web_search",
					"web_search": {
						"search_keywords": ["golang", "performance"],
						"search_results": [{
							"title": "Go Performance",
							"url": "https://example.com"
						}]
					}
				}]
			}
		}],
		"usage": {
			"completion_tokens": 150,
			"prompt_tokens": 50,
			"total_tokens": 200,
			"citation_tokens": 10,
			"reasoning_tokens": 20,
			"cost": {
				"input_tokens_cost": 0.001,
				"output_tokens_cost": 0.003,
				"total_cost": 0.004,
				"citation_tokens_cost": 0.0001,
				"reasoning_tokens_cost": 0.0002
			}
		},
		"citations": ["https://example.com/1", "https://example.com/2"],
		"search_results": [
			{"title": "Result 1", "url": "https://example.com/1"},
			{"title": "Result 2", "url": "https://example.com/2"}
		]
	}`)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var chunk StreamChunk
		err := json.Unmarshal(data, &chunk)
		if err != nil {
			b.Fatal(err)
		}
	}
}
