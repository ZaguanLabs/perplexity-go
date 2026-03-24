package embeddings

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
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/v1/embeddings" {
			t.Errorf("Expected path /v1/embeddings, got %s", r.URL.Path)
		}

		var params CreateParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}
		if params.Model != ModelEmbedV14B {
			t.Errorf("Model = %s, want %s", params.Model, ModelEmbedV14B)
		}

		response := CreateResponse{
			Data:   []EmbeddingObject{{Embedding: types.String("AQID"), Index: types.Int(0), Object: types.String("embedding")}},
			Model:  types.String(string(ModelEmbedV14B)),
			Object: types.String("list"),
			Usage:  &Usage{PromptTokens: types.Int(3), TotalTokens: types.Int(3)},
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := internalhttp.NewClient(server.Client(), server.URL, "test-api-key", 2, nil, "test-agent", nil)
	service := NewService(httpClient)

	result, err := service.Create(context.Background(), &CreateParams{Input: Input{Texts: []string{"hello"}}, Model: ModelEmbedV14B})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(result.Data) != 1 {
		t.Fatalf("Data length = %d, want 1", len(result.Data))
	}
	if result.Model == nil || *result.Model != string(ModelEmbedV14B) {
		t.Fatalf("Model = %v, want %s", result.Model, ModelEmbedV14B)
	}
}

func TestService_CreateValidation(t *testing.T) {
	service := NewService(internalhttp.NewClient(&http.Client{}, "https://example.com", "test-api-key", 0, nil, "test-agent", nil))

	if _, err := service.Create(context.Background(), nil); err == nil {
		t.Fatal("expected error for nil params")
	}
	if _, err := service.Create(context.Background(), &CreateParams{Input: Input{Text: types.String("hello")}}); err == nil {
		t.Fatal("expected error for missing model")
	}
	if _, err := service.Create(context.Background(), &CreateParams{Input: Input{Text: types.String("")}, Model: ModelEmbedV106B}); err == nil {
		t.Fatal("expected error for empty string input")
	}
	if _, err := service.Create(context.Background(), &CreateParams{Input: Input{Texts: []string{"ok", ""}}, Model: ModelEmbedV106B}); err == nil {
		t.Fatal("expected error for empty string in input slice")
	}
}
