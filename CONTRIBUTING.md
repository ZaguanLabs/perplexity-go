# Contributing to Perplexity Go SDK

Thank you for your interest in contributing to the Perplexity Go SDK! This document provides guidelines and instructions for contributing.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [Development Workflow](#development-workflow)
- [Testing](#testing)
- [Code Style](#code-style)
- [Submitting Changes](#submitting-changes)
- [Reporting Issues](#reporting-issues)

## Code of Conduct

This project follows a standard code of conduct. Please be respectful and constructive in all interactions.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/perplexity-go.git
   cd perplexity-go
   ```
3. **Add upstream remote**:
   ```bash
   git remote add upstream https://github.com/perplexityai/perplexity-go.git
   ```

## Development Setup

### Prerequisites

- **Go 1.21 or higher**
- **Git**
- A Perplexity API key for integration testing (optional)

### Install Dependencies

The SDK has zero external dependencies for core functionality. To run tests:

```bash
go test ./...
```

### Environment Variables

For running examples and integration tests:

```bash
export PERPLEXITY_API_KEY="your-api-key"
```

## Project Structure

```
perplexity-go/
├── perplexity/              # Main package
│   ├── client.go            # Client implementation
│   ├── config.go            # Configuration options
│   ├── errors.go            # Error types
│   ├── chat/                # Chat completions service
│   │   ├── chat.go          # Service implementation
│   │   ├── params.go        # Parameter types
│   │   ├── stream.go        # Streaming support
│   │   └── *_test.go        # Tests
│   ├── types/               # Type definitions
│   │   ├── message.go       # Message types
│   │   ├── choice.go        # Choice types
│   │   ├── usage.go         # Usage types
│   │   └── ...
│   └── internal/            # Internal packages
│       ├── http/            # HTTP client wrapper
│       └── sse/             # SSE decoder
├── examples/                # Example code
│   ├── basic/
│   ├── chat/
│   └── streaming/
├── docs/                    # Documentation
└── README.md
```

## Development Workflow

### 1. Create a Branch

```bash
git checkout -b feature/your-feature-name
```

Use descriptive branch names:
- `feature/add-search-api` - New features
- `fix/streaming-error` - Bug fixes
- `docs/update-readme` - Documentation updates
- `refactor/client-structure` - Code refactoring

### 2. Make Changes

- Write clear, idiomatic Go code
- Follow the existing code style
- Add tests for new functionality
- Update documentation as needed

### 3. Commit Changes

Write clear, descriptive commit messages:

```bash
git commit -m "feat: add search API support

- Implement SearchService with Create method
- Add SearchParams and SearchResponse types
- Include comprehensive tests
- Add example usage"
```

Commit message format:
- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `test:` - Test additions/changes
- `refactor:` - Code refactoring
- `chore:` - Maintenance tasks

### 4. Keep Your Branch Updated

```bash
git fetch upstream
git rebase upstream/main
```

## Testing

### Running Tests

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run tests verbosely:
```bash
go test -v ./...
```

Run specific package tests:
```bash
go test ./perplexity/chat/...
```

### Writing Tests

- **Unit tests**: Test individual functions and methods
- **Integration tests**: Test with mock HTTP servers
- **Table-driven tests**: Use for multiple test cases

Example test structure:

```go
func TestFeature(t *testing.T) {
    tests := []struct {
        name    string
        input   string
        want    string
        wantErr bool
    }{
        {
            name:  "valid input",
            input: "test",
            want:  "result",
        },
        // More test cases...
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := Feature(tt.input)
            if (err != nil) != tt.wantErr {
                t.Errorf("Feature() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if got != tt.want {
                t.Errorf("Feature() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

### Test Coverage Goals

- Aim for **80%+ coverage** for new code
- All exported functions should have tests
- Test both success and error cases
- Test edge cases and boundary conditions

## Code Style

### Go Standards

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Use `golint` for linting
- Follow Go naming conventions

### Formatting

Format your code before committing:

```bash
go fmt ./...
```

### Linting

Run linters:

```bash
go vet ./...
```

### Documentation

- Add GoDoc comments for all exported types, functions, and methods
- Use complete sentences in comments
- Provide usage examples in documentation

Example:

```go
// Create generates a chat completion for the given parameters.
// It returns a StreamChunk containing the completion response.
//
// Example:
//
//	result, err := client.Chat.Create(ctx, &chat.CompletionParams{
//	    Model: "sonar",
//	    Messages: []types.ChatMessage{
//	        types.UserMessage("Hello!"),
//	    },
//	})
func (s *Service) Create(ctx context.Context, params *CompletionParams) (*types.StreamChunk, error) {
    // Implementation...
}
```

### Error Handling

- Return errors, don't panic
- Wrap errors with context: `fmt.Errorf("failed to do X: %w", err)`
- Use custom error types for specific errors
- Check for nil before dereferencing pointers

### Naming Conventions

- **Packages**: Short, lowercase, no underscores (e.g., `chat`, `types`)
- **Files**: Lowercase with underscores (e.g., `chat_test.go`)
- **Types**: PascalCase (e.g., `CompletionParams`)
- **Functions**: PascalCase for exported, camelCase for unexported
- **Variables**: camelCase
- **Constants**: PascalCase or SCREAMING_SNAKE_CASE for enums

## Submitting Changes

### 1. Push Your Branch

```bash
git push origin feature/your-feature-name
```

### 2. Create a Pull Request

- Go to the repository on GitHub
- Click "New Pull Request"
- Select your branch
- Fill out the PR template with:
  - Clear description of changes
  - Related issue numbers
  - Testing performed
  - Breaking changes (if any)

### 3. PR Review Process

- Maintainers will review your PR
- Address any feedback or requested changes
- Keep the PR updated with main branch
- Once approved, a maintainer will merge

### PR Checklist

- [ ] Tests pass locally (`go test ./...`)
- [ ] Code is formatted (`go fmt ./...`)
- [ ] No linting errors (`go vet ./...`)
- [ ] Documentation updated
- [ ] Examples added/updated (if applicable)
- [ ] CHANGELOG updated (for significant changes)
- [ ] Commit messages follow convention

## Reporting Issues

### Bug Reports

Include:
- Go version (`go version`)
- SDK version
- Minimal code to reproduce
- Expected behavior
- Actual behavior
- Error messages/stack traces

### Feature Requests

Include:
- Use case description
- Proposed API/interface
- Examples of how it would be used
- Alternatives considered

### Security Issues

**Do not** open public issues for security vulnerabilities. Instead:
- Email security concerns to the maintainers
- Provide detailed description
- Allow time for a fix before public disclosure

## Development Tips

### Running Examples

```bash
# Set API key
export PERPLEXITY_API_KEY="your-key"

# Run example
go run examples/chat/main.go
```

### Building

```bash
# Build all packages
go build ./...

# Build specific example
go build -o bin/chat examples/chat/main.go
```

### Debugging

Use the `-v` flag for verbose test output:
```bash
go test -v ./perplexity/chat/...
```

### Benchmarking

Write benchmarks for performance-critical code:

```go
func BenchmarkFeature(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Feature()
    }
}
```

Run benchmarks:
```bash
go test -bench=. ./...
```

## Questions?

- Check existing issues and PRs
- Read the documentation
- Ask in discussions or issues

## License

By contributing, you agree that your contributions will be licensed under the Apache 2.0 License.

---

Thank you for contributing to the Perplexity Go SDK! 🎉
