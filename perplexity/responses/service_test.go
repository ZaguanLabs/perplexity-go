package responses

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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
		if params.Background == nil || !*params.Background {
			t.Fatal("background should be present")
		}
		if len(params.Tools) != 4 {
			t.Fatalf("tools length = %d, want 4", len(params.Tools))
		}
		if _, ok := params.Tools[0].AsWebSearch(); !ok {
			t.Fatal("expected web_search tool")
		}
		if _, ok := params.Tools[1].AsPeopleSearch(); !ok {
			t.Fatal("expected people_search tool")
		}
		if _, ok := params.Tools[2].AsFinanceSearch(); !ok {
			t.Fatal("expected finance_search tool")
		}
		if _, ok := params.Tools[3].AsSandbox(); !ok {
			t.Fatal("expected sandbox tool")
		}

		response := CreateResponse{
			ID:         "resp_123",
			CreatedAt:  1234567890,
			Model:      "sonar-pro",
			Object:     "response",
			Status:     StatusCompleted,
			Background: types.Bool(true),
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

	searchContextSize := SearchContextSizeLow
	result, err := service.Create(context.Background(), &CreateParams{
		Input:      Input{Text: types.String("hello")},
		Background: types.Bool(true),
		Model:      types.String("sonar-pro"),
		Tools: []Tool{
			NewToolFromWebSearch(WebSearchTool{
				Type:              ToolTypeWebSearch,
				SearchContextSize: &searchContextSize,
			}),
			NewToolFromPeopleSearch(PeopleSearchTool{Type: ToolTypePeopleSearch}),
			NewToolFromFinanceSearch(FinanceSearchTool{Type: ToolTypeFinanceSearch}),
			NewToolFromSandbox(SandboxTool{Type: ToolTypeSandbox}),
		},
	})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
	if result.ID != "resp_123" {
		t.Fatalf("ID = %s, want resp_123", result.ID)
	}
	if got := result.OutputText(); got != "Hello from responses" {
		t.Fatalf("OutputText = %q", got)
	}
	if result.Background == nil || !*result.Background {
		t.Fatal("Background should be true")
	}
}

func TestService_Retrieve(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if r.URL.Path != "/v1/responses/resp_123" {
			t.Errorf("Expected path /v1/responses/resp_123, got %s", r.URL.Path)
		}

		response := RetrieveResponse{
			ID:         "resp_123",
			CreatedAt:  1234567890,
			Model:      "sonar-pro",
			Object:     "response",
			Status:     StatusQueued,
			Background: types.Bool(true),
			Output:     []OutputItem{},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	httpClient := internalhttp.NewClient(server.Client(), server.URL, "test-api-key", 2, nil, "test-agent", nil)
	service := NewService(httpClient)

	result, err := service.Retrieve(context.Background(), "resp_123")
	if err != nil {
		t.Fatalf("Retrieve failed: %v", err)
	}
	if result.ID != "resp_123" {
		t.Fatalf("ID = %s, want resp_123", result.ID)
	}
	if result.Status != StatusQueued {
		t.Fatalf("Status = %s, want %s", result.Status, StatusQueued)
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
	if _, err := service.Retrieve(context.Background(), ""); err == nil {
		t.Fatal("expected error for empty response ID")
	}
}

func TestService_CancelAndFiles(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/responses/resp_123/cancel":
			if r.Method != http.MethodPost {
				t.Fatalf("method = %s", r.Method)
			}
			_, _ = io.WriteString(w, `{"response_id":"resp_123","status":"cancelling"}`)
		case "/v1/responses/resp_123/files":
			_, _ = io.WriteString(w, `{"object":"list","data":[{"id":"file_1","bytes":4,"created_at":1,"filename":"a.txt","object":"file"}]}`)
		case "/v1/responses/resp_123/files/file_1/content":
			if got := r.Header.Get("Accept"); got != "application/octet-stream" {
				t.Fatalf("Accept = %q", got)
			}
			_, _ = io.WriteString(w, "data")
		default:
			t.Fatalf("unexpected path %s", r.URL.Path)
		}
	}))
	defer server.Close()

	service := NewService(internalhttp.NewClient(server.Client(), server.URL, "key", 0, nil, "agent", nil))
	cancelled, err := service.Cancel(context.Background(), "resp_123")
	if err != nil || cancelled.Status != "cancelling" {
		t.Fatalf("Cancel = %#v, %v", cancelled, err)
	}
	files, err := service.Files.List(context.Background(), "resp_123")
	if err != nil || len(files.Data) != 1 || files.Data[0].Filename != "a.txt" {
		t.Fatalf("List = %#v, %v", files, err)
	}
	content, err := service.Files.Content(context.Background(), "resp_123", "file_1")
	if err != nil || string(content) != "data" {
		t.Fatalf("Content = %q, %v", content, err)
	}
}

