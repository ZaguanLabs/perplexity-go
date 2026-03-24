package browser

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	internalhttp "github.com/ZaguanLabs/perplexity-go/perplexity/internal/http"
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

func TestSessionsService_CreateAndDelete(t *testing.T) {
	deleteCalled := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/v1/browser/sessions":
			response := SessionResponse{SessionID: types.String("sess_123"), Status: func() *SessionStatus { s := SessionStatusRunning; return &s }()}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(response)
		case r.Method == http.MethodDelete && r.URL.Path == "/v1/browser/sessions/sess_123":
			deleteCalled = true
			if r.Header.Get("Accept") != "*/*" {
				t.Errorf("Accept header = %q, want */*", r.Header.Get("Accept"))
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Fatalf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
	}))
	defer server.Close()

	httpClient := internalhttp.NewClient(server.Client(), server.URL, "test-api-key", 2, nil, "test-agent", nil)
	service := NewService(httpClient)

	created, err := service.Sessions.Create(context.Background())
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if created == nil || created.SessionID == nil || *created.SessionID != "sess_123" {
		t.Fatalf("unexpected create response: %#v", created)
	}

	if err := service.Sessions.Delete(context.Background(), "sess_123"); err != nil {
		t.Fatalf("Delete failed: %v", err)
	}
	if !deleteCalled {
		t.Fatal("expected delete to be called")
	}
}

func TestSessionsService_DeleteValidation(t *testing.T) {
	service := NewService(internalhttp.NewClient(&http.Client{}, "https://example.com", "test-api-key", 0, nil, "test-agent", nil))
	if err := service.Sessions.Delete(context.Background(), ""); err == nil {
		t.Fatal("expected error for empty session ID")
	}
}
