# Perplexity SDK Python to Go Porting Plan

## Executive Summary

This document outlines a comprehensive plan for porting the Perplexity Python SDK (v0.20.0) to Go. The Python SDK is a Stainless-generated client library providing synchronous and asynchronous access to the Perplexity REST API with full type safety.

**Source:** `docs/perplexity-py-0.20.0/`  
**Target:** Go SDK with idiomatic Go patterns  
**Version:** 0.1.0 (initial port)

---

## 1. Architecture Overview

### 1.1 Python SDK Structure

```
perplexity/
├── _client.py               # Main clients
├── _base_client.py          # HTTP client
├── _streaming.py            # SSE streaming
├── _exceptions.py           # Error types
├── _types.py                # Type definitions
├── resources/
│   ├── chat/completions.py  # Chat API
│   ├── search.py            # Search API
│   └── async_/              # Async endpoints
└── types/                   # Response types
```

### 1.2 Proposed Go SDK Structure

```
perplexity-go/
├── client.go                # Main Client
├── config.go                # Configuration
├── errors.go                # Error types
├── streaming.go             # SSE streaming
├── chat/
│   ├── completions.go       # Chat service
│   └── types.go
├── search/
│   ├── search.go            # Search service
│   └── types.go
├── types/                   # Shared types
│   ├── message.go
│   ├── stream.go
│   └── params.go
└── internal/
    ├── http/                # HTTP layer
    └── sse/                 # SSE decoder
```

---

## 2. Core Components

### 2.1 Client Implementation

**Priority:** HIGH | **Complexity:** MEDIUM

```go
type Client struct {
    apiKey     string
    baseURL    string
    httpClient *http.Client
    
    Chat   *chat.Service
    Search *search.Service
    
    maxRetries int
    timeout    time.Duration
}

func NewClient(apiKey string, opts ...ClientOption) (*Client, error)

type ClientOption func(*Client) error
```

**Key Features:**
- Functional options pattern
- Context-based cancellation
- Service-based API organization
- No separate async client (use goroutines)

### 2.2 HTTP Client Layer

**Priority:** HIGH | **Complexity:** MEDIUM

**Features:**
- Request building with JSON marshaling
- Response parsing and validation
- Retry logic (exponential backoff, 2 retries default)
- Timeout management (15min default)
- Error mapping to typed errors

```go
type HTTPClient struct {
    client     *http.Client
    baseURL    string
    apiKey     string
    maxRetries int
}

func (c *HTTPClient) Do(ctx context.Context, req *Request) (*Response, error)
```

### 2.3 Error Handling

**Priority:** HIGH | **Complexity:** LOW

```go
type Error struct {
    Message    string
    StatusCode int
    Body       json.RawMessage
    RequestID  string
}

// Specific errors
type BadRequestError struct{ *Error }          // 400
type AuthenticationError struct{ *Error }      // 401
type PermissionDeniedError struct{ *Error }    // 403
type NotFoundError struct{ *Error }            // 404
type RateLimitError struct{ *Error }           // 429
type InternalServerError struct{ *Error }      // 5xx
type ConnectionError struct{ *Error }
type TimeoutError struct{ *Error }
```

### 2.4 Type System

**Priority:** HIGH | **Complexity:** HIGH

**Challenges:**
- Optional fields → Use pointers
- Union types → Interface with type assertions
- Nested structures → Proper composition

```go
type ChatMessage struct {
    Role           Role                `json:"role"`
    Content        MessageContent      `json:"content"`
    ReasoningSteps []ReasoningStep     `json:"reasoning_steps,omitempty"`
    ToolCalls      []ToolCall          `json:"tool_calls,omitempty"`
}

type StreamChunk struct {
    ID            string         `json:"id"`
    Choices       []Choice       `json:"choices"`
    Created       int64          `json:"created"`
    Model         string         `json:"model"`
    Citations     []string       `json:"citations,omitempty"`
    SearchResults []SearchResult `json:"search_results,omitempty"`
    Usage         *UsageInfo     `json:"usage,omitempty"`
}
```

### 2.5 Streaming Support (SSE)

**Priority:** HIGH | **Complexity:** HIGH

```go
type Stream[T any] struct {
    reader   *bufio.Reader
    response *http.Response
    decoder  *sse.Decoder
    ctx      context.Context
}

func (s *Stream[T]) Next() (T, error)
func (s *Stream[T]) Close() error
```

