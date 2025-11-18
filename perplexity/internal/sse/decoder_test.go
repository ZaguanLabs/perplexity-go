package sse

import (
	"io"
	"strings"
	"testing"
)

func TestDecoder_Decode(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *Event
		wantErr bool
	}{
		{
			name:  "simple data event",
			input: "data: hello world\n\n",
			want: &Event{
				Data: "hello world",
			},
		},
		{
			name:  "event with type",
			input: "event: message\ndata: hello\n\n",
			want: &Event{
				Event: "message",
				Data:  "hello",
			},
		},
		{
			name:  "event with id",
			input: "id: 123\ndata: test\n\n",
			want: &Event{
				ID:   "123",
				Data: "test",
			},
		},
		{
			name:  "multiline data",
			input: "data: line 1\ndata: line 2\ndata: line 3\n\n",
			want: &Event{
				Data: "line 1\nline 2\nline 3",
			},
		},
		{
			name:  "with comment",
			input: ": this is a comment\ndata: hello\n\n",
			want: &Event{
				Data: "hello",
			},
		},
		{
			name:  "JSON data",
			input: "data: {\"message\":\"hello\"}\n\n",
			want: &Event{
				Data: `{"message":"hello"}`,
			},
		},
		{
			name:  "done marker",
			input: "data: [DONE]\n\n",
			want: &Event{
				Data: "[DONE]",
			},
		},
		{
			name:  "error event",
			input: "event: error\ndata: something went wrong\n\n",
			want: &Event{
				Event: "error",
				Data:  "something went wrong",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decoder := NewDecoder(strings.NewReader(tt.input))
			got, err := decoder.Decode()

			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			if got.Event != tt.want.Event {
				t.Errorf("Event = %q, want %q", got.Event, tt.want.Event)
			}
			if got.Data != tt.want.Data {
				t.Errorf("Data = %q, want %q", got.Data, tt.want.Data)
			}
			if got.ID != tt.want.ID {
				t.Errorf("ID = %q, want %q", got.ID, tt.want.ID)
			}
		})
	}
}

func TestDecoder_DecodeMultiple(t *testing.T) {
	input := `data: first event

data: second event

data: third event

`
	decoder := NewDecoder(strings.NewReader(input))

	// First event
	event1, err := decoder.Decode()
	if err != nil {
		t.Fatalf("First Decode() error = %v", err)
	}
	if event1.Data != "first event" {
		t.Errorf("First event data = %q, want %q", event1.Data, "first event")
	}

	// Second event
	event2, err := decoder.Decode()
	if err != nil {
		t.Fatalf("Second Decode() error = %v", err)
	}
	if event2.Data != "second event" {
		t.Errorf("Second event data = %q, want %q", event2.Data, "second event")
	}

	// Third event
	event3, err := decoder.Decode()
	if err != nil {
		t.Fatalf("Third Decode() error = %v", err)
	}
	if event3.Data != "third event" {
		t.Errorf("Third event data = %q, want %q", event3.Data, "third event")
	}

	// EOF
	_, err = decoder.Decode()
	if err != io.EOF {
		t.Errorf("Expected EOF, got %v", err)
	}
}

func TestDecoder_DecodeAll(t *testing.T) {
	input := `data: event 1

data: event 2

data: event 3

`
	decoder := NewDecoder(strings.NewReader(input))
	events, err := decoder.DecodeAll()
	if err != nil {
		t.Fatalf("DecodeAll() error = %v", err)
	}

	if len(events) != 3 {
		t.Fatalf("DecodeAll() returned %d events, want 3", len(events))
	}

	expectedData := []string{"event 1", "event 2", "event 3"}
	for i, event := range events {
		if event.Data != expectedData[i] {
			t.Errorf("Event %d data = %q, want %q", i, event.Data, expectedData[i])
		}
	}
}

func TestEvent_IsDone(t *testing.T) {
	tests := []struct {
		name string
		data string
		want bool
	}{
		{
			name: "done marker",
			data: "[DONE]",
			want: true,
		},
		{
			name: "done with extra",
			data: "[DONE] extra text",
			want: true,
		},
		{
			name: "not done",
			data: "regular data",
			want: false,
		},
		{
			name: "empty",
			data: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := &Event{Data: tt.data}
			if got := event.IsDone(); got != tt.want {
				t.Errorf("IsDone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_IsError(t *testing.T) {
	tests := []struct {
		name  string
		event string
		want  bool
	}{
		{
			name:  "error event",
			event: "error",
			want:  true,
		},
		{
			name:  "message event",
			event: "message",
			want:  false,
		},
		{
			name:  "empty event",
			event: "",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := &Event{Event: tt.event}
			if got := event.IsError(); got != tt.want {
				t.Errorf("IsError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvent_ParseJSON(t *testing.T) {
	event := &Event{
		Data: `{"message":"hello"}`,
	}

	data, err := event.ParseJSON()
	if err != nil {
		t.Fatalf("ParseJSON() error = %v", err)
	}

	expected := `{"message":"hello"}`
	if string(data) != expected {
		t.Errorf("ParseJSON() = %q, want %q", string(data), expected)
	}
}

func TestEvent_ParseJSON_Empty(t *testing.T) {
	event := &Event{
		Data: "",
	}

	_, err := event.ParseJSON()
	if err == nil {
		t.Error("ParseJSON() expected error for empty data")
	}
}
