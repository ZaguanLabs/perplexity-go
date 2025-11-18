package chat

import (
	"io"

	"github.com/perplexityai/perplexity-go/perplexity/types"
)

// Stream represents a stream of chat completion chunks.
// This will be fully implemented in Phase 4.
type Stream struct {
	// Implementation will be added in Phase 4
}

// Next returns the next chunk in the stream.
// Returns io.EOF when the stream is complete.
func (s *Stream) Next() (*types.StreamChunk, error) {
	// Will be implemented in Phase 4
	return nil, io.EOF
}

// Close closes the stream and releases resources.
func (s *Stream) Close() error {
	// Will be implemented in Phase 4
	return nil
}
