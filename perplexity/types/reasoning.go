package types

// ReasoningStep represents a step in the reasoning process.
type ReasoningStep struct {
	// Thought is the reasoning thought.
	Thought string `json:"thought"`

	// AgentProgress contains agent progress updates (optional).
	AgentProgress *ReasoningStepAgentProgress `json:"agent_progress,omitempty"`

	// BrowserAgent contains browser agent step summary (optional).
	BrowserAgent *ReasoningStepBrowserAgent `json:"browser_agent,omitempty"`

	// BrowserToolExecution contains tool input for browser automation (optional).
	BrowserToolExecution *ReasoningStepBrowserToolExecution `json:"browser_tool_execution,omitempty"`

	// ExecutePython contains code generation step details (optional).
	ExecutePython *ReasoningStepExecutePython `json:"execute_python,omitempty"`

	// FetchURLContent contains fetch URL content step details (optional).
	FetchURLContent *ReasoningStepFetchURLContent `json:"fetch_url_content,omitempty"`

	// FileAttachmentSearch contains file attachment search step details (optional).
	FileAttachmentSearch *ReasoningStepFileAttachmentSearch `json:"file_attachment_search,omitempty"`

	// Type is the type of reasoning step (optional).
	Type *string `json:"type,omitempty"`

	// WebSearch contains web search step details (optional).
	WebSearch *ReasoningStepWebSearch `json:"web_search,omitempty"`
}

// ReasoningStepAgentProgress represents agent progress for live-browsing updates.
type ReasoningStepAgentProgress struct {
	// Action is the current action (optional).
	Action *string `json:"action,omitempty"`

	// Screenshot is a screenshot URL or data (optional).
	Screenshot *string `json:"screenshot,omitempty"`

	// URL is the current URL (optional).
	URL *string `json:"url,omitempty"`
}

// ReasoningStepBrowserAgent represents a browser agent step summary.
type ReasoningStepBrowserAgent struct {
	// Result is the result of the browser action.
	Result string `json:"result"`

	// URL is the URL that was accessed.
	URL string `json:"url"`
}

// ReasoningStepBrowserToolExecution represents tool input for browser automation.
type ReasoningStepBrowserToolExecution struct {
	// Tool contains the tool configuration.
	Tool map[string]interface{} `json:"tool"`
}

// ReasoningStepExecutePython represents code generation step details.
type ReasoningStepExecutePython struct {
	// Code is the Python code that was executed.
	Code string `json:"code"`

	// Result is the result of the code execution.
	Result string `json:"result"`
}

// ReasoningStepFetchURLContent represents fetch URL content step details.
type ReasoningStepFetchURLContent struct {
	// Contents contains the fetched content.
	Contents []SearchResult `json:"contents"`
}

// ReasoningStepFileAttachmentSearch represents file attachment search step details.
type ReasoningStepFileAttachmentSearch struct {
	// AttachmentURLs contains URLs of the attachments.
	AttachmentURLs []string `json:"attachment_urls"`
}

// ReasoningStepWebSearch represents web search step details.
type ReasoningStepWebSearch struct {
	// SearchKeywords contains the search keywords used.
	SearchKeywords []string `json:"search_keywords"`

	// SearchResults contains the search results.
	SearchResults []SearchResult `json:"search_results"`
}
