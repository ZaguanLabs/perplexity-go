package perplexity_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/ZaguanLabs/perplexity-go/perplexity"
	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	"github.com/ZaguanLabs/perplexity-go/perplexity/batches"
	"github.com/ZaguanLabs/perplexity-go/perplexity/decisions"
	"github.com/ZaguanLabs/perplexity-go/perplexity/files"
	"github.com/ZaguanLabs/perplexity-go/perplexity/responses"
)

func TestManagement0461WireContracts(t *testing.T) {
	ctx := context.Background()
	var check func(*http.Request)
	var response string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test" {
			t.Error("missing auth")
		}
		if r.URL.Query().Get("tenant") != "test" {
			t.Errorf("base query lost: %s", r.URL)
		}
		check(r)
		w.Header().Set("X-Request-Id", "request-0461")
		io.WriteString(w, response)
	}))
	defer server.Close()
	client, err := perplexity.NewClient("test", perplexity.WithBaseURL(server.URL+"/proxy?tenant=test"))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name, method, path, response string
		call                         func() error
	}{
		{"models", "GET", "/proxy/v1/models", `{"object":"list","data":[{"id":"agent","created":1,"owned_by":"pplx","object":"model"}]}`, func() error {
			r, e := client.Models.ListRaw(ctx)
			if e == nil && (r.RequestID != "request-0461" || r.Data.Data[0].ID != "agent") {
				t.Error(r)
			}
			return e
		}},
		{"decisions", "POST", "/proxy/v1/decisions", `{"answers":{"ok":{"type":"noul","noul":0.7}},"model":"m","usage":{"input_tokens":1,"output_tokens":2}}`, func() error {
			_, e := client.Decisions.Create(ctx, &decisions.DecisionsRequest{Model: "m", State: map[string]any{"ready": false}, Questions: map[string]any{"ok": decisions.DecisionNoulQuestion{Type: "noul"}}})
			return e
		}},
		{"batch create", "POST", "/proxy/v1/batches", `{"id":"b","status":"validating"}`, func() error {
			_, e := client.Batches.Create(ctx, &batches.BatchCreateRequest{Endpoint: "/v1/responses", InputFileID: "f"}, api.WithHeader("Idempotency-Key", "key"))
			return e
		}},
		{"batch list", "GET", "/proxy/v1/batches", `{"data":[],"has_more":false}`, func() error { _, e := client.Batches.List(ctx, api.WithQuery("limit", 0)); return e }},
		{"batch get", "GET", "/proxy/v1/batches/a%2Fb", `{"id":"a/b"}`, func() error { _, e := client.Batches.Retrieve(ctx, "a/b"); return e }},
		{"batch cancel", "POST", "/proxy/v1/batches/b/cancel", `{"status":"cancelling"}`, func() error { _, e := client.Batches.Cancel(ctx, "b"); return e }},
		{"file list", "GET", "/proxy/v1/files", `{"data":[],"has_more":false}`, func() error { _, e := client.Files.List(ctx, api.WithQuery("purpose", "batch")); return e }},
		{"file get", "GET", "/proxy/v1/files/f", `{"id":"f","bytes":10}`, func() error { _, e := client.Files.Retrieve(ctx, "f"); return e }},
		{"file delete", "DELETE", "/proxy/v1/files/f", `{"id":"f","deleted":true}`, func() error { _, e := client.Files.Delete(ctx, "f"); return e }},
		{"file content", "GET", "/proxy/v1/files/f/content", "a\x00b", func() error {
			r, e := client.Files.Content(ctx, "f")
			if e == nil && string(r) != "a\x00b" {
				t.Errorf("binary content: %q", r)
			}
			return e
		}},
		{"skills list", "GET", "/proxy/v1/skills", `{"skills":[],"next_page_token":"next"}`, func() error { _, e := client.Skills.List(ctx); return e }},
		{"skill get", "GET", "/proxy/v1/skills/s", `{"skill_id":"s","revision":"1"}`, func() error { _, e := client.Skills.Retrieve(ctx, "s", api.WithQuery("revision", "1")); return e }},
		{"skill download", "GET", "/proxy/v1/skills/s/download", `{"download_url":"https://example.test/s"}`, func() error { _, e := client.Skills.Download(ctx, "s"); return e }},
		{"skill revisions", "GET", "/proxy/v1/skills/s/revisions", `{"revisions":[{"revision":"1"}]}`, func() error { _, e := client.Skills.Revisions.List(ctx, "s"); return e }},
		{"skill delete", "DELETE", "/proxy/v1/skills/s", "", func() error { return client.Skills.Delete(ctx, "s", "1") }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			check = func(r *http.Request) {
				if r.Method != tc.method || r.URL.EscapedPath() != tc.path {
					t.Errorf("got %s %s", r.Method, r.URL.EscapedPath())
				}
			}
			response = tc.response
			if err := tc.call(); err != nil {
				t.Fatal(err)
			}
		})
	}
	for _, id := range []string{"", ".", ".."} {
		if _, e := client.Files.Retrieve(ctx, id); e == nil {
			t.Errorf("accepted bad path %q", id)
		}
	}
}

