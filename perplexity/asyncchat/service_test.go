package asyncchat

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ZaguanLabs/perplexity-go/perplexity/chat"
	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

func TestService_CreateListGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/async/chat/completions":
			var body CompletionCreateParams
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Fatalf("failed to decode create body: %v", err)
			}
			if body.Request == nil {
				t.Fatal("expected request body")
			}
			if body.Request.Model != "sonar" {
				t.Fatalf("model = %q, want sonar", body.Request.Model)
			}
			if body.IdempotencyKey == nil || *body.IdempotencyKey != "idem_123" {
				t.Fatalf("idempotency_key = %#v, want idem_123", body.IdempotencyKey)
			}
			response := CompletionCreateResponse{
				ID:        "req_123",
				CreatedAt: 123,
				Model:     "sonar",
				Status:    CompletionStatusCreated,
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response)
		case r.Method == http.MethodGet && r.URL.Path == "/async/chat/completions":
			response := CompletionListResponse{
				Requests: []CompletionListRequest{{
					ID:        "req_123",
					CreatedAt: 123,
					Model:     "sonar",
					Status:    CompletionStatusInProgress,
				}},
				NextToken: types.String("next_123"),
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response)
		case r.Method == http.MethodGet && r.URL.Path == "/async/chat/completions/req_123":
			if got := r.URL.Query().Get("local_mode"); got != "true" {
				t.Fatalf("local_mode = %q, want true", got)
			}
			if got := r.Header.Get("x-client-env"); got != "dev" {
				t.Fatalf("x-client-env = %q, want dev", got)
			}
			if got := r.Header.Get("x-client-name"); got != "sdk-tests" {
				t.Fatalf("x-client-name = %q, want sdk-tests", got)
			}
			response := CompletionGetResponse{
				ID:        "req_123",
				CreatedAt: 123,
				Model:     "sonar",
				Status:    CompletionStatusCompleted,
				Response: &types.StreamChunk{
					ID:    "chatcmpl_123",
					Model: "sonar",
				},
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.String())
		}
	}))
	defer server.Close()

	httpClient := internalhttp.NewClient(server.Client(), server.URL, "test-api-key", 0, nil, "test-agent", nil)
	service := NewService(httpClient)

	createResp, err := service.Create(context.Background(), &CompletionCreateParams{
		Request: &chat.CompletionParams{
			Model:    "sonar",
			Messages: []types.ChatMessage{types.UserMessage("hello")},
		},
		IdempotencyKey: types.String("idem_123"),
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if createResp.ID != "req_123" {
		t.Fatalf("Create ID = %q, want req_123", createResp.ID)
	}

	listResp, err := service.List(context.Background())
	if err != nil {
		t.Fatalf("List failed: %v", err)
	}
	if len(listResp.Requests) != 1 {
		t.Fatalf("List requests len = %d, want 1", len(listResp.Requests))
	}
	if listResp.NextToken == nil || *listResp.NextToken != "next_123" {
		t.Fatalf("NextToken = %#v, want next_123", listResp.NextToken)
	}

	getResp, err := service.Get(context.Background(), "req_123", &CompletionGetParams{
		LocalMode:   types.Bool(true),
		XClientEnv:  types.String("dev"),
		XClientName: types.String("sdk-tests"),
	})
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if getResp.Status != CompletionStatusCompleted {
		t.Fatalf("Get status = %q, want %q", getResp.Status, CompletionStatusCompleted)
	}
	if getResp.Response == nil || getResp.Response.ID != "chatcmpl_123" {
		t.Fatalf("Get response = %#v, want stream chunk with id chatcmpl_123", getResp.Response)
	}
}

func TestService_CreateValidation(t *testing.T) {
	service := NewService(internalhttp.NewClient(&http.Client{}, "https://example.com", "test-api-key", 0, nil, "test-agent", nil))

	if _, err := service.Create(context.Background(), nil); err == nil {
		t.Fatal("expected error for nil params")
	}
	if _, err := service.Create(context.Background(), &CompletionCreateParams{}); err == nil {
		t.Fatal("expected error for nil request")
	}
	if _, err := service.Create(context.Background(), &CompletionCreateParams{Request: &chat.CompletionParams{}}); err == nil {
		t.Fatal("expected error for missing request fields")
	}
}

func TestService_GetValidation(t *testing.T) {
	service := NewService(internalhttp.NewClient(&http.Client{}, "https://example.com", "test-api-key", 0, nil, "test-agent", nil))
	if _, err := service.Get(context.Background(), "", nil); err == nil {
		t.Fatal("expected error for empty api request")
	}
}
