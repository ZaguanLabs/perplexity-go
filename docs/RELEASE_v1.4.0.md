# Perplexity Go SDK v1.4.0

This release ports the public API additions through the official Perplexity Python SDK v0.43.1.

## Added

- Responses cancellation: `Responses.Cancel` and `Responses.CancelRaw`.
- Response sandbox files: `Responses.Files.List`, `ListRaw`, `Content`, and `ContentRaw`.
- MCP request tools and MCP list/call output items.
- Built-in and inline response skills.
- Response continuation and storage fields.
- Advisor, sandbox, shared-file, and unknown output-item variants.
- `response.skill.loaded` streaming events.
- Chat response-formatting locale support.

## Upstream packaging

The official v0.43.1 tag cannot be built unmodified by setuptools: it has no
explicit build backend and combines a PEP 639 license expression with the
superseded Apache license classifier. The downloaded source snapshot remains
pristine. Build artifacts were generated from a temporary copy after removing
only that obsolete classifier.
