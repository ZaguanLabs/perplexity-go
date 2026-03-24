package responses

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ZaguanLabs/perplexity-go/perplexity/internal/sse"
)

type Stream struct {
	decoder  *sse.Decoder
	response *http.Response
	ctx      context.Context
	err      error
}

func newStream(ctx context.Context, resp *http.Response) *Stream {
	return &Stream{
		decoder:  sse.NewDecoder(resp.Body),
		response: resp,
		ctx:      ctx,
	}
}

func (s *Stream) Next() (*StreamEvent, error) {
	if s.err != nil {
		return nil, s.err
	}

	select {
	case <-s.ctx.Done():
		s.err = s.ctx.Err()
		return nil, s.err
	default:
	}

	event, err := s.decoder.Decode()
	if err != nil {
		if err == io.EOF {
			s.err = io.EOF
		}
		return nil, err
	}

	if event.IsDone() {
		s.err = io.EOF
		return nil, io.EOF
	}

	if event.IsError() {
		s.err = fmt.Errorf("stream error: %s", event.Data)
		return nil, s.err
	}

	data, err := event.ParseJSON()
	if err != nil {
		s.err = fmt.Errorf("failed to parse event data: %w", err)
		return nil, s.err
	}

	var chunk StreamEvent
	if err := json.Unmarshal(data, &chunk); err != nil {
		s.err = fmt.Errorf("failed to unmarshal chunk: %w", err)
		return nil, s.err
	}

	return &chunk, nil
}

func (s *Stream) Recv() (*StreamEvent, error) {
	return s.Next()
}

func (s *Stream) Close() error {
	if s.response != nil && s.response.Body != nil {
		return s.response.Body.Close()
	}
	return nil
}

func (s *Stream) Iter() <-chan *StreamEvent {
	ch := make(chan *StreamEvent)
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
