package responses

import (
	"context"
	"encoding/json"
	"io"
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
		if r.URL.Path != "/v1/responses" {
			t.Errorf("Expected path /v1/responses, got %s", r.URL.Path)
		}

		var params CreateParams
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			t.Fatalf("Failed to decode request body: %v", err)
		}
		if params.Input.Text == nil || *params.Input.Text != "hello" {
			t.Fatal("input should be present")
		}

		response := CreateResponse{
			ID:        "resp_123",
			CreatedAt: 1234567890,
			Model:     "sonar-pro",
			Object:    "response",
			Status:    StatusCompleted,
			Output: []OutputItem{NewOutputItemFromMessage(MessageOutputItem{
				ID:     "msg_1",
				Type:   OutputItemTypeMessage,
				Role:   "assistant",
				Status: StatusCompleted,
				Content: []ContentPart{{
					Type: "output_text",
					Text: "Hello from responses",
				}},
			})},
			Usage: &Usage{InputTokens: 5, OutputTokens: 4, TotalTokens: 9},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := internalhttp.NewClient(server.Client(), server.URL, "test-api-key", 2, nil, "test-agent", nil)
	service := NewService(httpClient)

	result, err := service.Create(context.Background(), &CreateParams{Input: Input{Text: types.String("hello")}, Model: types.String("sonar-pro")})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if result.ID != "resp_123" {
		t.Fatalf("ID = %s, want resp_123", result.ID)
	}
	if got := result.OutputText(); got != "Hello from responses" {
		t.Fatalf("OutputText = %q", got)
	}
}

func TestService_CreateStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = w.Write([]byte("data: {\"sequence_number\":1,\"type\":\"response.output_text.delta\",\"delta\":\"Hello\"}\n\n"))
		_, _ = w.Write([]byte("data: [DONE]\n\n"))
		if flusher, ok := w.(http.Flusher); ok {
			flusher.Flush()
		}
	}))
	defer server.Close()

	httpClient := internalhttp.NewClient(server.Client(), server.URL, "test-api-key", 2, nil, "test-agent", nil)
	service := NewService(httpClient)

	stream, err := service.CreateStream(context.Background(), &CreateParams{Input: Input{Text: types.String("hello")}, Model: types.String("sonar-pro")})
	if err != nil {
		t.Fatalf("CreateStream failed: %v", err)
	}
	defer func() { _ = stream.Close() }()

	event, err := stream.Next()
	if err != nil {
		t.Fatalf("Next failed: %v", err)
	}
	delta, ok := event.AsTextDelta()
	if !ok {
		t.Fatalf("expected text delta event, got %#v", event)
	}
	if delta.Type != EventTypeOutputTextDelta {
		t.Fatalf("Type = %s, want %s", delta.Type, EventTypeOutputTextDelta)
	}
	if delta.Delta != "Hello" {
		t.Fatalf("Delta = %v, want Hello", delta.Delta)
	}

	_, err = stream.Next()
	if err != io.EOF {
		t.Fatalf("expected EOF, got %v", err)
	}
}

func TestService_CreateValidation(t *testing.T) {
	service := NewService(internalhttp.NewClient(&http.Client{}, "https://example.com", "test-api-key", 0, nil, "test-agent", nil))

	if _, err := service.Create(context.Background(), nil); err == nil {
		t.Fatal("expected error for nil params")
	}
	if _, err := service.Create(context.Background(), &CreateParams{}); err == nil {
		t.Fatal("expected error for missing input")
	}
	streamEnabled := true
	if _, err := service.Create(context.Background(), &CreateParams{Input: Input{Text: types.String("hello")}, Stream: &streamEnabled}); err == nil {
		t.Fatal("expected error when stream is enabled")
	}
}

func TestInput_JSONVariants(t *testing.T) {
	t.Run("string input", func(t *testing.T) {
		payload := Input{Text: types.String("hello")}
		data, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		if string(data) != `"hello"` {
			t.Fatalf("Marshal = %s", data)
		}

		var decoded Input
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if decoded.Text == nil || *decoded.Text != "hello" {
			t.Fatalf("decoded text = %#v", decoded.Text)
		}
	})

	t.Run("input item array", func(t *testing.T) {
		payload := Input{Items: []InputItem{NewInputItemFromMessage(InputMessage{Role: InputMessageRoleUser, Type: InputMessageTypeMessage, Content: InputMessageContent{Text: types.String("hi")}})}}
		data, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}

		var decoded Input
		if err := json.Unmarshal(data, &decoded); err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}
		if len(decoded.Items) != 1 {
			t.Fatalf("Items length = %d, want 1", len(decoded.Items))
		}
		message, ok := decoded.Items[0].AsInputMessage()
		if !ok {
			t.Fatal("expected input message variant")
		}
		if message.Content.Text == nil || *message.Content.Text != "hi" {
			t.Fatalf("message content = %#v", message.Content)
		}
	})
}

func TestOutputItem_JSONVariant(t *testing.T) {
	payload := NewOutputItemFromFunctionCall(FunctionCallOutputItem{
		ID:        "fc_1",
		Arguments: `{"city":"SF"}`,
		CallID:    "call_1",
		Name:      "get_weather",
		Status:    StatusCompleted,
		Type:      OutputItemTypeFunctionCall,
	})

	data, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	var decoded OutputItem
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	fc, ok := decoded.AsFunctionCall()
	if !ok {
		t.Fatal("expected function call variant")
	}
	if fc.CallID != "call_1" {
		t.Fatalf("CallID = %s, want call_1", fc.CallID)
	}
}
