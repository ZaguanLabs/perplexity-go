# Perplexity Go SDK v1.5.0

This release ports the public API additions through the official Perplexity Python SDK v0.43.3.

## Added

- Connector tools with connector IDs and arbitrary connector configuration.
- Namespace tools and namespace function definitions.
- Typed `tool_search_output` response items.
- Responses `temperature` and `top_p` controls.
- MCP `defer_loading` and output-side `connector_id` fields.

## Compatibility

Existing v1.4.0 APIs remain source compatible. The additions extend tagged unions and
request structures without changing existing wire shapes.

## Upstream packaging

The official v0.43.3 tag retains an incompatible PEP 639 license expression and obsolete
license classifier combination. The pristine source snapshot remains unchanged; build
artifacts were generated from a temporary copy with only that classifier removed.
