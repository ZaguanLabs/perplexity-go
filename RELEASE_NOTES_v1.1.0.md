# 🎉 Perplexity Go SDK v1.1.0

**Release Date:** March 24, 2026  
**Status:** Stable Release ✅

---

## Overview

`v1.1.0` is a parity-focused release that brings the Go SDK into much closer alignment with the official Python SDK across client behavior, transport metadata, shared schemas, and newer API resources.

This release expands the public SDK surface, tightens request and response modeling, improves retry semantics, and significantly strengthens cross-resource test coverage.

---

## Highlights

### Expanded SDK Surface
- Added `Responses`
- Added `Embeddings`
- Added `ContextualizedEmbeddings`
- Added `Browser` with `Browser.Sessions`

### Python SDK Parity Improvements
- Aligned retry semantics with the Python SDK
- Added default `X-Stainless-*` transport metadata headers
- Tightened shared chat message unions and nullable content handling
- Improved `asyncchat` request and response coverage
- Updated examples to match the stricter typed `search.Query` API

### Validation and Reliability
- Added focused tests for retry behavior, typed errors, async chat flows, and shared content unions
- Verified full test suite with `go test ./...`

---

## What Changed

### Added
- Public client initialization for:
  - `Client.Responses`
  - `Client.Embeddings`
  - `Client.ContextualizedEmbeddings`
  - `Client.Browser`
- Dedicated `asyncchat` service tests covering:
  - create
  - list
  - get
  - validation
  - header and query propagation
- Shared message-model tests for:
  - `content: null`
  - `file_url` unions
  - `pdf_url` unions
  - `video_url` unions
- Retry tests for:
  - `retry-after-ms`
  - `retry-after`
  - `x-should-retry`
- Default `X-Stainless-*` client metadata headers:
  - `X-Stainless-Lang`
  - `X-Stainless-Package-Version`
  - `X-Stainless-OS`
  - `X-Stainless-Arch`
  - `X-Stainless-Runtime`
  - `X-Stainless-Runtime-Version`

### Changed
- Updated SDK version to `v1.1.0`
- Refined shared chat content parsing to better match Python SDK union behavior
- Updated backoff handling to honor server retry headers with capped exponential backoff and jitter
- Improved naming consistency in `asyncchat.Get()` by using `requestID`

### Fixed
- Fixed compile fallout in `examples/search/main.go` after stricter `search.Query` typing
- Fixed test coverage gaps around retry behavior and typed error mapping
- Fixed missing transport metadata behavior relative to the Python SDK

---

## Upgrade Notes

### Version Bump
Update to the new release with:

```bash
go get github.com/ZaguanLabs/perplexity-go/perplexity@v1.1.0
```

### Search Query Typing
If you were previously passing raw strings for search requests, use the typed `search.Query` wrapper instead.

### Transport Metadata
The client now sends default `X-Stainless-*` headers automatically. If you already set these manually, your explicit default headers still take precedence.

---

## Validation

This release was validated with:

```bash
go test ./...
```

All packages pass, including the expanded coverage for:
- `perplexity/internal/http`
- `perplexity/asyncchat`
- `perplexity/types`
- `perplexity/browser`
- `perplexity/responses`
- `perplexity/embeddings`
- `perplexity/contextualizedembeddings`

---

## Changelog Reference

For the structured release summary, see:

- `CHANGELOG.md` → `1.1.0`

---

**Thank you for using the Perplexity Go SDK.**