func TestUploads0461(t *testing.T) {
	const payload = "{\"custom_id\":\"one\"}\n"
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if err := r.ParseMultipartForm(1024); err != nil {
			t.Error(err)
			return
		}
		f, h, err := r.FormFile("file")
		if err != nil {
			t.Error(err)
			return
		}
		defer f.Close()
		data, _ := io.ReadAll(f)
		if string(data) != payload || h.Filename != "batch.jsonl" {
			t.Errorf("bad upload %q %q", data, h.Filename)
		}
		if r.URL.Path == "/v1/files" && r.FormValue("purpose") != "batch" {
			t.Error("missing purpose")
		}
		if requests == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(500)
			return
		}
		io.WriteString(w, `{"id":"f","skill_id":"s"}`)
	}))
	defer server.Close()
	client, err := perplexity.NewClient("test", perplexity.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	file, err := os.CreateTemp(t.TempDir(), "upload")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	if _, err = file.WriteString(payload); err != nil {
		t.Fatal(err)
	}
	if _, err = file.Seek(0, 0); err != nil {
		t.Fatal(err)
	}
	if _, err = client.Files.Create(context.Background(), &files.CreateParams{Purpose: "batch", File: api.Upload{Filename: "batch.jsonl", Reader: file}}); err != nil {
		t.Fatal(err)
	}
	if _, err = client.Skills.Create(context.Background(), api.Upload{Filename: "batch.jsonl", Reader: strings.NewReader(payload)}); err != nil {
		t.Fatal(err)
	}
	if _, err = client.Skills.Update(context.Background(), "s", "1", api.Upload{Filename: "batch.jsonl", Reader: strings.NewReader(payload)}); err != nil {
		t.Fatal(err)
	}
	if requests != 4 {
		t.Errorf("requests=%d", requests)
	}
}

func TestResponses0461FieldsAndReconnect(t *testing.T) {
	f := false
	effort := "high"
	key := "cache"
	tier := "fast"
	data, err := json.Marshal(responses.CreateParams{Effort: &effort, PromptCacheKey: &key, ServiceTier: &tier, StepBudgetReminder: &f})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"step_budget_reminder":false`) {
		t.Fatal(string(data))
	}
	var usage responses.Usage
	if err = json.Unmarshal([]byte(`{"tool_calls_details":{"extract":{"cost_usd":0.2,"token_usage":[{"model":"m","input_tokens_details":{"cached_tokens":2},"output_tokens_details":{"reasoning_tokens":3}}]}}}`), &usage); err != nil {
		t.Fatal(err)
	}
	if usage.ToolCallsDetails["extract"].TokenUsage[0].InputTokensDetails.CachedTokens != 2 {
		t.Fatal(usage)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("starting_after") != "0" || r.URL.Query().Get("stream") != "true" {
			t.Error(r.URL)
		}
		w.Header().Set("Content-Type", "text/event-stream")
		io.WriteString(w, "data: {\"type\":\"response.output_text.delta\",\"delta\":\"hi\"}\n\ndata: [DONE]\n\n")
	}))
	defer server.Close()
	client, err := perplexity.NewClient("test", perplexity.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	if _, err = client.Responses.RetrieveStream(context.Background(), "r", nil); err == nil {
		t.Fatal("missing cursor accepted")
	}
	zero := int64(0)
	stream, err := client.Responses.RetrieveStream(context.Background(), "r", &responses.RetrieveParams{StartingAfter: &zero})
	if err != nil {
		t.Fatal(err)
	}
	defer stream.Close()
	if _, err = stream.Next(); err != nil {
		t.Fatal(err)
	}
}
