package types

// ToolCall represents a tool call made by the assistant.
type ToolCall struct {
	// ID is the unique identifier for the tool call (optional).
	ID *string `json:"id,omitempty"`

	// Function contains the function call details (optional).
	Function *ToolCallFunction `json:"function,omitempty"`

	// Type is the type of tool call (optional).
	Type *ToolCallType `json:"type,omitempty"`
}

// ToolCallType represents the type of tool call.
type ToolCallType string

const (
	// ToolCallTypeFunction indicates a function call.
	ToolCallTypeFunction ToolCallType = "function"
)

// ToolCallFunction represents a function call.
type ToolCallFunction struct {
	// Arguments contains the function arguments as a JSON string (optional).
	Arguments *string `json:"arguments,omitempty"`

	// Name is the name of the function (optional).
	Name *string `json:"name,omitempty"`
}

// Tool represents a tool that can be called by the model.
type Tool struct {
	// Type is the type of tool.
	Type ToolType `json:"type"`

	// Function contains the function definition.
	Function ToolFunction `json:"function"`
}

// ToolType represents the type of tool.
type ToolType string

const (
	// ToolTypeFunction indicates a function tool.
	ToolTypeFunction ToolType = "function"
)

// ToolFunction represents a function tool definition.
type ToolFunction struct {
	// Name is the name of the function.
	Name string `json:"name"`

	// Description is a description of what the function does (optional).
	Description *string `json:"description,omitempty"`

	// Parameters contains the function parameters schema (optional).
	Parameters map[string]interface{} `json:"parameters,omitempty"`
}

// ToolChoice represents how the model should choose tools.
type ToolChoice string

const (
	// ToolChoiceNone means the model will not call any tools.
	ToolChoiceNone ToolChoice = "none"

	// ToolChoiceAuto means the model can choose whether to call tools.
	ToolChoiceAuto ToolChoice = "auto"

	// ToolChoiceRequired means the model must call at least one tool.
	ToolChoiceRequired ToolChoice = "required"
)
