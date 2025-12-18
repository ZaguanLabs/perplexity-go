package http

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// Benchmark HTTP request overhead

func BenchmarkClient_Do_Simple(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client := NewClient(
		&http.Client{Timeout: 30 * time.Second},
		server.URL,
		"test-key",
		0,
		nil,
		"test-agent",
		nil,
	)

	ctx := context.Background()
	req := &Request{
		Method: "GET",
		Path:   "/test",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Do(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_Do_WithJSONBody(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"id": "test-123", "status": "completed"}`))
	}))
	defer server.Close()

	client := NewClient(
		&http.Client{Timeout: 30 * time.Second},
		server.URL,
		"test-key",
		0,
		nil,
		"test-agent",
		nil,
	)

	ctx := context.Background()
	body := map[string]interface{}{
		"model":    "sonar",
		"messages": []map[string]string{{"role": "user", "content": "Hello"}},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := &Request{
			Method: "POST",
			Path:   "/chat/completions",
			Body:   body,
		}
		_, err := client.Do(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_Do_LargeResponse(b *testing.B) {
	largeResponse := strings.Repeat(`{"chunk": "data"}`, 1000)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(largeResponse))
	}))
	defer server.Close()

	client := NewClient(
		&http.Client{Timeout: 30 * time.Second},
		server.URL,
		"test-key",
		0,
		nil,
		"test-agent",
		nil,
	)

	ctx := context.Background()
	req := &Request{
		Method: "GET",
		Path:   "/test",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Do(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_Do_WithRetry(b *testing.B) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts%2 == 1 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client := NewClient(
		&http.Client{Timeout: 30 * time.Second},
		server.URL,
		"test-key",
		1,
		nil,
		"test-agent",
		nil,
	)

	ctx := context.Background()
	req := &Request{
		Method: "GET",
		Path:   "/test",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Do(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_Do_Concurrent(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client := NewClient(
		&http.Client{Timeout: 30 * time.Second},
		server.URL,
		"test-key",
		0,
		nil,
		"test-agent",
		nil,
	)

	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := &Request{
				Method: "GET",
				Path:   "/test",
			}
			_, err := client.Do(ctx, req)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkClient_Do_WithHeaders(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = r.Header.Get("X-Custom-Header")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client := NewClient(
		&http.Client{Timeout: 30 * time.Second},
		server.URL,
		"test-key",
		0,
		nil,
		"test-agent",
		nil,
	)

	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := &Request{
			Method: "GET",
			Path:   "/test",
			Headers: map[string]string{
				"X-Custom-Header": "test-value",
				"X-Request-ID":    "req-123",
			},
		}
		_, err := client.Do(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkClient_Do_Allocations(b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	}))
	defer server.Close()

	client := NewClient(
		&http.Client{Timeout: 30 * time.Second},
		server.URL,
		"test-key",
		0,
		nil,
		"test-agent",
		nil,
	)

	ctx := context.Background()
	req := &Request{
		Method: "GET",
		Path:   "/test",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Do(ctx, req)
		if err != nil {
			b.Fatal(err)
		}
	}
}
