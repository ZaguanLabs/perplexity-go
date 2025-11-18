package types

import (
	"encoding/json"
	"testing"
)

func TestChatMessage_UnmarshalJSON_TextContent(t *testing.T) {
	jsonData := `{
		"role": "user",
		"content": "Hello, world!"
	}`

	var msg ChatMessage
	err := json.Unmarshal([]byte(jsonData), &msg)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if msg.Role != RoleUser {
		t.Errorf("Role = %v, want %v", msg.Role, RoleUser)
	}

	textContent, ok := msg.Content.(TextContent)
	if !ok {
		t.Fatalf("Content is not TextContent, got %T", msg.Content)
	}

	if string(textContent) != "Hello, world!" {
		t.Errorf("Content = %q, want %q", textContent, "Hello, world!")
	}
}

func TestChatMessage_UnmarshalJSON_StructuredContent(t *testing.T) {
	jsonData := `{
		"role": "user",
		"content": [
			{
				"type": "text",
				"text": "What's in this image?"
			},
			{
				"type": "image_url",
				"image_url": "https://example.com/image.jpg"
			}
		]
	}`

	var msg ChatMessage
	err := json.Unmarshal([]byte(jsonData), &msg)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if msg.Role != RoleUser {
		t.Errorf("Role = %v, want %v", msg.Role, RoleUser)
	}

	structuredContent, ok := msg.Content.(StructuredContent)
	if !ok {
		t.Fatalf("Content is not StructuredContent, got %T", msg.Content)
	}

	if len(structuredContent) != 2 {
		t.Fatalf("StructuredContent length = %d, want 2", len(structuredContent))
	}

	// Check first chunk (text)
	textChunk, ok := structuredContent[0].(TextChunk)
	if !ok {
		t.Errorf("First chunk is not TextChunk, got %T", structuredContent[0])
	}
	if textChunk.Text != "What's in this image?" {
		t.Errorf("TextChunk.Text = %q, want %q", textChunk.Text, "What's in this image?")
	}

	// Check second chunk (image)
	imageChunk, ok := structuredContent[1].(ImageChunk)
	if !ok {
		t.Errorf("Second chunk is not ImageChunk, got %T", structuredContent[1])
	}
	if imageChunk.Type != "image_url" {
		t.Errorf("ImageChunk.Type = %q, want %q", imageChunk.Type, "image_url")
	}
}

func TestChatMessage_MarshalJSON_TextContent(t *testing.T) {
	msg := ChatMessage{
		Role:    RoleAssistant,
		Content: TextContent("Hello!"),
	}

	data, err := json.Marshal(msg)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	// Unmarshal to verify
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("Unmarshal verification failed: %v", err)
	}

	if result["role"] != "assistant" {
		t.Errorf("role = %v, want assistant", result["role"])
	}

	if result["content"] != "Hello!" {
		t.Errorf("content = %v, want Hello!", result["content"])
	}
}

func TestHelperFunctions(t *testing.T) {
	t.Run("SystemMessage", func(t *testing.T) {
		msg := SystemMessage("You are a helpful assistant")
		if msg.Role != RoleSystem {
			t.Errorf("Role = %v, want %v", msg.Role, RoleSystem)
		}
		content, ok := msg.Content.(TextContent)
		if !ok {
			t.Fatalf("Content is not TextContent")
		}
		if string(content) != "You are a helpful assistant" {
			t.Errorf("Content = %q, want %q", content, "You are a helpful assistant")
		}
	})

	t.Run("UserMessage", func(t *testing.T) {
		msg := UserMessage("Hello")
		if msg.Role != RoleUser {
			t.Errorf("Role = %v, want %v", msg.Role, RoleUser)
		}
	})

	t.Run("AssistantMessage", func(t *testing.T) {
		msg := AssistantMessage("Hi there")
		if msg.Role != RoleAssistant {
			t.Errorf("Role = %v, want %v", msg.Role, RoleAssistant)
		}
	})
}

func TestPointerHelpers(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		s := String("test")
		if s == nil {
			t.Fatal("String() returned nil")
		}
		if *s != "test" {
			t.Errorf("*String() = %q, want %q", *s, "test")
		}
	})

	t.Run("Int", func(t *testing.T) {
		i := Int(42)
		if i == nil {
			t.Fatal("Int() returned nil")
		}
		if *i != 42 {
			t.Errorf("*Int() = %d, want 42", *i)
		}
	})

	t.Run("Float64", func(t *testing.T) {
		f := Float64(3.14)
		if f == nil {
			t.Fatal("Float64() returned nil")
		}
		if *f != 3.14 {
			t.Errorf("*Float64() = %f, want 3.14", *f)
		}
	})

	t.Run("Bool", func(t *testing.T) {
		b := Bool(true)
		if b == nil {
			t.Fatal("Bool() returned nil")
		}
		if *b != true {
			t.Errorf("*Bool() = %v, want true", *b)
		}
	})
}
