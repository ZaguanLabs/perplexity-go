package chat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	internalhttp "github.com/perplexityai/perplexity-go/perplexity/internal/http"
	"github.com/perplexityai/perplexity-go/perplexity/types"
)

func TestService_Create(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/chat/completions" {
			t.Errorf("Expected path /chat/completions, got %s", r.URL.Path)
		}

		// Verify headers
		if r.Header.Get("Authorization") == "" {
			t.Error("Missing Authorization header")
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Error("Missing or incorrect Content-Type header")
		}

		// Parse request body
		var params CompletionParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}

		// Verify required fields
		if len(params.Messages) == 0 {
			t.Error("Messages are empty")
		}
		if params.Model == "" {
			t.Error("Model is empty")
		}

		// Return mock response
		response := types.StreamChunk{
			ID:      "test-123",
			Model:   "sonar",
			Created: 1234567890,
			Choices: []types.Choice{
				{
					Index: 0,
					Message: types.ChatMessage{
						Role:    types.RoleAssistant,
						Content: types.TextContent("Hello! How can I help you today?"),
					},
					FinishReason: (*types.FinishReason)(types.String(string(types.FinishReasonStop))),
				},
			},
			Usage: &types.UsageInfo{
				PromptTokens:     10,
				CompletionTokens: 8,
				TotalTokens:      18,
				Cost: types.Cost{
					TotalCost: 0.001,
				},
			},
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

	// Test successful request
	t.Run("successful request", func(t *testing.T) {
		params := &CompletionParams{
			Messages: []types.ChatMessage{
				types.UserMessage("Hello"),
			},
			Model: "sonar",
		}

		result, err := service.Create(context.Background(), params)
		if err != nil {
			t.Fatalf("Create failed: %v", err)
		}

		if result == nil {
			t.Fatal("Result is nil")
		}

		if result.ID != "test-123" {
			t.Errorf("ID = %s, want test-123", result.ID)
		}

		if len(result.Choices) != 1 {
			t.Fatalf("Choices length = %d, want 1", len(result.Choices))
		}

		if result.Choices[0].Message.Role != types.RoleAssistant {
			t.Errorf("Role = %v, want %v", result.Choices[0].Message.Role, types.RoleAssistant)
		}

		if result.Usage == nil {
			t.Error("Usage is nil")
		} else if result.Usage.TotalTokens != 18 {
			t.Errorf("TotalTokens = %d, want 18", result.Usage.TotalTokens)
		}
	})

	// Test validation errors
	t.Run("nil params", func(t *testing.T) {
		_, err := service.Create(context.Background(), nil)
		if err == nil {
			t.Error("Expected error for nil params")
		}
	})

	t.Run("empty messages", func(t *testing.T) {
		params := &CompletionParams{
			Model: "sonar",
		}
		_, err := service.Create(context.Background(), params)
		if err == nil {
			t.Error("Expected error for empty messages")
		}
	})

	t.Run("empty model", func(t *testing.T) {
		params := &CompletionParams{
			Messages: []types.ChatMessage{
				types.UserMessage("Hello"),
			},
		}
		_, err := service.Create(context.Background(), params)
		if err == nil {
			t.Error("Expected error for empty model")
		}
	})

	t.Run("stream enabled", func(t *testing.T) {
		params := &CompletionParams{
			Messages: []types.ChatMessage{
				types.UserMessage("Hello"),
			},
			Model:  "sonar",
			Stream: types.Bool(true),
		}
		_, err := service.Create(context.Background(), params)
		if err == nil {
			t.Error("Expected error when stream is enabled")
		}
	})
}

func TestCompletionParams_JSON(t *testing.T) {
	params := &CompletionParams{
		Messages: []types.ChatMessage{
			types.SystemMessage("You are helpful"),
			types.UserMessage("Hello"),
		},
		Model:       "sonar",
		MaxTokens:   types.Int(100),
		Temperature: types.Float64(0.7),
		TopP:        types.Float64(0.9),
		Tools: []types.Tool{
			{
				Type: types.ToolTypeFunction,
				Function: types.ToolFunction{
					Name:        "get_weather",
					Description: types.String("Get weather"),
				},
			},
		},
	}

	// Marshal
	data, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal
	var result CompletionParams
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	// Verify
	if len(result.Messages) != 2 {
		t.Errorf("Messages length = %d, want 2", len(result.Messages))
	}
	if result.Model != params.Model {
		t.Errorf("Model = %s, want %s", result.Model, params.Model)
	}
	if result.MaxTokens == nil || *result.MaxTokens != 100 {
		t.Error("MaxTokens mismatch")
	}
	if len(result.Tools) != 1 {
		t.Errorf("Tools length = %d, want 1", len(result.Tools))
	}
}
