package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ZaguanLabs/perplexity-go/perplexity/internal/sse"
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

// Stream represents a stream of chat completion chunks.
type Stream struct {
	decoder  *sse.Decoder
	response *http.Response
	ctx      context.Context
	err      error
}

// newStream creates a new stream from an HTTP response.
func newStream(ctx context.Context, resp *http.Response) *Stream {
	return &Stream{
		decoder:  sse.NewDecoder(resp.Body),
		response: resp,
		ctx:      ctx,
	}
}

// Next returns the next chunk in the stream.
// Returns io.EOF when the stream is complete.
func (s *Stream) Next() (*types.StreamChunk, error) {
	// Check if stream already errored
	if s.err != nil {
		return nil, s.err
	}

	// Check context cancellation
	select {
	case <-s.ctx.Done():
		s.err = s.ctx.Err()
		return nil, s.err
	default:
	}

	// Decode next SSE event
	event, err := s.decoder.Decode()
	if err != nil {
		if err == io.EOF {
			s.err = io.EOF
		}
		return nil, err
	}

	// Check for done marker
	if event.IsDone() {
		s.err = io.EOF
		return nil, io.EOF
	}

	// Check for error event
	if event.IsError() {
		s.err = fmt.Errorf("stream error: %s", event.Data)
		return nil, s.err
	}

	// Parse JSON data
	data, err := event.ParseJSON()
	if err != nil {
		s.err = fmt.Errorf("failed to parse event data: %w", err)
		return nil, s.err
	}

	// Unmarshal into StreamChunk
	var chunk types.StreamChunk
	if err := json.Unmarshal(data, &chunk); err != nil {
		s.err = fmt.Errorf("failed to unmarshal chunk: %w", err)
		return nil, s.err
	}

	return &chunk, nil
}

// Close closes the stream and releases resources.
func (s *Stream) Close() error {
	if s.response != nil && s.response.Body != nil {
		return s.response.Body.Close()
	}
	return nil
}

// Recv is an alias for Next for compatibility.
func (s *Stream) Recv() (*types.StreamChunk, error) {
	return s.Next()
}

// Iter returns a channel that yields stream chunks.
// The channel is closed when the stream ends or encounters an error.
// The caller must call Close() to release resources.
func (s *Stream) Iter() <-chan *types.StreamChunk {
	ch := make(chan *types.StreamChunk)
	go func() {
		defer close(ch)
		for {
			chunk, err := s.Next()
			if err != nil {
				return
			}
			select {
			case ch <- chunk:
			case <-s.ctx.Done():
				return
			}
		}
	}()
	return ch
}
