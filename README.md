# Perplexity Go SDK (Unofficial)

[![Go Reference](https://pkg.go.dev/badge/github.com/ZaguanLabs/perplexity-go.svg)](https://pkg.go.dev/github.com/ZaguanLabs/perplexity-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/ZaguanLabs/perplexity-go)](https://goreportcard.com/report/github.com/ZaguanLabs/perplexity-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

An **unofficial**, community-maintained Go client library for the [Perplexity API](https://docs.perplexity.ai/). This SDK provides idiomatic Go interfaces for chat completions, streaming responses, and web search.

> ⚠️ **Disclaimer**: This is an unofficial SDK and is not affiliated with, endorsed by, or supported by Perplexity AI. For official support, please refer to the [Perplexity API documentation](https://docs.perplexity.ai/).

## Why Use This SDK?

- 🚀 **Production Ready**: Comprehensive error handling, retries, and timeouts
- 📦 **Zero Dependencies**: Uses only the Go standard library
- 🔒 **Type Safe**: Full type definitions with compile-time safety
- ⚡ **Streaming Support**: Real-time responses with Server-Sent Events
- 🔍 **Complete API Coverage**: Chat, streaming, and search endpoints
- 📚 **Well Documented**: Extensive examples and GoDoc comments
- ✅ **Thoroughly Tested**: 70%+ test coverage with 50+ test cases

## Features

### Core Capabilities
- **Chat Completions**: Full support for Perplexity's chat API with 60+ parameters
- **Streaming Responses**: Real-time streaming with Server-Sent Events (SSE)
- **Web Search**: Advanced search with filtering, multiple queries, and specialized modes
- **Tool Calling**: Function calling and tool integration
- **Reasoning Traces**: Access to model reasoning steps

### Developer Experience
- **Context-Aware**: All methods accept `context.Context` for cancellation and timeouts
- **Automatic Retries**: Exponential backoff for transient errors
- **Type Safety**: Comprehensive type definitions with generics
- **Error Handling**: Detailed error types for all API responses
- **Flexible Configuration**: Functional options pattern for client setup

## Installation

```bash
go get github.com/ZaguanLabs/perplexity-go/perplexity
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ZaguanLabs/perplexity-go/perplexity"
    "github.com/ZaguanLabs/perplexity-go/perplexity/chat"
    "github.com/ZaguanLabs/perplexity-go/perplexity/search"
    "github.com/ZaguanLabs/perplexity-go/perplexity/types"
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

## Requirements

- Go 1.21 or higher
- A Perplexity API key ([get one here](https://www.perplexity.ai/settings/api))

## Project Resources

- 📖 [CHANGELOG.md](CHANGELOG.md) - Version history and release notes
- 🛠️ [DEVELOPMENT.md](docs/DEVELOPMENT.md) - Development status and roadmap
- 🤝 [CONTRIBUTING.md](CONTRIBUTING.md) - Contribution guidelines
- 📄 [LICENSE](LICENSE) - Apache 2.0 License

## Related Links

- [Perplexity API Documentation](https://docs.perplexity.ai/)
- [Official Python SDK](https://github.com/perplexityai/perplexity-py)
- [Issue Tracker](https://github.com/ZaguanLabs/perplexity-go/issues)

## Support

This is an unofficial, community-maintained project. For issues with this SDK:
- 🐛 [Report bugs](https://github.com/ZaguanLabs/perplexity-go/issues)
- 💡 [Request features](https://github.com/ZaguanLabs/perplexity-go/issues)
- 🤝 [Contribute](CONTRIBUTING.md)

For Perplexity API support, please contact [Perplexity AI](https://www.perplexity.ai/) directly.

## License

Apache 2.0 - See [LICENSE](LICENSE) for details.

## Acknowledgments

This SDK was built with reference to the official [Python SDK](https://github.com/perplexityai/perplexity-py) to ensure API compatibility and feature parity.
