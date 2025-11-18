package types

import (
	"encoding/json"
	"fmt"
)

// Role represents a chat message role.
type Role string

const (
	// RoleSystem is the system role.
	RoleSystem Role = "system"

	// RoleUser is the user role.
	RoleUser Role = "user"

	// RoleAssistant is the assistant role.
	RoleAssistant Role = "assistant"

	// RoleTool is the tool role.
	RoleTool Role = "tool"
)

// ChatMessage represents a chat message.
type ChatMessage struct {
	// Role is the role of the message sender.
	Role Role `json:"role"`

	// Content is the message content (can be string or structured content).
	Content MessageContent `json:"content"`

	// ReasoningSteps contains reasoning steps (optional).
	ReasoningSteps []ReasoningStep `json:"reasoning_steps,omitempty"`

	// ToolCalls contains tool calls made by the assistant (optional).
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`
}

// MessageContent represents message content that can be either a simple string
// or structured content with multiple chunks.
type MessageContent interface {
	isMessageContent()
}

// TextContent is a simple text message.
type TextContent string

func (TextContent) isMessageContent() {}

// MarshalJSON implements json.Marshaler for TextContent.
func (t TextContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

// StructuredContent is structured message content with multiple chunks.
type StructuredContent []ContentChunk

func (StructuredContent) isMessageContent() {}

// ContentChunk represents a single chunk of structured content.
type ContentChunk interface {
	isContentChunk()
	GetType() string
}

// TextChunk represents a text content chunk.
type TextChunk struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func (TextChunk) isContentChunk() {}

// GetType returns the type of the text chunk.
func (t TextChunk) GetType() string { return t.Type }

// ImageChunk represents an image content chunk.
type ImageChunk struct {
	Type     string   `json:"type"`
	ImageURL ImageURL `json:"image_url"`
}

func (ImageChunk) isContentChunk() {}

// GetType returns the type of the image chunk.
func (i ImageChunk) GetType() string { return i.Type }

// UnmarshalJSON implements custom unmarshaling for ImageChunk.
func (i *ImageChunk) UnmarshalJSON(data []byte) error {
	var temp struct {
		Type     string          `json:"type"`
		ImageURL json.RawMessage `json:"image_url"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	i.Type = temp.Type

	// Try to unmarshal as string first
	var urlStr string
	if err := json.Unmarshal(temp.ImageURL, &urlStr); err == nil {
		i.ImageURL = ImageURLString(urlStr)
		return nil
	}

	// Otherwise, unmarshal as object
	var urlObj ImageURLObject
	if err := json.Unmarshal(temp.ImageURL, &urlObj); err != nil {
		return err
	}
	i.ImageURL = urlObj
	return nil
}

// ImageURL can be either a string URL or an object with additional properties.
type ImageURL interface {
	isImageURL()
}

// ImageURLString is a simple string URL.
type ImageURLString string

func (ImageURLString) isImageURL() {}

// MarshalJSON implements json.Marshaler for ImageURLString.
func (u ImageURLString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(u))
}

// ImageURLObject is an image URL with additional properties.
type ImageURLObject struct {
	URL string `json:"url"`
}

func (ImageURLObject) isImageURL() {}

// FileChunk represents a file content chunk.
type FileChunk struct {
	Type     string  `json:"type"`
	FileURL  FileURL `json:"file_url"`
	FileName *string `json:"file_name,omitempty"`
}

func (FileChunk) isContentChunk() {}

// GetType returns the type of the file chunk.
func (f FileChunk) GetType() string { return f.Type }

// FileURL can be either a string URL or an object.
type FileURL interface {
	isFileURL()
}

// FileURLString is a simple string URL.
type FileURLString string

func (FileURLString) isFileURL() {}

// MarshalJSON implements json.Marshaler for FileURLString.
func (u FileURLString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(u))
}

// FileURLObject is a file URL object.
type FileURLObject struct {
	URL string `json:"url"`
}

func (FileURLObject) isFileURL() {}

// PDFChunk represents a PDF content chunk.
type PDFChunk struct {
	Type   string `json:"type"`
	PDFURL PDFURL `json:"pdf_url"`
}

func (PDFChunk) isContentChunk() {}

