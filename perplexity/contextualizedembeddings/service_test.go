package contextualizedembeddings

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
		if r.URL.Path != "/v1/contextualizedembeddings" {
			t.Errorf("Expected path /v1/contextualizedembeddings, got %s", r.URL.Path)
		}

		var params CreateParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}
		if len(params.Input) != 1 || len(params.Input[0]) != 2 {
			t.Fatalf("unexpected input shape: %#v", params.Input)
		}

		response := CreateResponse{
			Data:   []ContextualizedEmbeddingObject{{Data: []EmbeddingObject{{Embedding: types.String("AQID"), Index: types.Int(0), Object: types.String("embedding")}}, Index: types.Int(0), Object: types.String("contextualized_embedding")}},
			Model:  types.String(string(ModelEmbedContextV14B)),
			Object: types.String("list"),
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := internalhttp.NewClient(server.Client(), server.URL, "test-api-key", 2, nil, "test-agent", nil)
	service := NewService(httpClient)

	result, err := service.Create(context.Background(), &CreateParams{Input: [][]string{{"chunk-1", "chunk-2"}}, Model: ModelEmbedContextV14B})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if result == nil {
		t.Fatal("Result is nil")
	}
	if len(result.Data) != 1 {
		t.Fatalf("Data length = %d, want 1", len(result.Data))
	}
}

func TestService_CreateValidation(t *testing.T) {
	service := NewService(internalhttp.NewClient(&http.Client{}, "https://example.com", "test-api-key", 0, nil, "test-agent", nil))

	if _, err := service.Create(context.Background(), nil); err == nil {
		t.Fatal("expected error for nil params")
	}
	if _, err := service.Create(context.Background(), &CreateParams{Input: [][]string{{"chunk"}}}); err == nil {
		t.Fatal("expected error for missing model")
	}
	if _, err := service.Create(context.Background(), &CreateParams{Input: nil, Model: ModelEmbedContextV106B}); err == nil {
		t.Fatal("expected error for empty input")
	}
	if _, err := service.Create(context.Background(), &CreateParams{Input: [][]string{{""}}, Model: ModelEmbedContextV106B}); err == nil {
		t.Fatal("expected error for empty chunk")
	}
}