**Usage:**
```go
stream, err := client.Chat.Completions.CreateStream(ctx, params)
defer stream.Close()

for {
    chunk, err := stream.Next()
    if err == io.EOF {
        break
    }
    // Process chunk
}
```

---

## 3. API Endpoints

### 3.1 Chat Completions

**Priority:** HIGH | **Complexity:** HIGH

**Endpoint:** `POST /chat/completions`

```go
type CompletionParams struct {
    Messages             []types.ChatMessage   `json:"messages"`
    Model                string                `json:"model"`
    Stream               *bool                 `json:"stream,omitempty"`
    MaxTokens            *int                  `json:"max_tokens,omitempty"`
    Temperature          *float64              `json:"temperature,omitempty"`
    TopP                 *float64              `json:"top_p,omitempty"`
    Tools                []Tool                `json:"tools,omitempty"`
    WebSearchOptions     *WebSearchOptions     `json:"web_search_options,omitempty"`
    SearchDomainFilter   []string              `json:"search_domain_filter,omitempty"`
    SearchRecencyFilter  *SearchRecencyFilter  `json:"search_recency_filter,omitempty"`
    ReturnImages         *bool                 `json:"return_images,omitempty"`
    // ... 50+ more optional fields
}

func (s *Service) Create(ctx context.Context, params *CompletionParams) (*types.StreamChunk, error)
func (s *Service) CreateStream(ctx context.Context, params *CompletionParams) (*Stream[types.StreamChunk], error)
```

### 3.2 Search API

**Priority:** MEDIUM | **Complexity:** LOW

**Endpoint:** `POST /search`

```go
type SearchParams struct {
    Query                 SearchQuery          `json:"query"`
    MaxResults            *int                 `json:"max_results,omitempty"`
    SearchMode            *SearchMode          `json:"search_mode,omitempty"`
    SearchRecencyFilter   *SearchRecencyFilter `json:"search_recency_filter,omitempty"`
    SearchDomainFilter    []string             `json:"search_domain_filter,omitempty"`
}

type SearchResponse struct {
    ID      string         `json:"id"`
    Results []SearchResult `json:"results"`
}

func (s *Service) Create(ctx context.Context, params *SearchParams) (*SearchResponse, error)
```

### 3.3 Async Chat Completions

**Priority:** LOW | **Complexity:** MEDIUM

**Endpoints:**
- `POST /async/chat/completions` - Create
- `GET /async/chat/completions` - List
- `GET /async/chat/completions/{id}` - Get status

---

## 4. Implementation Phases

### Phase 1: Foundation (Week 1-2)
- [ ] Project setup
- [ ] Client struct and configuration
- [ ] HTTP client wrapper
- [ ] Error types
- [ ] Basic request/response handling

### Phase 2: Type System (Week 2-3)
- [ ] Core types (ChatMessage, Choice, UsageInfo)
- [ ] Stream types
- [ ] Search types
- [ ] Parameter types
- [ ] JSON marshaling tests

### Phase 3: Chat Completions (Week 3-4)
- [ ] Chat service
- [ ] Create() method
- [ ] Parameter validation
- [ ] Response parsing
- [ ] Tests and examples

### Phase 4: Streaming (Week 4-5)
- [ ] SSE decoder
- [ ] Stream type
- [ ] CreateStream() method
- [ ] Context cancellation
- [ ] Tests and examples

### Phase 5: Search API (Week 5)
- [ ] Search service
- [ ] Create() method
- [ ] Tests and examples

### Phase 6: Advanced Features (Week 6)
- [ ] Retry logic
- [ ] Timeout configuration
- [ ] Custom headers
- [ ] Request logging

### Phase 7: Documentation (Week 7)
- [ ] README
- [ ] GoDoc comments
- [ ] Examples
- [ ] API reference

### Phase 8: Release (Week 8)
- [ ] Code review
- [ ] Performance testing
- [ ] CI/CD setup
- [ ] v0.1.0 release

---

## 5. Testing Strategy

### 5.1 Unit Tests (80%+ coverage)

- Client initialization
- HTTP client (request/response/retry)
- Streaming (SSE decoding, iteration)
- Type marshaling/unmarshaling
- Error handling

### 5.2 Integration Tests

