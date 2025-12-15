package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	httpClient := &http.Client{}
	baseURL := "https://api.example.com"
	apiKey := "test-key"
	maxRetries := 3
	headers := map[string]string{"X-Custom": "value"}
	userAgent := "test-agent"

	client := NewClient(httpClient, baseURL, apiKey, maxRetries, headers, userAgent, nil)

	if client.httpClient != httpClient {
		t.Error("httpClient not set correctly")
	}
	if client.baseURL != baseURL {
		t.Error("baseURL not set correctly")
	}
	if client.apiKey != apiKey {
		t.Error("apiKey not set correctly")
	}
	if client.maxRetries != maxRetries {
		t.Error("maxRetries not set correctly")
	}
	if client.userAgent != userAgent {
		t.Error("userAgent not set correctly")
	}
}

func TestClient_Do_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify headers
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Error("Authorization header incorrect")
		}
		if r.Header.Get("User-Agent") != "test-agent" {
			t.Error("User-Agent header incorrect")
		}
		if r.Header.Get("X-Custom") != "custom-value" {
			t.Error("Custom header not set")
		}

		// Verify request body
		var body map[string]string
		json.NewDecoder(r.Body).Decode(&body)
		if body["test"] != "data" {
			t.Error("Request body incorrect")
		}

		// Return success response
		w.Header().Set("X-Request-Id", "req-123")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}))
	defer server.Close()

	client := NewClient(
		server.Client(),
		server.URL,
		"test-key",
		2,
		map[string]string{"X-Custom": "custom-value"},
		"test-agent",
		nil,
	)

	req := &Request{
		Method: "POST",
		Path:   "/test",
		Body:   map[string]string{"test": "data"},
	}

	resp, err := client.Do(context.Background(), req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	if resp.RequestID != "req-123" {
		t.Errorf("RequestID = %s, want req-123", resp.RequestID)
	}

	var result map[string]string
	json.Unmarshal(resp.Body, &result)
	if result["result"] != "success" {
		t.Error("Response body incorrect")
	}
}

func TestClient_Do_Retry(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			// Fail first 2 attempts
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Succeed on 3rd attempt
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", 3, nil, "test-agent", nil)

	req := &Request{
		Method: "GET",
		Path:   "/test",
	}

	resp, err := client.Do(context.Background(), req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
}

func TestClient_Do_MaxRetriesExceeded(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", 2, nil, "test-agent", nil)

	req := &Request{
		Method: "GET",
		Path:   "/test",
	}

	_, err := client.Do(context.Background(), req)
	if err == nil {
		t.Error("Expected error after max retries")
	}

	// Should try initial + 2 retries = 3 total
	if attempts != 3 {
		t.Errorf("Expected 3 attempts, got %d", attempts)
	}
}

func TestClient_Do_ContextCancellation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", 2, nil, "test-agent", nil)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	req := &Request{
		Method: "GET",
		Path:   "/test",
	}

	_, err := client.Do(ctx, req)
	if err == nil {
		t.Error("Expected context deadline exceeded error")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("Expected context.DeadlineExceeded, got %v", err)
	}
}

func TestClient_Do_NonRetryableError(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		w.WriteHeader(http.StatusBadRequest) // 400 is not retryable
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", 2, nil, "test-agent", nil)

	req := &Request{
		Method: "GET",
		Path:   "/test",
	}

	_, err := client.Do(context.Background(), req)
	if err == nil {
		t.Error("Expected error for 400 status")
	}

	// Should not retry 400 errors
	if attempts != 1 {
		t.Errorf("Expected 1 attempt, got %d", attempts)
	}
}

func TestClient_Do_RateLimitRetry(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts == 1 {
			w.WriteHeader(http.StatusTooManyRequests) // 429 is retryable
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", 2, nil, "test-agent", nil)

	req := &Request{
		Method: "GET",
		Path:   "/test",
	}

	resp, err := client.Do(context.Background(), req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}

	if attempts != 2 {
		t.Errorf("Expected 2 attempts, got %d", attempts)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}
}

func TestClient_DoStream_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify streaming headers
		if r.Header.Get("Accept") != "text/event-stream" {
			t.Error("Accept header not set for streaming")
		}
		if r.Header.Get("Cache-Control") != "no-cache" {
			t.Error("Cache-Control header not set")
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("X-Request-Id", "stream-123")
		w.WriteHeader(http.StatusOK)

		// Write some SSE data
		w.Write([]byte("data: test\n\n"))
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", 2, nil, "test-agent", nil)

	req := &Request{
		Method: "POST",
		Path:   "/stream",
		Body:   map[string]string{"test": "data"},
	}

	resp, err := client.DoStream(context.Background(), req)
	if err != nil {
		t.Fatalf("DoStream() error = %v", err)
	}
	defer resp.Response.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("StatusCode = %d, want %d", resp.StatusCode, http.StatusOK)
	}

	if resp.RequestID != "stream-123" {
		t.Errorf("RequestID = %s, want stream-123", resp.RequestID)
	}

	// Verify we can read from the stream
	data, err := io.ReadAll(resp.Response.Body)
	if err != nil {
		t.Fatalf("Failed to read stream: %v", err)
	}

	if !strings.Contains(string(data), "data: test") {
		t.Error("Stream data incorrect")
	}
}

