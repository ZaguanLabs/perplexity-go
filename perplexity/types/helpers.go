package types

// Helper functions for creating pointers to primitive types.
// These are useful for setting optional fields in request parameters.

// String returns a pointer to the given string value.
func String(v string) *string {
	return &v
}

// Int returns a pointer to the given int value.
func Int(v int) *int {
	return &v
}

// Int64 returns a pointer to the given int64 value.
func Int64(v int64) *int64 {
	return &v
}

// Float64 returns a pointer to the given float64 value.
func Float64(v float64) *float64 {
	return &v
}

// Bool returns a pointer to the given bool value.
func Bool(v bool) *bool {
	return &v
}

// Role helpers for common roles.

// SystemMessage creates a system message with text content.
func SystemMessage(content string) ChatMessage {
	return ChatMessage{
		Role:    RoleSystem,
		Content: TextContent(content),
	}
}

// UserMessage creates a user message with text content.
func UserMessage(content string) ChatMessage {
	return ChatMessage{
		Role:    RoleUser,
		Content: TextContent(content),
	}
}

// AssistantMessage creates an assistant message with text content.
func AssistantMessage(content string) ChatMessage {
	return ChatMessage{
		Role:    RoleAssistant,
		Content: TextContent(content),
	}
}

// ToolMessage creates a tool message with text content.
func ToolMessage(content string) ChatMessage {
	return ChatMessage{
		Role:    RoleTool,
		Content: TextContent(content),
	}
}
