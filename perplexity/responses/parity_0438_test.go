package responses

import (
	"encoding/json"
	"testing"
)

func TestParity0438RequestAndReplay(t *testing.T) {
	// Replay a server-issued approval request and a rejected approval decision.
	wire := `[{"type":"mcp_approval_request","id":"a1","arguments":"{}","name":"write","server_label":"srv","thought_signature":"sig"},{"type":"mcp_approval_response","approval_request_id":"a1","approve":false},{"type":"mcp_call","id":"c1","arguments":"{}","name":"write","server_label":"srv","approval_request_id":"a1"},{"type":"mcp_list_tools","tools":[{"name":"read","input_schema":{},"annotations":{"read_only":false}}]}]`
	var items []InputItem
	if err := json.Unmarshal([]byte(wire), &items); err != nil {
		t.Fatal(err)
	}
	data, err := json.Marshal(items)
	if err != nil {
		t.Fatal(err)
	}
	var got, want any
	json.Unmarshal(data, &got)
	json.Unmarshal([]byte(wire), &want)
	// Input/output MCP calls may omit optional null fields, but all supplied fields survive.
	for i := range got.([]any) {
		for k, v := range want.([]any)[i].(map[string]any) {
			g := got.([]any)[i].(map[string]any)[k]
			a, _ := json.Marshal(g)
			b, _ := json.Marshal(v)
			if string(a) != string(b) {
				t.Fatalf("field %s: %s != %s", k, a, b)
			}
		}
	}
	no := false
	params := CreateParams{Input: Input{Items: items}, ToolChoice: "required", Tools: []Tool{NewToolFromMCP(MCPTool{Type: ToolTypeMCP, ServerLabel: "srv", ServerURL: "https://example.test", RequireApproval: &RequireApproval{Filters: &ApprovalFilters{Never: &ApprovalFilter{ReadOnly: &no}}}}), NewToolFromImageSearch(ImageSearchTool{Type: ToolTypeImageSearch, Filters: &ImageSearchFilters{SafeSearch: &no}})}}
	data, err = json.Marshal(params)
	if err != nil {
		t.Fatal(err)
	}
	var body map[string]any
	json.Unmarshal(data, &body)
	if body["tool_choice"] != "required" {
		t.Fatal(string(data))
	}
	tools := body["tools"].([]any)
	if tools[0].(map[string]any)["require_approval"].(map[string]any)["never"].(map[string]any)["read_only"] != false {
		t.Fatal(string(data))
	}
	if tools[1].(map[string]any)["filters"].(map[string]any)["safe_search"] != false {
		t.Fatal(string(data))
	}
	for _, mode := range []string{"always", "never"} {
		var r RequireApproval
		if err := json.Unmarshal([]byte(`"`+mode+`"`), &r); err != nil {
			t.Fatal(err)
		}
		data, _ := json.Marshal(r)
		if string(data) != `"`+mode+`"` {
			t.Fatal(string(data))
		}
	}
}

func TestParity0438ImageOutputAndStream(t *testing.T) {
	var response CreateResponse
	if err := json.Unmarshal([]byte(`{"output":[{"type":"image_search_results","results":[{"height":10,"width":20,"image_url":"img","origin_url":"src"}]},{"type":"mcp_approval_request","id":"a","arguments":"{}","name":"write","server_label":"s"}]}`), &response); err != nil {
		t.Fatal(err)
	}
	images, ok := response.Output[0].AsImageSearchResults()
	if !ok || images.Results[0].Width != 20 {
		t.Fatal(response)
	}
	if _, ok := response.Output[1].AsMCPApprovalRequest(); !ok {
		t.Fatal(response)
	}
	for _, wire := range []string{`{"type":"response.reasoning.image_search_queries","call_id":"c","sequence_number":1,"queries":["bird"]}`, `{"type":"response.reasoning.image_search_results","call_id":"c","sequence_number":2,"results":[]}`} {
		var event StreamEvent
		if err := json.Unmarshal([]byte(wire), &event); err != nil {
			t.Fatal(err)
		}
		data, err := json.Marshal(event)
		if err != nil {
			t.Fatal(err)
		}
		var a, b any
		json.Unmarshal([]byte(wire), &a)
		json.Unmarshal(data, &b)
		aa, _ := json.Marshal(a)
		bb, _ := json.Marshal(b)
		if string(aa) != string(bb) {
			t.Fatalf("%s != %s", aa, bb)
		}
	}
	filename := "report.pdf"
	data, err := json.Marshal(InputMessageContentPart{Type: InputMessageContentPartTypeFile, Filename: &filename})
	if err != nil {
		t.Fatal(err)
	}
	var part map[string]any
	json.Unmarshal(data, &part)
	if part["filename"] != filename || part["type"] != "input_file" {
		t.Fatal(string(data))
	}
}