func TestClient_DoStream_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", 2, nil, "test-agent", nil)

	req := &Request{
		Method: "POST",
		Path:   "/stream",
	}

	_, err := client.DoStream(context.Background(), req)
	if err == nil {
		t.Error("Expected error for 400 status")
	}

	if !strings.Contains(err.Error(), "400") {
		t.Errorf("Error should mention status code: %v", err)
	}
}

func TestClient_CalculateBackoff(t *testing.T) {
	client := NewClient(&http.Client{}, "https://api.example.com", "key", 3, nil, "agent", nil)

	tests := []struct {
		attempt  int
		minDelay time.Duration
		maxDelay time.Duration
	}{
		{1, 375 * time.Millisecond, 625 * time.Millisecond},   // 500ms ±25%
		{2, 750 * time.Millisecond, 1250 * time.Millisecond},  // 1s ±25%
		{3, 1500 * time.Millisecond, 2500 * time.Millisecond}, // 2s ±25%
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("attempt_%d", tt.attempt), func(t *testing.T) {
			delay := client.calculateBackoff(tt.attempt)
			if delay < tt.minDelay || delay > tt.maxDelay {
				t.Errorf("Backoff delay %v out of range [%v, %v]", delay, tt.minDelay, tt.maxDelay)
			}
		})
	}
}

func TestClient_CalculateBackoff_MaxDelay(t *testing.T) {
	client := NewClient(&http.Client{}, "https://api.example.com", "key", 10, nil, "agent", nil)

	// Test that backoff is capped at MaxRetryDelay
	delay := client.calculateBackoff(10)                       // Very high attempt number
	maxAllowed := time.Duration(float64(MaxRetryDelay) * 1.25) // Allow for jitter
	if delay > maxAllowed {
		t.Errorf("Backoff delay %v exceeds max %v", delay, maxAllowed)
	}
}

func TestClient_ShouldRetryStatus(t *testing.T) {
	client := NewClient(&http.Client{}, "https://api.example.com", "key", 3, nil, "agent", nil)

	retryableStatuses := []int{
		http.StatusRequestTimeout,      // 408
		http.StatusConflict,            // 409
		http.StatusTooManyRequests,     // 429
		http.StatusInternalServerError, // 500
		http.StatusBadGateway,          // 502
		http.StatusServiceUnavailable,  // 503
		http.StatusGatewayTimeout,      // 504
	}

	for _, status := range retryableStatuses {
		t.Run(fmt.Sprintf("status_%d", status), func(t *testing.T) {
			if !client.shouldRetryStatus(status) {
				t.Errorf("Status %d should be retryable", status)
			}
		})
	}

	nonRetryableStatuses := []int{
		http.StatusOK,                  // 200
		http.StatusBadRequest,          // 400
		http.StatusUnauthorized,        // 401
		http.StatusForbidden,           // 403
		http.StatusNotFound,            // 404
		http.StatusUnprocessableEntity, // 422
	}

	for _, status := range nonRetryableStatuses {
		t.Run(fmt.Sprintf("status_%d_not_retryable", status), func(t *testing.T) {
			if client.shouldRetryStatus(status) {
				t.Errorf("Status %d should not be retryable", status)
			}
		})
	}
}

func TestClient_RequestHeaders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify all expected headers
		headers := map[string]string{
			"Authorization":      "Bearer test-key",
			"Content-Type":       "application/json",
			"User-Agent":         "test-agent/1.0",
			"X-Custom":           "custom-value",
			"X-Request-Specific": "specific-value",
		}

		for key, expected := range headers {
			if got := r.Header.Get(key); got != expected {
				t.Errorf("Header %s = %s, want %s", key, got, expected)
			}
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := NewClient(
		server.Client(),
		server.URL,
		"test-key",
		0,
		map[string]string{"X-Custom": "custom-value"},
		"test-agent/1.0",
		nil,
	)

	req := &Request{
		Method: "POST",
		Path:   "/test",
		Headers: map[string]string{
			"X-Request-Specific": "specific-value",
		},
		Body: map[string]string{"test": "data"},
	}

	_, err := client.Do(context.Background(), req)
	if err != nil {
		t.Fatalf("Do() error = %v", err)
	}
}

func TestClient_ErrorResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Request-Id", "err-123")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Invalid request parameters",
		})
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", 0, nil, "test-agent", nil)

	req := &Request{
		Method: "POST",
		Path:   "/test",
	}

	_, err := client.Do(context.Background(), req)
	if err == nil {
		t.Fatal("Expected error for 400 status")
	}

	errMsg := err.Error()
	if !strings.Contains(errMsg, "Invalid request parameters") {
		t.Errorf("Error should contain message from response: %v", errMsg)
	}
	if !strings.Contains(errMsg, "400") {
		t.Errorf("Error should contain status code: %v", errMsg)
	}
	if !strings.Contains(errMsg, "err-123") {
		t.Errorf("Error should contain request ID: %v", errMsg)
	}
}
