package perplexity_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/ZaguanLabs/perplexity-go/perplexity"
	"github.com/ZaguanLabs/perplexity-go/perplexity/api"
	"github.com/ZaguanLabs/perplexity-go/perplexity/embeddings"
	"github.com/ZaguanLabs/perplexity-go/perplexity/monitors"
	"github.com/ZaguanLabs/perplexity-go/perplexity/search"
	"github.com/ZaguanLabs/perplexity-go/perplexity/webcontext"
)

func TestMonitorRoutes(t *testing.T) {
	ctx := context.Background()
	const monitorID = "monitor/a"
	const prefix = "/v1/monitors"
	const member = prefix + "/monitor%2Fa"
	var clearGoal *string
	create := &monitors.CreateMonitorRequest{Name: "news", Schedule: monitors.MonitorScheduleRequest{IntervalScheduleRequest: &monitors.MonitorIntervalScheduleRequest{Type: "interval", Interval: "1d"}}, Targets: monitors.MonitorTargets{{PageTargetRequest: &monitors.MonitorPageTargetRequest{Type: "page", URL: "https://example.org"}}}}
	for _, tc := range []struct {
		method, path, query, body string
		call                      func(*monitors.Service) (any, error)
	}{
		{"POST", prefix, "", `{"name":"news","schedule":{"type":"interval","interval":"1d"},"targets":[{"type":"page","url":"https://example.org"}]}`, func(s *monitors.Service) (any, error) {
			return s.Create(ctx, create, api.WithHeader("Idempotency-Key", "new"))
		}},
		{"GET", prefix, "limit=3", "", func(s *monitors.Service) (any, error) { return s.List(ctx, api.WithQuery("limit", 3)) }},
		{"GET", member, "", "", func(s *monitors.Service) (any, error) { return s.Retrieve(ctx, monitorID) }},
		{"PATCH", member, "", `{"goal":null}`, func(s *monitors.Service) (any, error) {
			return s.Update(ctx, monitorID, &monitors.UpdateMonitorRequest{Goal: &clearGoal})
		}},
		{"DELETE", member, "", "", func(s *monitors.Service) (any, error) { return s.Delete(ctx, monitorID) }},
		{"POST", member + "/pause", "", "", func(s *monitors.Service) (any, error) { return s.Pause(ctx, monitorID) }},
		{"POST", member + "/resume", "", "", func(s *monitors.Service) (any, error) { return s.Resume(ctx, monitorID) }},
		{"POST", member + "/run", "", "", func(s *monitors.Service) (any, error) { return s.Run(ctx, monitorID) }},
		{"GET", member + "/runs", "after=abc", "", func(s *monitors.Service) (any, error) {
			return s.Runs.List(ctx, monitorID, api.WithQuery("after", "abc"))
		}},
		{"GET", member + "/runs/run%2Fb", "", "", func(s *monitors.Service) (any, error) { return s.Runs.Retrieve(ctx, monitorID, "run/b") }},
		{"POST", member + "/webhook/test", "", "", func(s *monitors.Service) (any, error) { return s.Webhook.Test(ctx, monitorID) }},
	} {
		t.Run(tc.method+tc.path, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tc.method || r.URL.EscapedPath() != tc.path || r.URL.RawQuery != tc.query {
					t.Errorf("request = %s %s", r.Method, r.URL)
				}
				if tc.body != "" {
					var got, want any
					if err := json.NewDecoder(r.Body).Decode(&got); err != nil {
						t.Error(err)
					}
					if err := json.Unmarshal([]byte(tc.body), &want); err != nil {
						t.Error(err)
					}
					if !reflect.DeepEqual(got, want) {
						t.Errorf("body = %#v, want %#v", got, want)
					}
				}
				if tc.method == "POST" && tc.path == prefix && r.Header.Get("Idempotency-Key") != "new" {
					t.Error("missing idempotency key")
				}
				io.WriteString(w, `{}`)
			}))
			defer server.Close()
			client, err := perplexity.NewClient("test", perplexity.WithBaseURL(server.URL))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := tc.call(client.Monitors); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestSnapshotAndEmbeddingWire(t *testing.T) {
	var timestamp = time.Date(2026, 11, 24, 13, 25, 0, 0, time.UTC)
	ctx := context.Background()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var body map[string]any
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		switch r.URL.Path {
		case "/search", "/v1/context":
			if body["snapshot_as_of"] != timestamp.Format(time.RFC3339) {
				t.Errorf("snapshot = %v", body)
			}
			if r.URL.Path == "/v1/context" && body["cache_only"] != false {
				t.Errorf("cache_only = %v", body["cache_only"])
			}
			w.Header().Set("X-Request-Id", "ctx-1")
			io.WriteString(w, `{"results":[{"url":"https://example.org","status":{"type":"ok"}}]}`)
		case "/v1/embeddings":
			if body["model"] != "pplx-embed-v2-context-9b" || !reflect.DeepEqual(body["input"], []any{[]any{"a", "b"}}) {
				t.Errorf("embedding body = %v", body)
			}
			io.WriteString(w, `{"data":[{"data":[{"embedding":"abc"}]}]}`)
		default:
			t.Errorf("unexpected URL %s", r.URL)
		}
	}))
	defer server.Close()
	client, err := perplexity.NewClient("test", perplexity.WithBaseURL(server.URL))
	if err != nil {
		t.Fatal(err)
	}
	query := "test"
	_, err = client.Search.Create(ctx, &search.SearchParams{Query: search.Query{Text: &query}, SnapshotAsOf: &timestamp})
	if err != nil {
		t.Fatal(err)
	}
	cache := false
	raw, err := client.Context.CreateRaw(ctx, &webcontext.ContextRequest{URLs: []string{"https://example.org"}, CacheOnly: &cache, SnapshotAsOf: &timestamp})
	if err != nil {
		t.Fatal(err)
	}
	if raw.RequestID != "ctx-1" || len(raw.Data.Results) != 1 || raw.Data.Results[0].Status.Type != "ok" {
		t.Fatalf("context = %+v", raw)
	}
	res, err := client.Embeddings.Create(ctx, &embeddings.CreateParams{Model: embeddings.ModelEmbedV2Context9B, Input: embeddings.Input{Documents: [][]string{{"a", "b"}}}})
	if err != nil {
		t.Fatal(err)
	}
	if len(res.Data) != 1 || len(res.Data[0].Data) != 1 || res.Data[0].Data[0].Embedding == nil || *res.Data[0].Data[0].Embedding != "abc" {
		t.Fatalf("embedding = %+v", res)
	}
	var input embeddings.Input
	if err := json.Unmarshal([]byte(`[{"image":"data:image/png;base64,AA","type":"image"}]`), &input); err != nil || len(input.Images) != 1 {
		t.Fatalf("image = %+v, %v", input, err)
	}
}
