# Perplexity Go SDK v1.3.0

**Release Date:** June 4, 2026  
**Status:** Stable Release

---

## Release Summary

`v1.3.0` updates the Go SDK to match the official Python SDK v0.37.0 API surface.

This release focuses on closing schema and resource drift introduced since the previous parity pass. It adds the response retrieval endpoint, newer response request fields and tool variants, updated search parameters, and the official nested user location shape for chat web search options.

---

## Highlights

### Responses API parity

The responses service now supports retrieving background responses by ID:

- `Responses.Retrieve()`
- `Responses.RetrieveRaw()`

The response types now include the Python SDK v0.37.0 status, reasoning, background, web search, and tool variants.

### Search API parity

`search.SearchParams` now includes:

- `search_context_size`
- `search_type`

### Chat request parity

`chat.WebSearchOptions.UserLocation` now matches the official optional location object:

- `city`
- `country`
- `latitude`
- `longitude`
- `region`

---

## What’s New

### Added

- `Responses.Retrieve(ctx, responseID, opts...)`
- `Responses.RetrieveRaw(ctx, responseID, opts...)`
- `responses.CreateParams.Background`
- `responses.CreateResponse.Background`
- `responses.StatusQueued`
- `responses.StatusCancelled`
- `responses.ReasoningEffortMinimal`
- `responses.ReasoningEffortXHigh`
- `responses.WebSearchTool.SearchContextSize`
- `responses.PeopleSearchTool`
- `responses.FinanceSearchTool`
- `responses.SandboxTool`
- `search.SearchParams.SearchContextSize`
- `search.SearchParams.SearchType`
- `.gitignore` entry for `docs/perplexity-py/`

### Changed

- Updated SDK version to `v1.3.0`.
- Updated `chat.WebSearchOptions.UserLocation` to use optional city/country/region and optional latitude/longitude fields.
- Removed the non-official `_is_browser_agent` chat request field from `chat.CompletionParams`.

---

## Upgrade Notes

### Install or update

```bash
go get github.com/ZaguanLabs/perplexity-go/perplexity@v1.3.0
```

### Response retrieval

Use `Retrieve` to poll a background response:

```go
response, err := client.Responses.Retrieve(ctx, "resp_123")
if err != nil {
    return err
}
fmt.Println(response.Status)
```

### Chat user location

`chat.UserLocation` now uses pointers for all fields:

```go
lat := 37.7749
lng := -122.4194
country := "US"

params.WebSearchOptions = &chat.WebSearchOptions{
    UserLocation: &chat.UserLocation{
        Country:   &country,
        Latitude:  &lat,
        Longitude: &lng,
    },
}
```

### Source compatibility note

The previously exposed `_is_browser_agent` request field was not present in the official Python SDK v0.37.0 and has been removed to maintain official SDK parity.

---

## Validation

This release was validated with:

```bash
go test ./...
```

All packages pass, including:

- `perplexity`
- `perplexity/internal/http`
- `perplexity/internal/sse`
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

For the structured release summary, see `CHANGELOG.md` under `1.3.0`.

