# Perplexity Go SDK v1.8.0

Targets official Perplexity Python SDK 0.43.8.

- Adds image-search tools, filters, results and both reasoning stream events.
- Adds file input parts and Responses tool-choice selection (string modes or object selectors).
- Adds MCP approval policy modes and per-tool/read-only filters, approval request outputs, approval response inputs, MCP call/list-tools replay, approval linkage and read-only annotations.
- Registers namespace tools for construction, decoding and access.

The official GitHub v0.43.8 archive is retained in the ignored `docs/perplexity-py/` directory with its SHA-256 manifest. Built with `python -m build` and the generated sdist unpacked. Upstream still has conflicting PEP 639 license metadata: the obsolete classifier was temporarily removed for the build, and the original pyproject restored byte for byte.

Validation: complete offline race suite, build, vet and diff checks. Regression tests exercise approval decision false values, replay round-trips, tool filters, file inputs, image outputs and streaming events. No live API calls.
