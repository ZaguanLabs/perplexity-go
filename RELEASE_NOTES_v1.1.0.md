# Perplexity Go SDK v1.1.0

**Release Date:** March 24, 2026  
**Status:** Stable Release

---

## Release Summary

`v1.1.0` is a major parity release for the Perplexity Go SDK.

This version brings the Go SDK substantially closer to the official Python SDK by expanding the public client surface, tightening request and response schemas, aligning transport behavior, and strengthening validation across resources.

If you want the short version: this release makes the Go SDK more complete, more predictable, and much closer to the reference implementation.

---

## Highlights

### New public SDK resources
The Go client now exposes the newer resource families that were previously missing from the public surface:

- `Responses`
- `Embeddings`
- `ContextualizedEmbeddings`
- `Browser` with `Browser.Sessions`

### Python SDK parity improvements
This release closes a broad set of parity gaps with the official Python SDK, including:

- shared message-content unions and nullable content handling
- retry behavior and retry header support
- default client metadata headers
- async chat request and response coverage
- stricter typed search query handling

### Stronger validation and tests
`v1.1.0` also expands the validation story around the SDK by adding focused tests for the areas most likely to drift from parity:

- retry semantics
- typed error mapping
- async chat flows
- structured chat content unions
- cross-resource request validation

---

## What’s New

### Added

- Public client initialization for:
  - `Client.Responses`
  - `Client.Embeddings`
  - `Client.ContextualizedEmbeddings`
  - `Client.Browser`
- Dedicated `asyncchat` service coverage for:
  - create
  - list
  - get
  - validation
  - header and query propagation
- Shared message-model tests covering:
  - `content: null`
  - `file_url` unions
  - `pdf_url` unions
  - `video_url` unions
- Retry coverage for:
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
- Fixed coverage gaps around retry behavior and typed error mapping
- Fixed missing transport metadata behavior relative to the Python SDK

---

## Upgrade Notes

### Install or update

```bash
go get github.com/ZaguanLabs/perplexity-go/perplexity@v1.1.0
```

### Search query typing

If you were previously passing raw strings into search requests, update call sites to use the typed `search.Query` wrapper.

### Transport metadata

The client now sends default `X-Stainless-*` headers automatically. If you already provide those headers explicitly through client defaults, your explicit values still win.

---

## Validation

This release was validated with:

```bash
go test ./...
```

All packages pass, including expanded coverage for:

- `perplexity/internal/http`
- `perplexity/asyncchat`
- `perplexity/types`
- `perplexity/browser`
- `perplexity/responses`
- `perplexity/embeddings`
- `perplexity/contextualizedembeddings`

---

## Changelog

For the structured release summary, see `CHANGELOG.md` under `1.1.0`.

---

Thank you for using the Perplexity Go SDK.
