# Perplexity Go SDK v1.10.0

Targets the official Perplexity Python SDK 0.50.0.

- Use `client.Context.Create` to fetch documents, links and highlights. Search and Context accept a `*time.Time` snapshot.
- Use `client.Monitors` to create, retrieve, list, update, delete, pause, resume and run monitors. `Monitors.Runs` exposes run history, and `Monitors.Webhook.Test` checks delivery. All methods have raw-response variants. Pass pagination through `api.WithQuery` and the creation idempotency key through `api.WithHeader`.
- Embeddings supports both v2 context models, document batches, image inputs and nested embedding results. Decisions reports per-tool invocation counts. Responses accepts the built-in search skill.
- The upstream Python tag and the PyPI wheel contain 118 matching Python modules. Building the source required temporarily removing its obsolete Apache license classifier. The checkout's original metadata was restored after building and unpacking the source distribution.

Verified offline with HTTP contract tests, the full race-enabled test suite, vet and build.
