# Perplexity Go SDK v1.9.0

Targets official Python SDK 0.46.1. Adds the Agent model catalog, Decisions, Files, Batches, managed Skills and nested revisions. Typed methods share the established authentication, request-option, retry, error, and raw-response paths. Pagination and revision queries use `api.WithQuery`; idempotency keys use `api.WithHeader`.

File and skill uploads accept `api.Upload` readers, including open files. `api.UploadPath` reads a local path. Multipart bodies are prepared once so retries resend the complete upload. Binary file contents remain bytes.

Responses gains effort (preset remains available but is deprecated upstream), prompt cache keys, service tiers, step-budget reminders, extraction token/cost details, optional input-message discriminators and typed stream reconnection. `RetrieveStream` requires a non-negative safe-integer cursor. Base URL queries and escaped identifiers survive path joining.

Source: official GitHub v0.46.1 archive, SHA-256 `ac045213f94a14dd134ab7a822c4ffa09511e935e756e479d2486d652d706f79`. PyPI publishes a wheel only. Built using `python -m build`; the obsolete license classifier was temporarily removed to resolve upstream PEP 639 metadata conflict, then the original pyproject was restored byte-for-byte. Built sdist unpacked under ignored `docs/perplexity-py/`.

Offline tests exercise every new resource route, raw metadata, multipart file handles, retry replay, binary bytes, query preservation, escaped and invalid IDs, zero/false values and streaming reconnection.