Use `httptest` for mock server:

```go
server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(mockResponse)
}))
defer server.Close()

client, _ := perplexity.NewClient("test-key", 
    perplexity.WithBaseURL(server.URL))
```

### 5.3 Examples

1. Basic chat completion
2. Streaming chat
3. Function calling
4. Web search
5. Search API
6. Error handling
7. Custom HTTP client
8. Context cancellation

---

## 6. Dependencies

**Core:** Standard library only (no external dependencies)

**Testing:**
```
github.com/stretchr/testify v1.8.4
```

---

## 7. Design Decisions

### 7.1 No Separate Async Client
Use goroutines instead of separate async client (idiomatic Go)

### 7.2 Functional Options Pattern
```go
client, err := perplexity.NewClient(
    apiKey,
    perplexity.WithTimeout(30*time.Second),
    perplexity.WithMaxRetries(5),
)
```

### 7.3 Pointers for Optional Fields
Distinguishes zero value from not set

### 7.4 Context Throughout
All API methods accept `context.Context` for cancellation

### 7.5 Generic Streams
`Stream[T]` for type-safe streaming

---

## 8. API Compatibility Matrix

| Feature | Python | Go | Priority |
|---------|--------|-----|----------|
| Chat Completions | ✅ | 🔄 | HIGH |
| Streaming | ✅ | 🔄 | HIGH |
| Search API | ✅ | 🔄 | MEDIUM |
| Function Calling | ✅ | 🔄 | HIGH |
| Web Search Options | ✅ | 🔄 | HIGH |
| Retry Logic | ✅ | 🔄 | HIGH |
| Error Types | ✅ | 🔄 | HIGH |
| Async Completions | ✅ | 🔄 | LOW |

---

## 9. Documentation Requirements

- README with quick start
- GoDoc for all exports
- Comprehensive examples
- API reference
- Migration guide from Python
- CONTRIBUTING guide
- CHANGELOG

---

## 10. Success Criteria

- [ ] Feature parity with Python SDK (core features)
- [ ] 80%+ test coverage
- [ ] Zero external dependencies (core)
- [ ] Idiomatic Go code
- [ ] Complete documentation
- [ ] Working examples
- [ ] CI/CD pipeline
- [ ] v0.1.0 release

---

## Appendix A: Key Python SDK Files

**Core Files:**
- `_client.py` (426 lines) - Main client
- `_base_client.py` (1996 lines) - HTTP client
- `_streaming.py` (370 lines) - SSE streaming
- `_exceptions.py` (109 lines) - Errors
- `resources/chat/completions.py` (891 lines) - Chat API
- `resources/search.py` (221 lines) - Search API

**Type Files:**
- `types/stream_chunk.py` - StreamChunk
- `types/search_create_response.py` - Search response
- `types/shared/chat_message_input.py` (205 lines) - Message types
- `types/shared/choice.py` - Choice type
- `types/shared/usage_info.py` - Usage info

**Total LOC:** ~4,000+ lines (excluding tests)

---

## Appendix B: Example Usage Comparison

### Python
```python
from perplexity import Perplexity

client = Perplexity(api_key="...")

# Non-streaming
response = client.chat.completions.create(
    messages=[{"role": "user", "content": "Hello"}],
    model="sonar"
)

# Streaming
stream = client.chat.completions.create(
    messages=[{"role": "user", "content": "Hello"}],
    model="sonar",
    stream=True
)
for chunk in stream:
    print(chunk.choices[0].delta.content)
```

### Go (Proposed)
```go
import "github.com/perplexityai/perplexity-go"

client, _ := perplexity.NewClient("...")

// Non-streaming
response, err := client.Chat.Completions.Create(ctx, &chat.CompletionParams{
    Messages: []types.ChatMessage{
        {Role: types.RoleUser, Content: types.TextContent("Hello")},
    },
    Model: "sonar",
})

// Streaming
stream, err := client.Chat.Completions.CreateStream(ctx, &chat.CompletionParams{
    Messages: []types.ChatMessage{
        {Role: types.RoleUser, Content: types.TextContent("Hello")},
    },
    Model: "sonar",
})
defer stream.Close()

for {
    chunk, err := stream.Next()
    if err == io.EOF {
        break
    }
    fmt.Println(chunk.Choices[0].Delta.Content)
}
```
