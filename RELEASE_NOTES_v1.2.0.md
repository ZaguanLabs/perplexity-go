# Perplexity Go SDK v1.2.0

**Release Date:** May 2, 2026  
**Status:** Stable Release

---

## Release Summary

`v1.2.0` is a Python SDK parity infrastructure release for the Perplexity Go SDK.

This version focuses on matching the official Python SDK's client ergonomics and transport behavior more closely. It adds per-request options, default query parameters, client cloning, and raw response access across the public resource surface, while also closing a confirmed search request schema drift.

If you want the short version: this release makes the Go SDK easier to customize per request, easier to fork into configured clients, and better aligned with the Python SDK's raw response and request option patterns.

---

## Highlights

### Python-style request options

Resource methods now accept optional per-request options without breaking existing call sites. These options support:

- request headers
- query parameters
- extra JSON body fields
- per-request timeouts

### Default query parameters

Clients can now be configured with default query parameters that are automatically applied to outgoing requests.

### Client cloning

`Client.Copy()` and `Client.WithOptions()` make it easy to derive a new client from an existing client while overriding selected options such as base URL, headers, query parameters, timeout, retries, or transport.

### Raw response access

Most resource methods now have raw response variants that return typed SDK data together with HTTP metadata such as status code, headers, response body, and request ID.

### Search schema parity fix

`search.SearchParams` now includes the Python SDK's missing `last_updated_after_filter` and `last_updated_before_filter` fields.

---

## What’s New

### Added

- New `perplexity/api` package with reusable request option helpers:
  - `api.WithHeader()`
  - `api.WithHeaders()`
  - `api.WithQuery()`
  - `api.WithQueryParams()`
  - `api.WithExtraBody()`
  - `api.WithExtraBodyFields()`
  - `api.WithTimeout()`
- New raw response envelope type:
  - `api.RawResponse[T]`
- Default query client options:
  - `WithDefaultQuery()`
  - `WithDefaultQueryParams()`
- Bulk default header client option:
  - `WithDefaultHeaders()`
- Client cloning helpers:
  - `Client.Copy()`
  - `Client.WithOptions()`
- Raw response APIs for:
  - `Chat.CreateRaw()`
  - `Search.CreateRaw()`
  - `Responses.CreateRaw()`
  - `Embeddings.CreateRaw()`
  - `ContextualizedEmbeddings.CreateRaw()`
  - `AsyncChat.CreateRaw()`
  - `AsyncChat.ListRaw()`
  - `AsyncChat.GetRaw()`
  - `Browser.Sessions.CreateRaw()`
  - `Browser.Sessions.DeleteRaw()`
- Per-request options on resource methods across chat, search, responses, embeddings, contextualized embeddings, async chat, and browser sessions.
- Tests covering default query behavior, per-request options, extra body merging, and client copy behavior.

### Changed

- Updated SDK version to `v1.2.0`.
- Extended the internal HTTP client to merge:
  - default query parameters
  - request-level query parameters
  - per-request option query parameters
  - extra JSON body fields
- Extended streaming requests to support the same per-request headers, query parameters, extra body fields, and timeout options as non-streaming requests.
- Preserved source compatibility by making request options variadic on existing resource methods.

### Fixed

- Added missing `search.SearchParams.LastUpdatedAfterFilter` and `search.SearchParams.LastUpdatedBeforeFilter` fields for parity with Python SDK `SearchCreateParams`.
- Added JSON round-trip test coverage for the new search filter fields.

---

## Upgrade Notes

### Install or update

```bash
go get github.com/ZaguanLabs/perplexity-go/perplexity@v1.2.0
```

### Existing resource calls remain compatible

Existing calls continue to work because request options are variadic:

```go
result, err := client.Chat.Create(ctx, params)
```

You can opt into per-request customization when needed:

```go
result, err := client.Chat.Create(ctx, params,
    api.WithHeader("X-Request-Source", "example"),
    api.WithQuery("debug", true),
)
```

### Raw responses

Use raw response helpers when you need response metadata:

```go
raw, err := client.Search.CreateRaw(ctx, params)
if err != nil {
    return err
}
fmt.Println(raw.StatusCode, raw.RequestID)
fmt.Println(raw.Data)
```

### Client cloning

Use `WithOptions()` to derive a configured client without mutating the original:

```go
regionalClient, err := client.WithOptions(
    perplexity.WithBaseURL("https://api.perplexity.ai"),
    perplexity.WithDefaultQuery("region", "us"),
)
```

---

## Validation

This release was validated with:

```bash
go test ./...
```

All packages pass, including coverage for:

- `perplexity`
- `perplexity/internal/http`
- `perplexity/chat`
- `perplexity/search`
- `perplexity/responses`
- `perplexity/embeddings`
- `perplexity/contextualizedembeddings`
- `perplexity/asyncchat`
- `perplexity/browser`
- `perplexity/types`

---

## Changelog

For the structured release summary, see `CHANGELOG.md` under `1.2.0`.

---

Thank you for using the Perplexity Go SDK.
