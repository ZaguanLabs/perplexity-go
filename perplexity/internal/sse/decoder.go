package sse

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Event represents a Server-Sent Event.
type Event struct {
	// Event is the event type (e.g., "message", "error").
	Event string

	// Data is the event data.
	Data string

	// ID is the event ID.
	ID string

	// Retry is the reconnection time in milliseconds.
	Retry int
}

// Decoder decodes Server-Sent Events from a stream.
type Decoder struct {
	reader *bufio.Reader
}

// NewDecoder creates a new SSE decoder.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: bufio.NewReader(r),
	}
}

// Decode reads and decodes the next SSE event.
// Returns io.EOF when the stream is closed.
func (d *Decoder) Decode() (*Event, error) {
	event := &Event{}
	var dataLines []string

	for {
		line, err := d.reader.ReadString('\n')
		if err != nil {
			if err == io.EOF && len(dataLines) > 0 {
				// Return accumulated data before EOF
				event.Data = strings.Join(dataLines, "\n")
				return event, nil
			}
			return nil, err
		}

		// Remove trailing newline
		line = strings.TrimSuffix(line, "\n")
		line = strings.TrimSuffix(line, "\r")

		// Empty line signals end of event
		if line == "" {
			if len(dataLines) > 0 || event.Event != "" || event.ID != "" {
				event.Data = strings.Join(dataLines, "\n")
				return event, nil
			}
			// Skip empty lines between events
			continue
		}

		// Skip comments
		if strings.HasPrefix(line, ":") {
			continue
		}

		// Parse field
		colonIndex := strings.Index(line, ":")
		if colonIndex == -1 {
			// Field with no value
			continue
		}

		field := line[:colonIndex]
		value := line[colonIndex+1:]

		// Remove leading space from value
		if len(value) > 0 && value[0] == ' ' {
			value = value[1:]
		}

		switch field {
		case "event":
			event.Event = value
		case "data":
			dataLines = append(dataLines, value)
		case "id":
			event.ID = value
		case "retry":
			// Parse retry as integer (not implemented for now)
			event.Retry = 0
		}
	}
}

// DecodeAll reads all events from the stream until EOF.
func (d *Decoder) DecodeAll() ([]*Event, error) {
	var events []*Event

	for {
		event, err := d.Decode()
		if err == io.EOF {
			break
		}
		if err != nil {
			return events, err
		}
		events = append(events, event)
	}

	return events, nil
}

// ParseJSON attempts to parse the event data as JSON.
// This is a helper that returns the raw data for JSON unmarshaling.
func (e *Event) ParseJSON() ([]byte, error) {
	if e.Data == "" {
		return nil, fmt.Errorf("event data is empty")
	}
	return []byte(e.Data), nil
}

// IsError returns true if this is an error event.
func (e *Event) IsError() bool {
	return e.Event == "error"
}

// IsDone returns true if this is a done/completion event.
func (e *Event) IsDone() bool {
	return strings.HasPrefix(e.Data, "[DONE]")
}

// String returns a string representation of the event.
func (e *Event) String() string {
	var buf bytes.Buffer
	if e.Event != "" {
		fmt.Fprintf(&buf, "event: %s\n", e.Event)
	}
	if e.ID != "" {
		fmt.Fprintf(&buf, "id: %s\n", e.ID)
	}
	if e.Data != "" {
		for _, line := range strings.Split(e.Data, "\n") {
			fmt.Fprintf(&buf, "data: %s\n", line)
		}
	}
	return buf.String()
}
