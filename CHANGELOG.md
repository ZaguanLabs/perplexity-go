# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2025-01-18

### Added

#### Core Features
- Initial release of the Perplexity Go SDK
- Complete implementation of Perplexity API in Go
- Zero external dependencies (uses only Go standard library)

#### Chat Completions
- `Chat.Create()` method for synchronous chat completions
- `Chat.CreateStream()` method for streaming responses
- Support for 60+ parameters including:
  - Temperature, top_p, frequency/presence penalties
  - Tools and function calling
  - Web search options and filters
  - Response format control (text, JSON schema, regex)
  - Reasoning effort control
  - Image and file attachments
  - Location-based search

#### Streaming Support
- Server-Sent Events (SSE) decoder
- `Stream.Next()` method for sequential iteration
- `Stream.Iter()` method for channel-based iteration
- Context cancellation support
- Error event handling
- Automatic resource cleanup

#### Search API
- `Search.Create()` method for web search
- Support for single and multiple queries
- Search modes: web, academic, SEC
- Advanced filtering:
  - Domain filtering
  - Language filtering
  - Recency filtering (hour, day, week, month, year)
  - Date range filtering
  - Country localization

#### Type System
- Complete type definitions matching Python SDK
- `ChatMessage` with text and structured content support
- `StreamChunk` for streaming responses
- `SearchResponse` for search results
- `Tool` and `ToolCall` for function calling
- `ReasoningStep` for reasoning traces
- Helper functions for creating messages and pointers

#### Error Handling
- Comprehensive error hierarchy
- Typed errors for all HTTP status codes:
  - `BadRequestError` (400)
  - `AuthenticationError` (401)
  - `PermissionDeniedError` (403)
  - `NotFoundError` (404)
  - `ConflictError` (409)
  - `UnprocessableEntityError` (422)
  - `RateLimitError` (429)
  - `InternalServerError` (5xx)
  - `ConnectionError`, `TimeoutError`
- Error helper functions (`IsRetryable`, `IsRateLimitError`, etc.)

#### Client Features
- Functional options pattern for configuration
- Automatic retry with exponential backoff
- Configurable timeouts and max retries
- Custom HTTP client support
- Default headers support
- Environment variable support for API key
- Version information (`client.Version()`)

#### Documentation
- Comprehensive README with examples
- CONTRIBUTING.md with development guidelines
- Inline GoDoc comments throughout
- 4 complete working examples:
  - Basic usage
  - Chat completions
  - Streaming responses
  - Search API

#### Testing
- 50+ test cases across all packages
- ~70% code coverage
- Mock HTTP servers for integration tests
- Table-driven tests
- All tests passing

### Project Structure
- Clean package organization
- Internal packages for implementation details
- Examples directory with working code
- Proper separation of concerns

### Dependencies
- **Zero external dependencies**
- Uses only Go 1.21+ standard library

---

## Release Notes

### v0.1.0 - Initial Release

This is the first release of the Perplexity Go SDK, providing complete API coverage for the Perplexity API with idiomatic Go patterns.

**Highlights:**
- 🚀 Full feature parity with Python SDK v0.20.0
- 📦 Zero dependencies
- ✅ Comprehensive test coverage
- 📚 Complete documentation
- 🔄 Streaming support with SSE
- 🔍 Full search API support
- 🛠️ Type-safe with generics
- ⚡ Production-ready

**Installation:**
```bash
go get github.com/ZaguanLabs/perplexity-go/perplexity
```

**Quick Start:**
```go
client, _ := perplexity.NewClient("your-api-key")

result, _ := client.Chat.Create(ctx, &chat.CompletionParams{
    Model: "sonar",
    Messages: []types.ChatMessage{
        types.UserMessage("Hello!"),
    },
})
```

See [README.md](README.md) for full documentation.

---

[0.1.0]: https://github.com/ZaguanLabs/perplexity-go/releases/tag/v0.1.0