func TestNewResponseVariants(t *testing.T) {
	payload := `{"type":"mcp_call","id":"mcp_1","arguments":"{}","name":"lookup","server_label":"docs","output":"ok"}`
	var item OutputItem
	if err := json.Unmarshal([]byte(payload), &item); err != nil {
		t.Fatal(err)
	}
	call, ok := item.AsMCPCall()
	if !ok || call.Output == nil || *call.Output != "ok" {
		t.Fatalf("MCP call = %#v", call)
	}

	skill := NewSkillFromInline(InlineSkill{Type: SkillTypeInline, Name: "review", Description: "Review", Instructions: "Check it"})
	data, err := json.Marshal(CreateParams{Input: Input{Text: types.String("hello")}, Skills: []Skill{skill}, Store: types.Bool(false), PreviousResponseID: types.String("resp_prev")})
	if err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{`"skills"`, `"store":false`, `"previous_response_id":"resp_prev"`} {
		if !strings.Contains(string(data), expected) {
			t.Fatalf("payload %s missing %s", data, expected)
		}
	}

	var event StreamEvent
	if err := json.Unmarshal([]byte(`{"type":"response.skill.loaded","name":"office/pdf","sequence_number":4}`), &event); err != nil {
		t.Fatal(err)
	}
	loaded, ok := event.AsSkillLoaded()
	if !ok || loaded.Name != "office/pdf" {
		t.Fatalf("event = %#v", loaded)
	}
}

func TestPython0433ResponseAdditions(t *testing.T) {
	connector := NewToolFromConnector(ConnectorTool{
		ID:                "conn_1",
		ServerLabel:       "github",
		Type:              ToolTypeConnector,
		AllowedTools:      []string{"search"},
		ServerDescription: types.String("GitHub connector"),
	})
	deferLoading := true
	temperature := 0.4
	topP := 0.8
	payload, err := json.Marshal(CreateParams{
		Input:       Input{Text: types.String("hello")},
		Temperature: &temperature,
		TopP:        &topP,
		Tools: []Tool{
			connector,
			NewToolFromMCP(MCPTool{
				ServerLabel:  "docs",
				ServerURL:    "https://example.test/mcp",
				Type:         ToolTypeMCP,
				DeferLoading: &deferLoading,
			}),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, expected := range []string{
		`"type":"connector"`,
		`"server_description":"GitHub connector"`,
		`"defer_loading":true`,
		`"temperature":0.4`,
		`"top_p":0.8`,
	} {
		if !strings.Contains(string(payload), expected) {
			t.Fatalf("payload %s missing %s", payload, expected)
		}
	}

	var decoded Tool
	if err := json.Unmarshal([]byte(`{"id":"conn_1","server_label":"github","type":"connector"}`), &decoded); err != nil {
		t.Fatal(err)
	}
	if value, ok := decoded.AsConnector(); !ok || value.ID != "conn_1" {
		t.Fatalf("connector = %#v, %t", value, ok)
	}

	toolSearchJSON := `{"type":"tool_search_output","id":"ts_1","call_id":null,"execution":"server","status":"completed","tools":[{"description":"repo tools","name":"github","type":"namespace","tools":[{"name":"search","type":"function","parameters":{"type":"object"}}]}]}`
	var item OutputItem
	if err := json.Unmarshal([]byte(toolSearchJSON), &item); err != nil {
		t.Fatal(err)
	}
	search, ok := item.AsToolSearch()
	if !ok || len(search.Tools) != 1 || len(search.Tools[0].Tools) != 1 || search.Tools[0].Tools[0].Name != "search" {
		t.Fatalf("tool search = %#v", search)
	}

	connectorID := "conn_1"
	mcp := NewOutputItemFromMCPCall(MCPCallOutputItem{
		ID: "mcp_1", Arguments: "{}", Name: "search", ServerLabel: "github",
		Type: OutputItemTypeMCPCall, ConnectorID: &connectorID,
	})
	mcpJSON, err := json.Marshal(mcp)
	if err != nil || !strings.Contains(string(mcpJSON), `"connector_id":"conn_1"`) {
		t.Fatalf("MCP output = %s, %v", mcpJSON, err)
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