// GetType returns the type of the PDF chunk.
func (p PDFChunk) GetType() string { return p.Type }

// PDFURL can be either a string URL or an object.
type PDFURL interface {
	isPDFURL()
}

// PDFURLString is a simple string URL.
type PDFURLString string

func (PDFURLString) isPDFURL() {}

// MarshalJSON implements json.Marshaler for PDFURLString.
func (u PDFURLString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(u))
}

// PDFURLObject is a PDF URL object.
type PDFURLObject struct {
	URL string `json:"url"`
}

func (PDFURLObject) isPDFURL() {}

// VideoChunk represents a video content chunk.
type VideoChunk struct {
	Type     string   `json:"type"`
	VideoURL VideoURL `json:"video_url"`
}

func (VideoChunk) isContentChunk() {}

// GetType returns the type of the video chunk.
func (v VideoChunk) GetType() string { return v.Type }

// VideoURL can be either a string URL or an object.
type VideoURL interface {
	isVideoURL()
}

// VideoURLString is a simple string URL.
type VideoURLString string

func (VideoURLString) isVideoURL() {}

// MarshalJSON implements json.Marshaler for VideoURLString.
func (u VideoURLString) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(u))
}

// VideoURLObject is a video URL with additional properties.
type VideoURLObject struct {
	URL           string      `json:"url"`
	FrameInterval interface{} `json:"frame_interval,omitempty"` // Can be string or int
}

func (VideoURLObject) isVideoURL() {}

// UnmarshalJSON implements custom unmarshaling for ChatMessage to handle
// the content field which can be either a string or structured content.
func (m *ChatMessage) UnmarshalJSON(data []byte) error {
	// First, unmarshal into a temporary structure
	var temp struct {
		Role           Role            `json:"role"`
		Content        json.RawMessage `json:"content"`
		ReasoningSteps []ReasoningStep `json:"reasoning_steps,omitempty"`
		ToolCalls      []ToolCall      `json:"tool_calls,omitempty"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	m.Role = temp.Role
	m.ReasoningSteps = temp.ReasoningSteps
	m.ToolCalls = temp.ToolCalls

	// Try to unmarshal content as a string first
	var contentStr string
	if err := json.Unmarshal(temp.Content, &contentStr); err == nil {
		m.Content = TextContent(contentStr)
		return nil
	}

	// If that fails, try as structured content
	var chunks []json.RawMessage
	if err := json.Unmarshal(temp.Content, &chunks); err != nil {
		return fmt.Errorf("content must be string or array: %w", err)
	}

	structuredContent := make(StructuredContent, 0, len(chunks))
	for _, chunkData := range chunks {
		// Peek at the type field
		var typeCheck struct {
			Type string `json:"type"`
		}
		if err := json.Unmarshal(chunkData, &typeCheck); err != nil {
			return fmt.Errorf("failed to determine chunk type: %w", err)
		}

		var chunk ContentChunk
		switch typeCheck.Type {
		case "text":
			var tc TextChunk
			if err := json.Unmarshal(chunkData, &tc); err != nil {
				return err
			}
			chunk = tc
		case "image_url":
			var ic ImageChunk
			if err := json.Unmarshal(chunkData, &ic); err != nil {
				return err
			}
			chunk = ic
		case "file_url":
			var fc FileChunk
			if err := json.Unmarshal(chunkData, &fc); err != nil {
				return err
			}
			chunk = fc
		case "pdf_url":
			var pc PDFChunk
			if err := json.Unmarshal(chunkData, &pc); err != nil {
				return err
			}
			chunk = pc
		case "video_url":
			var vc VideoChunk
			if err := json.Unmarshal(chunkData, &vc); err != nil {
				return err
			}
			chunk = vc
		default:
			return fmt.Errorf("unknown content chunk type: %s", typeCheck.Type)
		}

		structuredContent = append(structuredContent, chunk)
	}

	m.Content = structuredContent
	return nil
}

// MarshalJSON implements custom marshaling for ChatMessage.
func (m ChatMessage) MarshalJSON() ([]byte, error) {
	type Alias ChatMessage
	return json.Marshal(struct {
		Alias
		Content interface{} `json:"content"`
	}{
		Alias:   Alias(m),
		Content: m.Content,
	})
}
