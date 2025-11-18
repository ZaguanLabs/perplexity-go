# Perplexity Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/perplexityai/perplexity-go.svg)](https://pkg.go.dev/github.com/perplexityai/perplexity-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/perplexityai/perplexity-go)](https://goreportcard.com/report/github.com/perplexityai/perplexity-go)

The **unofficial** Go client library for the [Perplexity API](https://docs.perplexity.ai/).

> **Version 0.1.0** - Initial release with full chat completions, streaming, and search support.

## Features

- ✅ **Type-safe**: Full type definitions for all request and response types
- ✅ **Context-aware**: All methods accept `context.Context` for cancellation and timeouts
- ✅ **Retry logic**: Automatic exponential backoff for transient errors
- ✅ **Zero dependencies**: Uses only the Go standard library
- ✅ **Chat completions**: Full support for chat API with 60+ parameters
- ✅ **Streaming support**: Server-Sent Events (SSE) for real-time responses
- ✅ **Search API**: Web search with filtering and multiple query support

## Installation

```bash
go get github.com/perplexityai/perplexity-go/perplexity
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/perplexityai/perplexity-go/perplexity"
    "github.com/perplexityai/perplexity-go/perplexity/chat"
    "github.com/perplexityai/perplexity-go/perplexity/search"
    "github.com/perplexityai/perplexity-go/perplexity/types"
)

func main() {
    // Create a new client
    client, err := perplexity.NewClient("your-api-key")
    if err != nil {
        log.Fatal(err)
    }

    // Create a chat completion
    result, err := client.Chat.Create(context.Background(), &chat.CompletionParams{
        Model: "sonar",
        Messages: []types.ChatMessage{
            types.UserMessage("What is the capital of France?"),
        },
        MaxTokens: types.Int(100),
    })
    if err != nil {
        log.Fatal(err)
    }

    // Print the response
    fmt.Println(result.Choices[0].Message.Content)

    // Perform a web search
    searchResult, err := client.Search.Create(context.Background(), &search.SearchParams{
        Query:      "latest AI developments",
        MaxResults: types.Int(5),
    })
    if err != nil {
        log.Fatal(err)
    }

    // Print search results
    for _, item := range searchResult.Results {
        fmt.Printf("%s: %s\n", item.Title, item.URL)
    }
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

### Phase 3: Chat Completions ✅ (Completed)
- [x] Chat service implementation
- [x] CompletionParams with 60+ parameters
- [x] Create() method for non-streaming completions
- [x] Parameter validation
- [x] Comprehensive tests
- [x] Working examples

### Phase 4: Streaming ✅ (Completed)
- [x] SSE (Server-Sent Events) decoder
- [x] Stream type with Next() and Iter() methods
- [x] CreateStream() method
- [x] Context cancellation support
- [x] Error event handling
- [x] Comprehensive streaming tests
- [x] Streaming examples

### Phase 5: Search API ✅ (Completed)
- [x] Search service implementation
- [x] SearchParams with all filter options
- [x] Create() method for web search
- [x] Support for single and multiple queries
- [x] Search mode support (web, academic, SEC)
- [x] Domain and language filtering
- [x] Recency filtering
- [x] Comprehensive tests
- [x] Search examples

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
