# Perplexity Go SDK v1.6.0

This release ports the public API additions through the official Perplexity Python SDK v0.43.5.

## Added

- Responses profile references and custom skill references.
- Reasoning items in the Responses input union, including summaries, encrypted content, status, and optional reasoning text.
- Multimodal function-call output parts with text and image variants.
- Optional function-call output ID and status fields.

## Compatibility

String function-call outputs remain accepted. The additions extend tagged unions and request structures without changing existing wire shapes.

## Upstream packaging

The official v0.43.5 tag contains an incompatible PEP 639 license expression and obsolete license classifier combination. The pristine source snapshot was restored exactly after build artifacts were generated with a temporary metadata-only workaround.
