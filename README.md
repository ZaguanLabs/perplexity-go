# Perplexity Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/perplexityai/perplexity-go.svg)](https://pkg.go.dev/github.com/perplexityai/perplexity-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/perplexityai/perplexity-go)](https://goreportcard.com/report/github.com/perplexityai/perplexity-go)

The official Go client library for the [Perplexity API](https://docs.perplexity.ai/).

> **⚠️ Work in Progress**: This SDK is currently under active development. The API may change before the v1.0.0 release.

## Features

- ✅ **Type-safe**: Full type definitions for all request and response types
- ✅ **Context-aware**: All methods accept `context.Context` for cancellation and timeouts
- ✅ **Retry logic**: Automatic exponential backoff for transient errors
- ✅ **Zero dependencies**: Uses only the Go standard library
- 🚧 **Streaming support**: Server-Sent Events (SSE) for real-time responses (coming soon)
- 🚧 **Chat completions**: Full support for chat API (coming soon)
- 🚧 **Search API**: Web search capabilities (coming soon)

## Installation

```bash
go get github.com/perplexityai/perplexity-go
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/perplexityai/perplexity-go"
)

func main() {
    // Create a new client
    client, err := perplexity.NewClient("your-api-key")
    if err != nil {
        log.Fatal(err)
    }

    // API methods will be available in upcoming releases
    // Example (coming soon):
    // resp, err := client.Chat.Completions.Create(ctx, &chat.CompletionParams{
    //     Model: "sonar",
    //     Messages: []types.ChatMessage{
    //         {Role: types.RoleUser, Content: types.TextContent("Hello!")},
    //     },
    // })
}
```

## Configuration

### API Key

The API key can be provided in two ways:

1. **Direct parameter**:
```go
client, err := perplexity.NewClient("your-api-key")
```

2. **Environment variable**:
```bash
export PERPLEXITY_API_KEY="your-api-key"
```
```go
client, err := perplexity.NewClient("") // Reads from PERPLEXITY_API_KEY
```

### Client Options

Customize the client with functional options:

```go
client, err := perplexity.NewClient(
    "your-api-key",
    perplexity.WithBaseURL("https://custom.api.com"),
    perplexity.WithTimeout(30*time.Second),
    perplexity.WithMaxRetries(5),
    perplexity.WithDefaultHeader("X-Custom-Header", "value"),
)
```

Available options:
- `WithBaseURL(url string)` - Set a custom API base URL
- `WithHTTPClient(client *http.Client)` - Use a custom HTTP client
- `WithTimeout(timeout time.Duration)` - Set request timeout (default: 15 minutes)
- `WithMaxRetries(retries int)` - Set maximum retry attempts (default: 2)
- `WithDefaultHeader(key, value string)` - Add a default header to all requests

## Error Handling

The SDK provides typed errors for different HTTP status codes:

```go
resp, err := client.Chat.Completions.Create(ctx, params)
if err != nil {
    switch e := err.(type) {
    case *perplexity.AuthenticationError:
        log.Fatal("Invalid API key:", e)
    case *perplexity.RateLimitError:
        log.Println("Rate limited, retrying...")
    case *perplexity.InternalServerError:
        log.Println("Server error, will retry automatically")
    default:
        log.Fatal("Unexpected error:", err)
    }
}
```

Error types:
- `BadRequestError` (400)
- `AuthenticationError` (401)
- `PermissionDeniedError` (403)
- `NotFoundError` (404)
- `ConflictError` (409)
- `UnprocessableEntityError` (422)
- `RateLimitError` (429)
- `InternalServerError` (5xx)
- `ConnectionError` (network errors)
- `TimeoutError` (request timeout)

Helper functions:
- `IsRetryable(err error) bool` - Check if error is retryable
- `IsRateLimitError(err error) bool` - Check for rate limit errors
- `IsAuthenticationError(err error) bool` - Check for auth errors
- `IsTimeoutError(err error) bool` - Check for timeout errors

## Development Status

### Phase 1: Foundation ✅ (Completed)
- [x] Project setup
- [x] Client configuration
- [x] Error types
- [x] HTTP client wrapper
- [x] Retry logic
- [x] Unit tests

### Phase 2: Type System ✅ (Completed)
- [x] Core types (ChatMessage, Choice, UsageInfo)
- [x] Stream types (StreamChunk)
- [x] Search types (SearchResult, SearchResponse)
- [x] Tool types (Tool, ToolCall)
- [x] Reasoning types (ReasoningStep)
- [x] Helper functions
- [x] Comprehensive tests (20+ test cases)

### Phase 3: Chat Completions 📋 (Next)
- [ ] Chat service
- [ ] Create method
- [ ] Parameter validation

### Phase 4: Streaming 📋 (Planned)
- [ ] SSE decoder
- [ ] Stream type
- [ ] Streaming support

### Phase 5: Search API 📋 (Planned)
- [ ] Search service
- [ ] Search methods

## Requirements

- Go 1.21 or higher

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## License

Apache 2.0 - See [LICENSE](LICENSE) for details.

## Links

- [API Documentation](https://docs.perplexity.ai/)
- [Python SDK](https://github.com/perplexityai/perplexity-py)
- [Issue Tracker](https://github.com/perplexityai/perplexity-go/issues)
