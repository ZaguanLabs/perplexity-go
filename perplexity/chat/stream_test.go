package chat

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	internalhttp "github.com/perplexityai/perplexity-go/perplexity/internal/http"
	"github.com/perplexityai/perplexity-go/perplexity/types"
)

func TestService_CreateStream(t *testing.T) {
	// Create a mock SSE server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify streaming headers
		if r.Header.Get("Accept") != "text/event-stream" {
			t.Error("Missing Accept: text/event-stream header")
		}

		// Set SSE headers
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		// Write SSE events
		flusher, ok := w.(http.Flusher)
		if !ok {
			t.Fatal("Streaming not supported")
		}

		// Event 1
		w.Write([]byte(`data: {"id":"test-1","model":"sonar","created":1234567890,"choices":[{"index":0,"delta":{"role":"assistant","content":"Hello"},"finish_reason":null}]}` + "\n\n"))
		flusher.Flush()

		// Event 2
		w.Write([]byte(`data: {"id":"test-1","model":"sonar","created":1234567890,"choices":[{"index":0,"delta":{"role":"assistant","content":" world"},"finish_reason":null}]}` + "\n\n"))
		flusher.Flush()

		// Event 3 (final)
		w.Write([]byte(`data: {"id":"test-1","model":"sonar","created":1234567890,"choices":[{"index":0,"delta":{"role":"assistant","content":"!"},"finish_reason":"stop"}],"usage":{"prompt_tokens":5,"completion_tokens":3,"total_tokens":8,"cost":{"total_cost":0.001}}}` + "\n\n"))
		flusher.Flush()

		// Done marker
		w.Write([]byte("data: [DONE]\n\n"))
		flusher.Flush()
	}))
	defer server.Close()

	// Create service
	httpClient := internalhttp.NewClient(
		server.Client(),
		server.URL,
		"test-api-key",
		2,
		nil,
		"test-agent",
	)
	service := NewService(httpClient)

	// Test streaming
	t.Run("successful stream", func(t *testing.T) {
		params := &CompletionParams{
			Messages: []types.ChatMessage{
				types.UserMessage("Hello"),
			},
			Model: "sonar",
		}

		stream, err := service.CreateStream(context.Background(), params)
		if err != nil {
			t.Fatalf("CreateStream failed: %v", err)
		}
		defer stream.Close()

		// Read chunks
		var chunks []*types.StreamChunk
		for {
			chunk, err := stream.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				t.Fatalf("Next() error: %v", err)
			}
			chunks = append(chunks, chunk)
		}

		// Verify we got 3 chunks (before [DONE])
		if len(chunks) != 3 {
			t.Fatalf("Expected 3 chunks, got %d", len(chunks))
		}

		// Verify first chunk
		if chunks[0].ID != "test-1" {
			t.Errorf("Chunk 0 ID = %s, want test-1", chunks[0].ID)
		}

		// Verify last chunk has usage
		if chunks[2].Usage == nil {
			t.Error("Last chunk missing usage")
		} else if chunks[2].Usage.TotalTokens != 8 {
			t.Errorf("TotalTokens = %d, want 8", chunks[2].Usage.TotalTokens)
		}
	})

	// Test validation
	t.Run("nil params", func(t *testing.T) {
		_, err := service.CreateStream(context.Background(), nil)
		if err == nil {
			t.Error("Expected error for nil params")
		}
	})

	t.Run("empty messages", func(t *testing.T) {
		params := &CompletionParams{
			Model: "sonar",
		}
		_, err := service.CreateStream(context.Background(), params)
		if err == nil {
			t.Error("Expected error for empty messages")
		}
	})
}

func TestStream_Next(t *testing.T) {
	// Create a mock response with SSE data
	sseData := `data: {"id":"test-1","model":"sonar","created":1234567890,"choices":[{"index":0,"delta":{"role":"assistant","content":"Hello"},"finish_reason":null}]}

data: [DONE]

`
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(sseData)),
		Header:     make(http.Header),
	}

	stream := newStream(context.Background(), resp)
	defer stream.Close()

	// First chunk
	chunk, err := stream.Next()
	if err != nil {
		t.Fatalf("Next() error: %v", err)
	}
	if chunk.ID != "test-1" {
		t.Errorf("ID = %s, want test-1", chunk.ID)
	}

	// Done marker
	_, err = stream.Next()
	if err != io.EOF {
		t.Errorf("Expected EOF, got %v", err)
	}
}

func TestStream_ContextCancellation(t *testing.T) {
	// Create a long-running SSE stream
	sseData := `data: {"id":"test-1","model":"sonar","created":1234567890,"choices":[{"index":0,"delta":{"role":"assistant","content":"Hello"},"finish_reason":null}]}

`
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(sseData)),
		Header:     make(http.Header),
	}

	ctx, cancel := context.WithCancel(context.Background())
	stream := newStream(ctx, resp)
	defer stream.Close()

	// Read first chunk
	_, err := stream.Next()
	if err != nil {
		t.Fatalf("First Next() error: %v", err)
	}

	// Cancel context
	cancel()

	// Try to read next chunk
	_, err = stream.Next()
	if err != context.Canceled {
		t.Errorf("Expected context.Canceled, got %v", err)
	}
}

func TestStream_ErrorEvent(t *testing.T) {
	sseData := `event: error
data: something went wrong

`
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(sseData)),
		Header:     make(http.Header),
	}

	stream := newStream(context.Background(), resp)
	defer stream.Close()

	_, err := stream.Next()
	if err == nil {
		t.Error("Expected error for error event")
	}
	if !strings.Contains(err.Error(), "stream error") {
		t.Errorf("Error message = %v, want to contain 'stream error'", err)
	}
}

func TestStream_Iter(t *testing.T) {
	sseData := `data: {"id":"test-1","model":"sonar","created":1234567890,"choices":[{"index":0,"delta":{"role":"assistant","content":"Hello"},"finish_reason":null}]}

data: {"id":"test-2","model":"sonar","created":1234567891,"choices":[{"index":0,"delta":{"role":"assistant","content":" world"},"finish_reason":"stop"}]}

data: [DONE]

`
	resp := &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(sseData)),
		Header:     make(http.Header),
	}

	stream := newStream(context.Background(), resp)
	defer stream.Close()

	var chunks []*types.StreamChunk
	for chunk := range stream.Iter() {
		chunks = append(chunks, chunk)
	}

	if len(chunks) != 2 {
		t.Errorf("Expected 2 chunks, got %d", len(chunks))
	}
}
