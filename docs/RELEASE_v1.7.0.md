# Perplexity Go SDK v1.7.0

Targets official Perplexity Python SDK 0.43.6.

- Search requests accept `search.SearchTypeFast`.
- Responses web-search tools expose optional `SearchType` with `web` and `fast` modes.
- Responses reasoning accepts `responses.ReasoningEffortNone`.
- Existing modes and omitted optional fields retain their wire behavior.

The official GitHub v0.43.6 source was downloaded (PyPI provides only a wheel), built with `python -m build`, and its generated source distribution unpacked. As in v0.43.5, building required temporarily removing the obsolete Apache license classifier; the original source metadata was restored byte for byte. Source archive SHA-256: `4edd597adad8bad36ea7f781100367578b3006a14ed79c99555ed5473d84a4d0`.

Validation: full race-tested suite, `go build ./...`, `go vet ./...`, and `git diff --check` pass. Tests cover new wire values, tagged-union decoding, and optional-field omission. No live API calls were made.
