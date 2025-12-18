package sse

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

// Benchmark SSE event parsing

func BenchmarkDecoder_Decode_SimpleEvent(b *testing.B) {
	data := "data: {\"test\": \"value\"}\n\n"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(data))
		decoder := NewDecoder(reader)
		_, err := decoder.Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecoder_Decode_ComplexEvent(b *testing.B) {
	data := `event: message
id: msg-123
data: {"id":"test","model":"sonar","choices":[{"index":0,"delta":{"role":"assistant","content":"Hello"}}]}

`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(data))
		decoder := NewDecoder(reader)
		_, err := decoder.Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecoder_Decode_MultilineData(b *testing.B) {
	data := `data: {"chunk": "line1"}
data: {"chunk": "line2"}
data: {"chunk": "line3"}

`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(data))
		decoder := NewDecoder(reader)
		_, err := decoder.Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecoder_Decode_LargePayload(b *testing.B) {
	// Create a large JSON payload
	largeJSON := `{"id":"test-123","model":"sonar-pro","choices":[{"index":0,"delta":{"role":"assistant","content":"` +
		strings.Repeat("This is a long response with lots of text. ", 100) +
		`"}}],"usage":{"completion_tokens":500,"prompt_tokens":100,"total_tokens":600}}`
	data := "data: " + largeJSON + "\n\n"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(data))
		decoder := NewDecoder(reader)
		_, err := decoder.Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecoder_Decode_MultipleEvents(b *testing.B) {
	data := `data: {"chunk": 1}

data: {"chunk": 2}

data: {"chunk": 3}

data: [DONE]

`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(data))
		decoder := NewDecoder(reader)

		// Decode all events
		for {
			event, err := decoder.Decode()
			if err != nil {
				break
			}
			if event.Data == "[DONE]" {
				break
			}
		}
	}
}

func BenchmarkDecoder_Decode_WithAllFields(b *testing.B) {
	data := `event: message
id: msg-456
retry: 1000
data: {"status": "processing"}

`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(data))
		decoder := NewDecoder(reader)
		_, err := decoder.Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecoder_Decode_StreamSimulation(b *testing.B) {
	// Simulate a realistic streaming response
	events := []string{
		"data: {\"id\":\"1\",\"choices\":[{\"delta\":{\"role\":\"assistant\"}}]}\n\n",
		"data: {\"id\":\"1\",\"choices\":[{\"delta\":{\"content\":\"Hello\"}}]}\n\n",
		"data: {\"id\":\"1\",\"choices\":[{\"delta\":{\"content\":\" world\"}}]}\n\n",
		"data: {\"id\":\"1\",\"choices\":[{\"delta\":{\"content\":\"!\"}}]}\n\n",
		"data: {\"id\":\"1\",\"choices\":[{\"finish_reason\":\"stop\"}],\"usage\":{\"total_tokens\":10}}\n\n",
		"data: [DONE]\n\n",
	}

	stream := strings.Join(events, "")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(stream))
		decoder := NewDecoder(reader)

		count := 0
		for {
			event, err := decoder.Decode()
			if err != nil {
				break
			}
			count++
			if event.Data == "[DONE]" {
				break
			}
		}
	}
}

func BenchmarkDecoder_Decode_Allocations(b *testing.B) {
	data := "data: {\"test\": \"value\"}\n\n"

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(data))
		decoder := NewDecoder(reader)
		_, err := decoder.Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkDecoder_Decode_BufferedStream(b *testing.B) {
	// Create a buffered stream with multiple events
	var buf bytes.Buffer
	for i := 0; i < 100; i++ {
		buf.WriteString("data: {\"chunk\": ")
		buf.WriteString(string(rune('0' + (i % 10))))
		buf.WriteString("}\n\n")
	}

	data := buf.Bytes()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(bytes.NewReader(data))
		decoder := NewDecoder(reader)

		// Decode all events
		for {
			_, err := decoder.Decode()
			if err != nil {
				break
			}
		}
	}
}

func BenchmarkDecoder_Decode_CommentFiltering(b *testing.B) {
	data := `: This is a comment
: Another comment
data: {"actual": "data"}

`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reader := bufio.NewReader(strings.NewReader(data))
		decoder := NewDecoder(reader)
		_, err := decoder.Decode()
		if err != nil {
			b.Fatal(err)
		}
	}
}
