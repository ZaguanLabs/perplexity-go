package responses

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Status string

type ReasoningEffort string

// WebSearchType selects the Responses web-search mode.
type WebSearchType string

const (
	WebSearchTypeWeb  WebSearchType = "web"
	WebSearchTypeFast WebSearchType = "fast"
)

type EventType string

type InputMessageRole string

type InputMessageType string

type InputMessageContentPartType string

type InputItemType string

type FunctionCallOutputPartType string

type FunctionCallOutputImageDetail string

type FunctionCallOutputStatus string

type ReasoningInputStatus string

type ReasoningSummaryType string

type ReasoningTextType string

type ToolType string

type SearchContextSize string

type OutputItemType string

type ResponseFormatType string

type ContentPartType string

type MessageOutputRole string

type SearchResultSource string

type SkillType string

type ProfileReferenceType string

const (
	StatusCompleted      Status = "completed"
	StatusFailed         Status = "failed"
	StatusInProgress     Status = "in_progress"
	StatusQueued         Status = "queued"
	StatusCancelled      Status = "cancelled"
	StatusRequiresAction Status = "requires_action"

	ReasoningEffortNone    ReasoningEffort = "none"
	ReasoningEffortMinimal ReasoningEffort = "minimal"
	ReasoningEffortLow     ReasoningEffort = "low"
	ReasoningEffortMedium  ReasoningEffort = "medium"
	ReasoningEffortHigh    ReasoningEffort = "high"
	ReasoningEffortXHigh   ReasoningEffort = "xhigh"

	EventTypeResponseCreated          EventType = "response.created"
	EventTypeResponseInProgress       EventType = "response.in_progress"
	EventTypeResponseCompleted        EventType = "response.completed"
	EventTypeResponseFailed           EventType = "response.failed"
	EventTypeOutputItemAdded          EventType = "response.output_item.added"
	EventTypeOutputItemDone           EventType = "response.output_item.done"
	EventTypeOutputTextDelta          EventType = "response.output_text.delta"
	EventTypeOutputTextDone           EventType = "response.output_text.done"
	EventTypeReasoningStarted         EventType = "response.reasoning.started"
	EventTypeReasoningSearchQueries   EventType = "response.reasoning.search_queries"
	EventTypeReasoningSearchResults   EventType = "response.reasoning.search_results"
	EventTypeReasoningFetchURLQueries EventType = "response.reasoning.fetch_url_queries"
	EventTypeReasoningFetchURLResults EventType = "response.reasoning.fetch_url_results"
	EventTypeReasoningStopped         EventType = "response.reasoning.stopped"
	EventTypeSkillLoaded              EventType = "response.skill.loaded"

	InputMessageRoleUser      InputMessageRole = "user"
	InputMessageRoleAssistant InputMessageRole = "assistant"
	InputMessageRoleSystem    InputMessageRole = "system"
	InputMessageRoleDeveloper InputMessageRole = "developer"

	InputMessageTypeMessage InputMessageType = "message"

	InputMessageContentPartTypeText  InputMessageContentPartType = "input_text"
	InputMessageContentPartTypeImage InputMessageContentPartType = "input_image"

	InputItemTypeMessage            InputItemType = "message"
	InputItemTypeFunctionCall       InputItemType = "function_call"
	InputItemTypeFunctionCallOutput InputItemType = "function_call_output"
	InputItemTypeReasoning          InputItemType = "reasoning"

	FunctionCallOutputPartTypeText  FunctionCallOutputPartType = "input_text"
	FunctionCallOutputPartTypeImage FunctionCallOutputPartType = "input_image"

	FunctionCallOutputImageDetailLow  FunctionCallOutputImageDetail = "low"
	FunctionCallOutputImageDetailHigh FunctionCallOutputImageDetail = "high"
	FunctionCallOutputImageDetailAuto FunctionCallOutputImageDetail = "auto"

	FunctionCallOutputStatusInProgress FunctionCallOutputStatus = "in_progress"
	FunctionCallOutputStatusCompleted  FunctionCallOutputStatus = "completed"
	FunctionCallOutputStatusIncomplete FunctionCallOutputStatus = "incomplete"

	ReasoningInputStatusInProgress ReasoningInputStatus = "in_progress"
	ReasoningInputStatusCompleted  ReasoningInputStatus = "completed"
	ReasoningInputStatusIncomplete ReasoningInputStatus = "incomplete"

	ReasoningSummaryTypeText ReasoningSummaryType = "summary_text"
	ReasoningTextTypeText    ReasoningTextType    = "reasoning_text"

	ToolTypeWebSearch     ToolType = "web_search"
	ToolTypeFetchURL      ToolType = "fetch_url"
	ToolTypeFunction      ToolType = "function"
	ToolTypePeopleSearch  ToolType = "people_search"
	ToolTypeFinanceSearch ToolType = "finance_search"
	ToolTypeSandbox       ToolType = "sandbox"
	ToolTypeMCP           ToolType = "mcp"
	ToolTypeConnector     ToolType = "connector"
	ToolTypeNamespace     ToolType = "namespace"

	SkillTypeBuiltin SkillType = "builtin"
	SkillTypeInline  SkillType = "inline"
	SkillTypeCustom  SkillType = "custom"

	ProfileReferenceTypeCustom ProfileReferenceType = "custom"

	SearchContextSizeLow    SearchContextSize = "low"
	SearchContextSizeMedium SearchContextSize = "medium"
	SearchContextSizeHigh   SearchContextSize = "high"

	OutputItemTypeMessage           OutputItemType = "message"
	OutputItemTypeSearchResults     OutputItemType = "search_results"
	OutputItemTypeFetchURLResults   OutputItemType = "fetch_url_results"
	OutputItemTypeFunctionCall      OutputItemType = "function_call"
	OutputItemTypeMCPListTools      OutputItemType = "mcp_list_tools"
	OutputItemTypeMCPCall           OutputItemType = "mcp_call"
	OutputItemTypeToolSearch        OutputItemType = "tool_search_output"
	OutputItemTypeSkillLoaded       OutputItemType = "skill_loaded"
	OutputItemTypeAdvisorResult     OutputItemType = "advisor_result"
	OutputItemTypeSandboxResults    OutputItemType = "sandbox_results"
	OutputItemTypeSandboxWriteFile  OutputItemType = "sandbox_write_file"
	OutputItemTypeSandboxReadFile   OutputItemType = "sandbox_read_file"
	OutputItemTypeSandboxEditFile   OutputItemType = "sandbox_edit_file"
	OutputItemTypeSandboxGrep       OutputItemType = "sandbox_grep"
	OutputItemTypeSandboxGlob       OutputItemType = "sandbox_glob"
	OutputItemTypeSandboxApplyPatch OutputItemType = "sandbox_apply_patch"
	OutputItemTypeShareFile         OutputItemType = "share_file"
	OutputItemTypeUnknown           OutputItemType = "unknown"

	ResponseFormatTypeJSONSchema ResponseFormatType = "json_schema"

	ContentPartTypeOutputText ContentPartType = "output_text"

	MessageOutputRoleAssistant MessageOutputRole = "assistant"

	SearchResultSourceWeb SearchResultSource = "web"
)

type InputMessageContentPart struct {
	FileData *string                     `json:"file_data,omitempty"`
	FileURL  *string                     `json:"file_url,omitempty"`
	Filename *string                     `json:"filename,omitempty"`
	Type     InputMessageContentPartType `json:"type"`
	ImageURL *string                     `json:"image_url,omitempty"`
	Text     *string                     `json:"text,omitempty"`
}

type InputMessageContent struct {
	Text  *string
	Parts []InputMessageContentPart
}

func (c InputMessageContent) MarshalJSON() ([]byte, error) {
	if c.Text != nil {
		return json.Marshal(*c.Text)
	}
	if c.Parts == nil {
		return []byte("null"), nil
	}
	return json.Marshal(c.Parts)
}

func (c *InputMessageContent) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*c = InputMessageContent{}
		return nil
	}
	var text string
	if err := json.Unmarshal(data, &text); err == nil {
		c.Text = &text
		c.Parts = nil
		return nil
	}
	var parts []InputMessageContentPart
	if err := json.Unmarshal(data, &parts); err != nil {
		return err
	}
	c.Text = nil
	c.Parts = parts
	return nil
}

type InputMessage struct {
	Content InputMessageContent `json:"content"`
	Role    InputMessageRole    `json:"role"`
	Type    InputMessageType    `json:"type,omitempty"`
}

type FunctionCallOutputInput struct {
	CallID           string                    `json:"call_id"`
	Output           any                       `json:"output"`
	Type             InputItemType             `json:"type"`
	ID               *string                   `json:"id,omitempty"`
	Name             *string                   `json:"name,omitempty"`
	Status           *FunctionCallOutputStatus `json:"status,omitempty"`
	ThoughtSignature *string                   `json:"thought_signature,omitempty"`
}

// FunctionCallOutputPart represents a text or image part returned by a tool.
// Text is used with input_text; ImageURL and Detail are used with input_image.
type FunctionCallOutputPart struct {
	Type     FunctionCallOutputPartType     `json:"type"`
	Text     *string                        `json:"text,omitempty"`
	ImageURL *string                        `json:"image_url,omitempty"`
	Detail   *FunctionCallOutputImageDetail `json:"detail,omitempty"`
}

type FunctionCallInput struct {
	Arguments        string        `json:"arguments"`
	CallID           string        `json:"call_id"`
	Name             string        `json:"name"`
	Type             InputItemType `json:"type"`
	ThoughtSignature *string       `json:"thought_signature,omitempty"`
}

type ReasoningSummaryInput struct {
	Text string               `json:"text"`
	Type ReasoningSummaryType `json:"type"`
}

type ReasoningTextInput struct {
	Text string            `json:"text"`
	Type ReasoningTextType `json:"type"`
}

type ReasoningInputItem struct {
	Summary          []ReasoningSummaryInput `json:"summary"`
	Type             InputItemType           `json:"type"`
	Content          []ReasoningTextInput    `json:"content,omitempty"`
	EncryptedContent *string                 `json:"encrypted_content,omitempty"`
	ID               *string                 `json:"id,omitempty"`
	Status           *ReasoningInputStatus   `json:"status,omitempty"`
}

type InputItem struct {
	value any
}

func NewInputItemFromMessage(v InputMessage) InputItem { return InputItem{value: v} }

func NewInputItemFromFunctionCallOutput(v FunctionCallOutputInput) InputItem {
	return InputItem{value: v}
}

func NewInputItemFromFunctionCall(v FunctionCallInput) InputItem { return InputItem{value: v} }

func NewInputItemFromReasoning(v ReasoningInputItem) InputItem { return InputItem{value: v} }

func (i InputItem) MarshalJSON() ([]byte, error) {
	return marshalUnionValue(i.value)
}

func (i *InputItem) UnmarshalJSON(data []byte) error {
	variant, err := unmarshalByTypeField(data, map[string]func() any{
		"":                                       func() any { return &InputMessage{} },
		string(InputItemTypeMCPListTools):        func() any { return &MCPListToolsInput{} },
		string(InputItemTypeMCPCall):             func() any { return &MCPCallOutputItem{} },
		string(InputItemTypeMCPApprovalRequest):  func() any { return &MCPApprovalRequest{} },
		string(InputItemTypeMCPApprovalResponse): func() any { return &MCPApprovalResponseInput{} },
		string(InputItemTypeMessage):             func() any { return &InputMessage{} },
		string(InputItemTypeFunctionCallOutput):  func() any { return &FunctionCallOutputInput{} },
		string(InputItemTypeFunctionCall):        func() any { return &FunctionCallInput{} },
		string(InputItemTypeReasoning):           func() any { return &ReasoningInputItem{} },
	})
	if err != nil {
		return err
	}
	i.value = dereferenceVariant(variant)
	return nil
}

func (i InputItem) AsInputMessage() (*InputMessage, bool) {
	v, ok := i.value.(InputMessage)
	if ok {
		return &v, true
	}
	return nil, false
}

func (i InputItem) AsFunctionCallOutput() (*FunctionCallOutputInput, bool) {
	v, ok := i.value.(FunctionCallOutputInput)
	if ok {
		return &v, true
	}
	return nil, false
}

func (i InputItem) AsFunctionCall() (*FunctionCallInput, bool) {
	v, ok := i.value.(FunctionCallInput)
	if ok {
		return &v, true
	}
	return nil, false
}

func (i InputItem) AsReasoning() (*ReasoningInputItem, bool) {
	v, ok := i.value.(ReasoningInputItem)
	if ok {
		return &v, true
	}
	return nil, false
}

type Input struct {
	Text  *string
	Items []InputItem
}

func (i Input) MarshalJSON() ([]byte, error) {
	if i.Text != nil {
		return json.Marshal(*i.Text)
	}
	if i.Items == nil {
		return []byte("null"), nil
	}
	return json.Marshal(i.Items)
}

func (i *Input) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*i = Input{}
		return nil
	}
	var text string
	if err := json.Unmarshal(data, &text); err == nil {
		i.Text = &text
		i.Items = nil
		return nil
	}
	var items []InputItem
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}
	i.Text = nil
	i.Items = items
	return nil
}

type Reasoning struct {
	Effort *ReasoningEffort `json:"effort,omitempty"`
}

const BuiltinSkillSearch = "search"

type BuiltinSkill struct {
	Name string    `json:"name"`
	Type SkillType `json:"type"`
}

type InlineSkill struct {
	Description  string    `json:"description"`
	Instructions string    `json:"instructions"`
	Name         string    `json:"name"`
	Type         SkillType `json:"type"`
}

type CustomSkill struct {
	ID      string    `json:"id"`
	Type    SkillType `json:"type"`
	Version *string   `json:"version,omitempty"`
}

type Skill struct {
	value any
}

func NewSkillFromBuiltin(v BuiltinSkill) Skill { return Skill{value: v} }

func NewSkillFromInline(v InlineSkill) Skill { return Skill{value: v} }

func NewSkillFromCustom(v CustomSkill) Skill { return Skill{value: v} }

func (s Skill) MarshalJSON() ([]byte, error) { return marshalUnionValue(s.value) }

func (s *Skill) UnmarshalJSON(data []byte) error {
	variant, err := unmarshalByTypeField(data, map[string]func() any{
		string(SkillTypeBuiltin): func() any { return &BuiltinSkill{} },
		string(SkillTypeInline):  func() any { return &InlineSkill{} },
		string(SkillTypeCustom):  func() any { return &CustomSkill{} },
	})
	if err != nil {
		return err
	}
	s.value = dereferenceVariant(variant)
	return nil
}

func (s Skill) AsBuiltin() (*BuiltinSkill, bool) {
	v, ok := s.value.(BuiltinSkill)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (s Skill) AsInline() (*InlineSkill, bool) {
	v, ok := s.value.(InlineSkill)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (s Skill) AsCustom() (*CustomSkill, bool) {
	v, ok := s.value.(CustomSkill)
	if !ok {
		return nil, false
	}
	return &v, true
}

type ProfileReference struct {
	ID      string               `json:"id"`
	Type    ProfileReferenceType `json:"type"`
	Version *string              `json:"version,omitempty"`
}

type WebSearchToolFilters struct {
	LastUpdatedAfterFilter  *string  `json:"last_updated_after_filter,omitempty"`
	LastUpdatedBeforeFilter *string  `json:"last_updated_before_filter,omitempty"`
	SearchAfterDateFilter   *string  `json:"search_after_date_filter,omitempty"`
	SearchBeforeDateFilter  *string  `json:"search_before_date_filter,omitempty"`
	SearchDomainFilter      []string `json:"search_domain_filter,omitempty"`
	SearchRecencyFilter     *string  `json:"search_recency_filter,omitempty"`
}

type UserLocation struct {
	City      *string  `json:"city,omitempty"`
	Country   *string  `json:"country,omitempty"`
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
	Region    *string  `json:"region,omitempty"`
}

type WebSearchTool struct {
	SearchType        *WebSearchType        `json:"search_type,omitempty"`
	Type              ToolType              `json:"type"`
	Filters           *WebSearchToolFilters `json:"filters,omitempty"`
	MaxTokens         *int                  `json:"max_tokens,omitempty"`
	MaxTokensPerPage  *int                  `json:"max_tokens_per_page,omitempty"`
	SearchContextSize *SearchContextSize    `json:"search_context_size,omitempty"`
	UserLocation      *UserLocation         `json:"user_location,omitempty"`
}

type FetchURLTool struct {
	Type    ToolType `json:"type"`
	MaxURLs *int     `json:"max_urls,omitempty"`
}

type PeopleSearchTool struct {
	Type ToolType `json:"type"`
}

type FinanceSearchTool struct {
	Type ToolType `json:"type"`
}

type SandboxTool struct {
	Type ToolType `json:"type"`
}

type MCPTool struct {
	RequireApproval *RequireApproval  `json:"require_approval,omitempty"`
	ServerLabel     string            `json:"server_label"`
	ServerURL       string            `json:"server_url"`
	Type            ToolType          `json:"type"`
	AllowedTools    []string          `json:"allowed_tools,omitempty"`
	Authorization   *string           `json:"authorization,omitempty"`
	DeferLoading    *bool             `json:"defer_loading,omitempty"`
	Headers         map[string]string `json:"headers,omitempty"`
}

type ConnectorTool struct {
	RequireApproval   *RequireApproval `json:"require_approval,omitempty"`
	ID                string           `json:"id"`
	ServerLabel       string           `json:"server_label"`
	Type              ToolType         `json:"type"`
	AllowedTools      []string         `json:"allowed_tools,omitempty"`
	ServerDescription *string          `json:"server_description,omitempty"`
}

type FunctionTool struct {
	Name        string         `json:"name"`
	Type        ToolType       `json:"type"`
	Description *string        `json:"description,omitempty"`
	Parameters  map[string]any `json:"parameters,omitempty"`
	Strict      *bool          `json:"strict,omitempty"`
}

type Tool struct {
	value any
}

func NewToolFromWebSearch(v WebSearchTool) Tool { return Tool{value: v} }

func NewToolFromFetchURL(v FetchURLTool) Tool { return Tool{value: v} }

func NewToolFromFunction(v FunctionTool) Tool { return Tool{value: v} }

func NewToolFromPeopleSearch(v PeopleSearchTool) Tool { return Tool{value: v} }

func NewToolFromFinanceSearch(v FinanceSearchTool) Tool { return Tool{value: v} }

func NewToolFromSandbox(v SandboxTool) Tool { return Tool{value: v} }

func NewToolFromMCP(v MCPTool) Tool { return Tool{value: v} }

func NewToolFromConnector(v ConnectorTool) Tool { return Tool{value: v} }

func (t Tool) MarshalJSON() ([]byte, error) {
	return marshalUnionValue(t.value)
}

func (t *Tool) UnmarshalJSON(data []byte) error {
	variant, err := unmarshalByTypeField(data, map[string]func() any{
		string(ToolTypeNamespace):     func() any { return &NamespaceTool{} },
		string(ToolTypeImageSearch):   func() any { return &ImageSearchTool{} },
		string(ToolTypeWebSearch):     func() any { return &WebSearchTool{} },
		string(ToolTypeFetchURL):      func() any { return &FetchURLTool{} },
		string(ToolTypeFunction):      func() any { return &FunctionTool{} },
		string(ToolTypePeopleSearch):  func() any { return &PeopleSearchTool{} },
		string(ToolTypeFinanceSearch): func() any { return &FinanceSearchTool{} },
		string(ToolTypeSandbox):       func() any { return &SandboxTool{} },
		string(ToolTypeMCP):           func() any { return &MCPTool{} },
		string(ToolTypeConnector):     func() any { return &ConnectorTool{} },
	})
	if err != nil {
		return err
	}
	t.value = dereferenceVariant(variant)
	return nil
}

func (t Tool) AsWebSearch() (*WebSearchTool, bool) {
	v, ok := t.value.(WebSearchTool)
	if ok {
		return &v, true
	}
	return nil, false
}

func (t Tool) AsFetchURL() (*FetchURLTool, bool) {
	v, ok := t.value.(FetchURLTool)
	if ok {
		return &v, true
	}
	return nil, false
}

func (t Tool) AsFunction() (*FunctionTool, bool) {
	v, ok := t.value.(FunctionTool)
	if ok {
		return &v, true
	}
	return nil, false
}

func (t Tool) AsPeopleSearch() (*PeopleSearchTool, bool) {
	v, ok := t.value.(PeopleSearchTool)
	if ok {
		return &v, true
	}
	return nil, false
}

func (t Tool) AsFinanceSearch() (*FinanceSearchTool, bool) {
	v, ok := t.value.(FinanceSearchTool)
	if ok {
		return &v, true
	}
	return nil, false
}

func (t Tool) AsSandbox() (*SandboxTool, bool) {
	v, ok := t.value.(SandboxTool)
	if ok {
		return &v, true
	}
	return nil, false
}

func (t Tool) AsMCP() (*MCPTool, bool) {
	v, ok := t.value.(MCPTool)
	return &v, ok
}

func (t Tool) AsConnector() (*ConnectorTool, bool) {
	v, ok := t.value.(ConnectorTool)
	return &v, ok
}

type JSONSchemaFormat struct {
	Name        string         `json:"name"`
	Schema      map[string]any `json:"schema"`
	Description *string        `json:"description,omitempty"`
	Strict      *bool          `json:"strict,omitempty"`
}

type ResponseFormat struct {
	Type       ResponseFormatType `json:"type"`
	JSONSchema *JSONSchemaFormat  `json:"json_schema,omitempty"`
}

type CreateParams struct {
	Effort             *string           `json:"effort,omitempty"`
	PromptCacheKey     *string           `json:"prompt_cache_key,omitempty"`
	ServiceTier        *string           `json:"service_tier,omitempty"`
	StepBudgetReminder *bool             `json:"step_budget_reminder,omitempty"`
	ToolChoice         any               `json:"tool_choice,omitempty"`
	Input              Input             `json:"input"`
	Background         *bool             `json:"background,omitempty"`
	Instructions       *string           `json:"instructions,omitempty"`
	LanguagePreference *string           `json:"language_preference,omitempty"`
	MaxOutputTokens    *int              `json:"max_output_tokens,omitempty"`
	MaxSteps           *int              `json:"max_steps,omitempty"`
	Model              *string           `json:"model,omitempty"`
	Models             []string          `json:"models,omitempty"`
	Preset             *string           `json:"preset,omitempty"`
	PreviousResponseID *string           `json:"previous_response_id,omitempty"`
	Profile            *ProfileReference `json:"profile,omitempty"`
	Reasoning          *Reasoning        `json:"reasoning,omitempty"`
	ResponseFormat     *ResponseFormat   `json:"response_format,omitempty"`
	Skills             []Skill           `json:"skills,omitempty"`
	Store              *bool             `json:"store,omitempty"`
	Stream             *bool             `json:"stream,omitempty"`
	Temperature        *float64          `json:"temperature,omitempty"`
	Tools              []Tool            `json:"tools,omitempty"`
	TopP               *float64          `json:"top_p,omitempty"`
}

type ErrorInfo struct {
	Message string  `json:"message"`
	Code    *string `json:"code,omitempty"`
	Type    *string `json:"type,omitempty"`
}

type Cost struct {
	Currency          string   `json:"currency"`
	InputCost         float64  `json:"input_cost"`
	OutputCost        float64  `json:"output_cost"`
	TotalCost         float64  `json:"total_cost"`
	CacheCreationCost *float64 `json:"cache_creation_cost,omitempty"`
	CacheReadCost     *float64 `json:"cache_read_cost,omitempty"`
	ToolCallsCost     *float64 `json:"tool_calls_cost,omitempty"`
}

type InputTokensDetails struct {
	CachedTokens             int  `json:"cached_tokens"`
	CacheCreationInputTokens *int `json:"cache_creation_input_tokens,omitempty"`
	CacheReadInputTokens     *int `json:"cache_read_input_tokens,omitempty"`
}

type ToolCallsDetails struct {
	CostUSD    *float64             `json:"cost_usd,omitempty"`
	TokenUsage []ToolCallTokenUsage `json:"token_usage,omitempty"`
	Invocation *int                 `json:"invocation,omitempty"`
}

type Usage struct {
	InputTokens        int                         `json:"input_tokens"`
	OutputTokens       int                         `json:"output_tokens"`
	TotalTokens        int                         `json:"total_tokens"`
	Cost               *Cost                       `json:"cost,omitempty"`
	InputTokensDetails *InputTokensDetails         `json:"input_tokens_details,omitempty"`
	ToolCallsDetails   map[string]ToolCallsDetails `json:"tool_calls_details,omitempty"`
}

type Annotation struct {
	EndIndex   *int    `json:"end_index,omitempty"`
	StartIndex *int    `json:"start_index,omitempty"`
	Title      *string `json:"title,omitempty"`
	Type       *string `json:"type,omitempty"`
	URL        *string `json:"url,omitempty"`
}

type ContentPart struct {
	Text        string          `json:"text"`
	Type        ContentPartType `json:"type"`
	Annotations []Annotation    `json:"annotations,omitempty"`
}

type SearchResult struct {
	ID          int                 `json:"id"`
	Snippet     string              `json:"snippet"`
	Title       string              `json:"title"`
	URL         string              `json:"url"`
	Date        *string             `json:"date,omitempty"`
	LastUpdated *string             `json:"last_updated,omitempty"`
	Source      *SearchResultSource `json:"source,omitempty"`
}

type FetchURLResultContent struct {
	Snippet string `json:"snippet"`
	Title   string `json:"title"`
	URL     string `json:"url"`
}

type MessageOutputItem struct {
	ID      string            `json:"id"`
	Content []ContentPart     `json:"content"`
	Role    MessageOutputRole `json:"role"`
	Status  Status            `json:"status"`
	Type    OutputItemType    `json:"type"`
}

type SearchResultsOutputItem struct {
	Results []SearchResult `json:"results"`
	Type    OutputItemType `json:"type"`
	Queries []string       `json:"queries,omitempty"`
}

type FetchURLResultsOutputItem struct {
	Contents []FetchURLResultContent `json:"contents"`
	Type     OutputItemType          `json:"type"`
}

type FunctionCallOutputItem struct {
	ID               string         `json:"id"`
	Arguments        string         `json:"arguments"`
	CallID           string         `json:"call_id"`
	Name             string         `json:"name"`
	Status           Status         `json:"status"`
	Type             OutputItemType `json:"type"`
	ThoughtSignature *string        `json:"thought_signature,omitempty"`
}

type MCPListToolsTool struct {
	Annotations *MCPToolDefAnnotations `json:"annotations,omitempty"`
	InputSchema map[string]any         `json:"input_schema"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
}

type MCPListToolsOutputItem struct {
	ID          string             `json:"id"`
	ServerLabel string             `json:"server_label"`
	Tools       []MCPListToolsTool `json:"tools"`
	Type        OutputItemType     `json:"type"`
	Error       *string            `json:"error,omitempty"`
	ConnectorID *string            `json:"connector_id,omitempty"`
}

type MCPCallOutputItem struct {
	ApprovalRequestID *string        `json:"approval_request_id,omitempty"`
	ID                string         `json:"id"`
	Arguments         string         `json:"arguments"`
	Name              string         `json:"name"`
	ServerLabel       string         `json:"server_label"`
	Type              OutputItemType `json:"type"`
	Error             *string        `json:"error,omitempty"`
	Output            *string        `json:"output,omitempty"`
	ConnectorID       *string        `json:"connector_id,omitempty"`
}

type NamespaceToolDef struct {
	Name        string         `json:"name"`
	Type        ToolType       `json:"type"`
	Description *string        `json:"description,omitempty"`
	Parameters  map[string]any `json:"parameters,omitempty"`
}

type NamespaceTool struct {
	Description string             `json:"description"`
	Name        string             `json:"name"`
	Tools       []NamespaceToolDef `json:"tools"`
	Type        ToolType           `json:"type"`
}

type ToolSearchOutputItem struct {
	CallID    *string         `json:"call_id"`
	Execution string          `json:"execution"`
	ID        string          `json:"id"`
	Status    string          `json:"status"`
	Tools     []NamespaceTool `json:"tools"`
	Type      OutputItemType  `json:"type"`
	Arguments *string         `json:"arguments,omitempty"`
}

type SkillLoadedOutputItem struct {
	Name string         `json:"name"`
	Type OutputItemType `json:"type"`
}

type AdvisorResultOutputItem struct {
	CallID       string         `json:"call_id"`
	Status       Status         `json:"status"`
	Type         OutputItemType `json:"type"`
	Advice       *string        `json:"advice,omitempty"`
	Arguments    *string        `json:"arguments,omitempty"`
	ErrorCode    *string        `json:"error_code,omitempty"`
	ErrorMessage *string        `json:"error_message,omitempty"`
	Question     *string        `json:"question,omitempty"`
}

type SandboxResult struct {
	DurationMS int    `json:"duration_ms"`
	ExitCode   int    `json:"exit_code"`
	Status     string `json:"status"`
	Stderr     string `json:"stderr"`
	Stdout     string `json:"stdout"`
}

type SandboxResultsOutputItem struct {
	CallID      string          `json:"call_id"`
	Code        string          `json:"code"`
	Language    string          `json:"language"`
	Results     []SandboxResult `json:"results"`
	Status      string          `json:"status"`
	Type        OutputItemType  `json:"type"`
	ContainerID *string         `json:"container_id,omitempty"`
}

type SandboxWriteFileOutputItem struct {
	CallID    string         `json:"call_id"`
	FilePath  string         `json:"file_path"`
	Type      OutputItemType `json:"type"`
	Error     *string        `json:"error,omitempty"`
	SizeBytes *int           `json:"size_bytes,omitempty"`
}

type SandboxReadFileOutputItem struct {
	CallID     string         `json:"call_id"`
	FilePath   string         `json:"file_path"`
	Type       OutputItemType `json:"type"`
	Content    *string        `json:"content,omitempty"`
	Error      *string        `json:"error,omitempty"`
	StartLine  *int           `json:"start_line,omitempty"`
	TotalLines *int           `json:"total_lines,omitempty"`
}

type SandboxEditFileOutputItem struct {
	CallID   string         `json:"call_id"`
	Type     OutputItemType `json:"type"`
	Error    *string        `json:"error,omitempty"`
	FilePath *string        `json:"file_path,omitempty"`
	Message  *string        `json:"message,omitempty"`
}

type SandboxGrepOutputItem struct {
	CallID    string         `json:"call_id"`
	Type      OutputItemType `json:"type"`
	Count     *int           `json:"count,omitempty"`
	Error     *string        `json:"error,omitempty"`
	Files     []string       `json:"files,omitempty"`
	Truncated *bool          `json:"truncated,omitempty"`
}

type SandboxGlobOutputItem struct {
	CallID    string         `json:"call_id"`
	Type      OutputItemType `json:"type"`
	Count     *int           `json:"count,omitempty"`
	Error     *string        `json:"error,omitempty"`
	Files     []string       `json:"files,omitempty"`
	Truncated *bool          `json:"truncated,omitempty"`
}

type SandboxApplyPatchOutputItem struct {
	CallID   string         `json:"call_id"`
	Type     OutputItemType `json:"type"`
	Added    []string       `json:"added,omitempty"`
	Deleted  []string       `json:"deleted,omitempty"`
	Error    *string        `json:"error,omitempty"`
	Modified []string       `json:"modified,omitempty"`
}

type ShareFileOutputItem struct {
	CallID    string         `json:"call_id"`
	Type      OutputItemType `json:"type"`
	Error     *string        `json:"error,omitempty"`
	FileID    *string        `json:"file_id,omitempty"`
	Filename  *string        `json:"filename,omitempty"`
	SizeBytes *int           `json:"size_bytes,omitempty"`
	URL       *string        `json:"url,omitempty"`
}

type UnknownOutputItem struct {
	ItemName string         `json:"item_name"`
	Payload  map[string]any `json:"payload"`
	Type     OutputItemType `json:"type"`
}

type OutputItem struct {
	value any
}

func NewOutputItemFromMessage(v MessageOutputItem) OutputItem { return OutputItem{value: v} }

func NewOutputItemFromSearchResults(v SearchResultsOutputItem) OutputItem {
	return OutputItem{value: v}
}

func NewOutputItemFromFetchURLResults(v FetchURLResultsOutputItem) OutputItem {
	return OutputItem{value: v}
}

func NewOutputItemFromFunctionCall(v FunctionCallOutputItem) OutputItem { return OutputItem{value: v} }

func NewOutputItemFromMCPListTools(v MCPListToolsOutputItem) OutputItem { return OutputItem{value: v} }
func NewOutputItemFromMCPCall(v MCPCallOutputItem) OutputItem           { return OutputItem{value: v} }
func NewOutputItemFromToolSearch(v ToolSearchOutputItem) OutputItem     { return OutputItem{value: v} }
func NewOutputItemFromSkillLoaded(v SkillLoadedOutputItem) OutputItem   { return OutputItem{value: v} }
func NewOutputItemFromAdvisorResult(v AdvisorResultOutputItem) OutputItem {
	return OutputItem{value: v}
}
func NewOutputItemFromSandboxResults(v SandboxResultsOutputItem) OutputItem {
	return OutputItem{value: v}
}
func NewOutputItemFromSandboxWriteFile(v SandboxWriteFileOutputItem) OutputItem {
	return OutputItem{value: v}
}
func NewOutputItemFromSandboxReadFile(v SandboxReadFileOutputItem) OutputItem {
	return OutputItem{value: v}
}
func NewOutputItemFromSandboxEditFile(v SandboxEditFileOutputItem) OutputItem {
	return OutputItem{value: v}
}
func NewOutputItemFromSandboxGrep(v SandboxGrepOutputItem) OutputItem { return OutputItem{value: v} }
func NewOutputItemFromSandboxGlob(v SandboxGlobOutputItem) OutputItem { return OutputItem{value: v} }
func NewOutputItemFromSandboxApplyPatch(v SandboxApplyPatchOutputItem) OutputItem {
	return OutputItem{value: v}
}
func NewOutputItemFromShareFile(v ShareFileOutputItem) OutputItem { return OutputItem{value: v} }
func NewOutputItemFromUnknown(v UnknownOutputItem) OutputItem     { return OutputItem{value: v} }

func (o OutputItem) MarshalJSON() ([]byte, error) {
	return marshalUnionValue(o.value)
}

func (o *OutputItem) UnmarshalJSON(data []byte) error {
	variant, err := unmarshalByTypeField(data, map[string]func() any{
		string(OutputItemTypeMCPApprovalRequest): func() any { return &MCPApprovalRequest{} },
		string(OutputItemTypeImageSearchResults): func() any { return &ImageSearchResultsOutputItem{} },
		string(OutputItemTypeMessage):            func() any { return &MessageOutputItem{} },
		string(OutputItemTypeSearchResults):      func() any { return &SearchResultsOutputItem{} },
		string(OutputItemTypeFetchURLResults):    func() any { return &FetchURLResultsOutputItem{} },
		string(OutputItemTypeFunctionCall):       func() any { return &FunctionCallOutputItem{} },
		string(OutputItemTypeMCPListTools):       func() any { return &MCPListToolsOutputItem{} },
		string(OutputItemTypeMCPCall):            func() any { return &MCPCallOutputItem{} },
		string(OutputItemTypeToolSearch):         func() any { return &ToolSearchOutputItem{} },
		string(OutputItemTypeSkillLoaded):        func() any { return &SkillLoadedOutputItem{} },
		string(OutputItemTypeAdvisorResult):      func() any { return &AdvisorResultOutputItem{} },
		string(OutputItemTypeSandboxResults):     func() any { return &SandboxResultsOutputItem{} },
		string(OutputItemTypeSandboxWriteFile):   func() any { return &SandboxWriteFileOutputItem{} },
		string(OutputItemTypeSandboxReadFile):    func() any { return &SandboxReadFileOutputItem{} },
		string(OutputItemTypeSandboxEditFile):    func() any { return &SandboxEditFileOutputItem{} },
		string(OutputItemTypeSandboxGrep):        func() any { return &SandboxGrepOutputItem{} },
		string(OutputItemTypeSandboxGlob):        func() any { return &SandboxGlobOutputItem{} },
		string(OutputItemTypeSandboxApplyPatch):  func() any { return &SandboxApplyPatchOutputItem{} },
		string(OutputItemTypeShareFile):          func() any { return &ShareFileOutputItem{} },
		string(OutputItemTypeUnknown):            func() any { return &UnknownOutputItem{} },
	})
	if err != nil {
		return err
	}
	o.value = dereferenceVariant(variant)
	return nil
}

func (o OutputItem) AsMessage() (*MessageOutputItem, bool) {
	v, ok := o.value.(MessageOutputItem)
	if ok {
		return &v, true
	}
	return nil, false
}

func (o OutputItem) AsSearchResults() (*SearchResultsOutputItem, bool) {
	v, ok := o.value.(SearchResultsOutputItem)
	if ok {
		return &v, true
	}
	return nil, false
}

func (o OutputItem) AsFetchURLResults() (*FetchURLResultsOutputItem, bool) {
	v, ok := o.value.(FetchURLResultsOutputItem)
	if ok {
		return &v, true
	}
	return nil, false
}

func (o OutputItem) AsFunctionCall() (*FunctionCallOutputItem, bool) {
	v, ok := o.value.(FunctionCallOutputItem)
	if ok {
		return &v, true
	}
	return nil, false
}

func (o OutputItem) AsMCPListTools() (*MCPListToolsOutputItem, bool) {
	v, ok := o.value.(MCPListToolsOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsMCPCall() (*MCPCallOutputItem, bool) {
	v, ok := o.value.(MCPCallOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsToolSearch() (*ToolSearchOutputItem, bool) {
	v, ok := o.value.(ToolSearchOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsSkillLoaded() (*SkillLoadedOutputItem, bool) {
	v, ok := o.value.(SkillLoadedOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsAdvisorResult() (*AdvisorResultOutputItem, bool) {
	v, ok := o.value.(AdvisorResultOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsSandboxResults() (*SandboxResultsOutputItem, bool) {
	v, ok := o.value.(SandboxResultsOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsSandboxWriteFile() (*SandboxWriteFileOutputItem, bool) {
	v, ok := o.value.(SandboxWriteFileOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsSandboxReadFile() (*SandboxReadFileOutputItem, bool) {
	v, ok := o.value.(SandboxReadFileOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsSandboxEditFile() (*SandboxEditFileOutputItem, bool) {
	v, ok := o.value.(SandboxEditFileOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsSandboxGrep() (*SandboxGrepOutputItem, bool) {
	v, ok := o.value.(SandboxGrepOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsSandboxGlob() (*SandboxGlobOutputItem, bool) {
	v, ok := o.value.(SandboxGlobOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsSandboxApplyPatch() (*SandboxApplyPatchOutputItem, bool) {
	v, ok := o.value.(SandboxApplyPatchOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsShareFile() (*ShareFileOutputItem, bool) {
	v, ok := o.value.(ShareFileOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func (o OutputItem) AsUnknown() (*UnknownOutputItem, bool) {
	v, ok := o.value.(UnknownOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

type CreateResponse struct {
	Effort             *string      `json:"effort,omitempty"`
	Preset             *string      `json:"preset,omitempty"`
	PromptCacheKey     *string      `json:"prompt_cache_key,omitempty"`
	ServiceTier        *string      `json:"service_tier,omitempty"`
	ToolChoice         any          `json:"tool_choice,omitempty"`
	ID                 string       `json:"id"`
	CreatedAt          int64        `json:"created_at"`
	Model              string       `json:"model"`
	Object             string       `json:"object"`
	Output             []OutputItem `json:"output"`
	Status             Status       `json:"status"`
	Background         *bool        `json:"background,omitempty"`
	Error              *ErrorInfo   `json:"error,omitempty"`
	PreviousResponseID *string      `json:"previous_response_id,omitempty"`
	Store              *bool        `json:"store,omitempty"`
	Usage              *Usage       `json:"usage,omitempty"`
}

type CancelResponse struct {
	ResponseID string `json:"response_id"`
	Status     string `json:"status"`
}

type ResponseFile struct {
	ID        string `json:"id"`
	Bytes     int    `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	Filename  string `json:"filename"`
	Object    string `json:"object"`
}

type ResponseFileList struct {
	Data   []ResponseFile `json:"data"`
	Object string         `json:"object"`
}

type RetrieveResponse = CreateResponse

func (r *CreateResponse) OutputText() string {
	if r == nil {
		return ""
	}
	var builder strings.Builder
	for _, output := range r.Output {
		message, ok := output.AsMessage()
		if !ok {
			continue
		}
		for _, content := range message.Content {
			if content.Type == ContentPartTypeOutputText {
				builder.WriteString(content.Text)
			}
		}
	}
	return builder.String()
}

type ResponseCreatedEvent struct {
	SequenceNumber int             `json:"sequence_number"`
	Type           EventType       `json:"type"`
	Response       *CreateResponse `json:"response,omitempty"`
}

type ResponseInProgressEvent struct {
	SequenceNumber int             `json:"sequence_number"`
	Type           EventType       `json:"type"`
	Response       *CreateResponse `json:"response,omitempty"`
}

type ResponseCompletedEvent struct {
	SequenceNumber int             `json:"sequence_number"`
	Type           EventType       `json:"type"`
	Response       *CreateResponse `json:"response,omitempty"`
}

type ResponseFailedEvent struct {
	Error          ErrorInfo `json:"error"`
	SequenceNumber int       `json:"sequence_number"`
	Type           EventType `json:"type"`
}

type OutputItemAddedEvent struct {
	Item           OutputItem `json:"item"`
	OutputIndex    int        `json:"output_index"`
	SequenceNumber int        `json:"sequence_number"`
	Type           EventType  `json:"type"`
}

type OutputItemDoneEvent struct {
	Item           OutputItem `json:"item"`
	OutputIndex    int        `json:"output_index"`
	SequenceNumber int        `json:"sequence_number"`
	Type           EventType  `json:"type"`
}

type TextDeltaEvent struct {
	ContentIndex   int       `json:"content_index"`
	Delta          string    `json:"delta"`
	ItemID         string    `json:"item_id"`
	OutputIndex    int       `json:"output_index"`
	SequenceNumber int       `json:"sequence_number"`
	Type           EventType `json:"type"`
}

type TextDoneEvent struct {
	ContentIndex   int       `json:"content_index"`
	ItemID         string    `json:"item_id"`
	OutputIndex    int       `json:"output_index"`
	SequenceNumber int       `json:"sequence_number"`
	Text           string    `json:"text"`
	Type           EventType `json:"type"`
}

type ReasoningStartedEvent struct {
	SequenceNumber int       `json:"sequence_number"`
	Type           EventType `json:"type"`
	Thought        *string   `json:"thought,omitempty"`
}

type SearchQueriesEvent struct {
	Queries        []string  `json:"queries"`
	SequenceNumber int       `json:"sequence_number"`
	Type           EventType `json:"type"`
	Thought        *string   `json:"thought,omitempty"`
}

type SearchResultsEvent struct {
	Results        []SearchResult `json:"results"`
	SequenceNumber int            `json:"sequence_number"`
	Type           EventType      `json:"type"`
	Thought        *string        `json:"thought,omitempty"`
	Usage          *Usage         `json:"usage,omitempty"`
}

type FetchURLQueriesEvent struct {
	SequenceNumber int       `json:"sequence_number"`
	Type           EventType `json:"type"`
	URLs           []string  `json:"urls"`
	Thought        *string   `json:"thought,omitempty"`
}

type FetchURLResultsEvent struct {
	Contents       []FetchURLResultContent `json:"contents"`
	SequenceNumber int                     `json:"sequence_number"`
	Type           EventType               `json:"type"`
	Thought        *string                 `json:"thought,omitempty"`
}

type ReasoningStoppedEvent struct {
	SequenceNumber int       `json:"sequence_number"`
	Type           EventType `json:"type"`
	Thought        *string   `json:"thought,omitempty"`
}

type SkillLoadedEvent struct {
	Name           string    `json:"name"`
	SequenceNumber int       `json:"sequence_number"`
	Type           EventType `json:"type"`
}

type StreamEvent struct {
	value any
}

func (e StreamEvent) MarshalJSON() ([]byte, error) {
	return marshalUnionValue(e.value)
}

func (e *StreamEvent) UnmarshalJSON(data []byte) error {
	variant, err := unmarshalByTypeField(data, map[string]func() any{
		string(EventTypeReasoningImageSearchResults): func() any { return &ImageSearchResultsEvent{} },
		string(EventTypeReasoningImageSearchQueries): func() any { return &ImageSearchQueriesEvent{} },
		string(EventTypeResponseCreated):             func() any { return &ResponseCreatedEvent{} },
		string(EventTypeResponseInProgress):          func() any { return &ResponseInProgressEvent{} },
		string(EventTypeResponseCompleted):           func() any { return &ResponseCompletedEvent{} },
		string(EventTypeResponseFailed):              func() any { return &ResponseFailedEvent{} },
		string(EventTypeOutputItemAdded):             func() any { return &OutputItemAddedEvent{} },
		string(EventTypeOutputItemDone):              func() any { return &OutputItemDoneEvent{} },
		string(EventTypeOutputTextDelta):             func() any { return &TextDeltaEvent{} },
		string(EventTypeOutputTextDone):              func() any { return &TextDoneEvent{} },
		string(EventTypeReasoningStarted):            func() any { return &ReasoningStartedEvent{} },
		string(EventTypeReasoningSearchQueries):      func() any { return &SearchQueriesEvent{} },
		string(EventTypeReasoningSearchResults):      func() any { return &SearchResultsEvent{} },
		string(EventTypeReasoningFetchURLQueries):    func() any { return &FetchURLQueriesEvent{} },
		string(EventTypeReasoningFetchURLResults):    func() any { return &FetchURLResultsEvent{} },
		string(EventTypeReasoningStopped):            func() any { return &ReasoningStoppedEvent{} },
		string(EventTypeSkillLoaded):                 func() any { return &SkillLoadedEvent{} },
	})
	if err != nil {
		return err
	}
	e.value = dereferenceVariant(variant)
	return nil
}

func (e StreamEvent) AsResponseCreated() (*ResponseCreatedEvent, bool) {
	v, ok := e.value.(ResponseCreatedEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsResponseInProgress() (*ResponseInProgressEvent, bool) {
	v, ok := e.value.(ResponseInProgressEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsResponseCompleted() (*ResponseCompletedEvent, bool) {
	v, ok := e.value.(ResponseCompletedEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsResponseFailed() (*ResponseFailedEvent, bool) {
	v, ok := e.value.(ResponseFailedEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsOutputItemAdded() (*OutputItemAddedEvent, bool) {
	v, ok := e.value.(OutputItemAddedEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsOutputItemDone() (*OutputItemDoneEvent, bool) {
	v, ok := e.value.(OutputItemDoneEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsTextDelta() (*TextDeltaEvent, bool) {
	v, ok := e.value.(TextDeltaEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsTextDone() (*TextDoneEvent, bool) {
	v, ok := e.value.(TextDoneEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsReasoningStarted() (*ReasoningStartedEvent, bool) {
	v, ok := e.value.(ReasoningStartedEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsSearchQueries() (*SearchQueriesEvent, bool) {
	v, ok := e.value.(SearchQueriesEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsSearchResults() (*SearchResultsEvent, bool) {
	v, ok := e.value.(SearchResultsEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsFetchURLQueries() (*FetchURLQueriesEvent, bool) {
	v, ok := e.value.(FetchURLQueriesEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsFetchURLResults() (*FetchURLResultsEvent, bool) {
	v, ok := e.value.(FetchURLResultsEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsReasoningStopped() (*ReasoningStoppedEvent, bool) {
	v, ok := e.value.(ReasoningStoppedEvent)
	if ok {
		return &v, true
	}
	return nil, false
}

func (e StreamEvent) AsSkillLoaded() (*SkillLoadedEvent, bool) {
	v, ok := e.value.(SkillLoadedEvent)
	if !ok {
		return nil, false
	}
	return &v, true
}

func marshalUnionValue(value any) ([]byte, error) {
	if value == nil {
		return []byte("null"), nil
	}
	return json.Marshal(value)
}

func unmarshalByTypeField(data []byte, mapping map[string]func() any) (any, error) {
	var envelope struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &envelope); err != nil {
		return nil, err
	}
	factory := mapping[envelope.Type]
	if factory == nil {
		return nil, fmt.Errorf("unknown discriminator type %q", envelope.Type)
	}
	variant := factory()
	if err := json.Unmarshal(data, variant); err != nil {
		return nil, err
	}
	return variant, nil
}

func dereferenceVariant(value any) any {
	switch v := value.(type) {
	case *NamespaceTool:
		return *v
	case *MCPListToolsInput:
		return *v
	case *MCPApprovalResponseInput:
		return *v
	case *MCPApprovalRequest:
		return *v
	case *ImageSearchTool:
		return *v
	case *ImageSearchResultsOutputItem:
		return *v
	case *ImageSearchResultsEvent:
		return *v
	case *ImageSearchQueriesEvent:
		return *v
	case *InputMessage:
		return *v
	case *FunctionCallOutputInput:
		return *v
	case *FunctionCallInput:
		return *v
	case *ReasoningInputItem:
		return *v
	case *WebSearchTool:
		return *v
	case *FetchURLTool:
		return *v
	case *FunctionTool:
		return *v
	case *PeopleSearchTool:
		return *v
	case *FinanceSearchTool:
		return *v
	case *SandboxTool:
		return *v
	case *MCPTool:
		return *v
	case *ConnectorTool:
		return *v
	case *BuiltinSkill:
		return *v
	case *InlineSkill:
		return *v
	case *CustomSkill:
		return *v
	case *MessageOutputItem:
		return *v
	case *SearchResultsOutputItem:
		return *v
	case *FetchURLResultsOutputItem:
		return *v
	case *FunctionCallOutputItem:
		return *v
	case *MCPListToolsOutputItem:
		return *v
	case *MCPCallOutputItem:
		return *v
	case *ToolSearchOutputItem:
		return *v
	case *SkillLoadedOutputItem:
		return *v
	case *AdvisorResultOutputItem:
		return *v
	case *SandboxResultsOutputItem:
		return *v
	case *SandboxWriteFileOutputItem:
		return *v
	case *SandboxReadFileOutputItem:
		return *v
	case *SandboxEditFileOutputItem:
		return *v
	case *SandboxGrepOutputItem:
		return *v
	case *SandboxGlobOutputItem:
		return *v
	case *SandboxApplyPatchOutputItem:
		return *v
	case *ShareFileOutputItem:
		return *v
	case *UnknownOutputItem:
		return *v
	case *ResponseCreatedEvent:
		return *v
	case *ResponseInProgressEvent:
		return *v
	case *ResponseCompletedEvent:
		return *v
	case *ResponseFailedEvent:
		return *v
	case *OutputItemAddedEvent:
		return *v
	case *OutputItemDoneEvent:
		return *v
	case *TextDeltaEvent:
		return *v
	case *TextDoneEvent:
		return *v
	case *ReasoningStartedEvent:
		return *v
	case *SearchQueriesEvent:
		return *v
	case *SearchResultsEvent:
		return *v
	case *FetchURLQueriesEvent:
		return *v
	case *FetchURLResultsEvent:
		return *v
	case *ReasoningStoppedEvent:
		return *v
	case *SkillLoadedEvent:
		return *v
	default:
		return value
	}
}

const (
	InputMessageContentPartTypeFile      InputMessageContentPartType = "input_file"
	InputItemTypeMCPApprovalResponse     InputItemType               = "mcp_approval_response"
	InputItemTypeMCPApprovalRequest      InputItemType               = "mcp_approval_request"
	InputItemTypeMCPCall                 InputItemType               = "mcp_call"
	InputItemTypeMCPListTools            InputItemType               = "mcp_list_tools"
	ToolTypeImageSearch                  ToolType                    = "image_search"
	OutputItemTypeImageSearchResults     OutputItemType              = "image_search_results"
	OutputItemTypeMCPApprovalRequest     OutputItemType              = "mcp_approval_request"
	EventTypeReasoningImageSearchQueries EventType                   = "response.reasoning.image_search_queries"
	EventTypeReasoningImageSearchResults EventType                   = "response.reasoning.image_search_results"
)

type ApprovalFilter struct {
	ReadOnly  *bool    `json:"read_only,omitempty"`
	ToolNames []string `json:"tool_names,omitempty"`
}
type ApprovalFilters struct {
	Always *ApprovalFilter `json:"always,omitempty"`
	Never  *ApprovalFilter `json:"never,omitempty"`
}

// RequireApproval is either the mode "always"/"never" or per-tool filters.
type RequireApproval struct {
	Mode    string
	Filters *ApprovalFilters
}

func (r RequireApproval) MarshalJSON() ([]byte, error) {
	if r.Filters != nil {
		return json.Marshal(r.Filters)
	}
	if r.Mode != "always" && r.Mode != "never" {
		return nil, fmt.Errorf("invalid approval mode %q", r.Mode)
	}
	return json.Marshal(r.Mode)
}
func (r *RequireApproval) UnmarshalJSON(data []byte) error {
	var mode string
	if err := json.Unmarshal(data, &mode); err == nil {
		if mode != "always" && mode != "never" {
			return fmt.Errorf("invalid approval mode %q", mode)
		}
		*r = RequireApproval{Mode: mode}
		return nil
	}
	var filters ApprovalFilters
	if len(data) == 0 || data[0] != '{' {
		return fmt.Errorf("approval must be a mode or filters")
	}
	if err := json.Unmarshal(data, &filters); err != nil {
		return err
	}
	*r = RequireApproval{Filters: &filters}
	return nil
}

type MCPToolDefAnnotations struct {
	ReadOnly *bool `json:"read_only,omitempty"`
}
type MCPApprovalRequest struct {
	Arguments        string  `json:"arguments"`
	ConnectorID      *string `json:"connector_id,omitempty"`
	ID               string  `json:"id"`
	Name             string  `json:"name"`
	ServerLabel      string  `json:"server_label"`
	ThoughtSignature *string `json:"thought_signature,omitempty"`
	Type             string  `json:"type"`
}
type MCPApprovalResponseInput struct {
	ApprovalRequestID string        `json:"approval_request_id"`
	Approve           bool          `json:"approve"`
	ID                *string       `json:"id,omitempty"`
	Reason            *string       `json:"reason,omitempty"`
	Type              InputItemType `json:"type"`
}
type MCPCallInput = MCPCallOutputItem
type MCPApprovalRequestInput = MCPApprovalRequest
type MCPApprovalRequestOutputItem = MCPApprovalRequest
type MCPListToolsInput struct {
	Error       *string            `json:"error,omitempty"`
	ID          *string            `json:"id,omitempty"`
	ServerLabel *string            `json:"server_label,omitempty"`
	Tools       []MCPListToolsTool `json:"tools,omitempty"`
	Type        InputItemType      `json:"type"`
}
type ImageResult struct {
	Height    int     `json:"height"`
	ImageURL  string  `json:"image_url"`
	OriginURL string  `json:"origin_url"`
	Title     *string `json:"title,omitempty"`
	Width     int     `json:"width"`
}
type ImageSearchFilters struct {
	DomainFilter []string `json:"domain_filter,omitempty"`
	FormatFilter []string `json:"format_filter,omitempty"`
	SafeSearch   *bool    `json:"safe_search,omitempty"`
}
type ImageSearchTool struct {
	Filters    *ImageSearchFilters `json:"filters,omitempty"`
	MaxResults *int                `json:"max_results,omitempty"`
	Type       ToolType            `json:"type"`
}
type ImageSearchResultsOutputItem struct {
	Error   *string        `json:"error,omitempty"`
	Queries []string       `json:"queries,omitempty"`
	Results []ImageResult  `json:"results"`
	Type    OutputItemType `json:"type"`
}
type ImageSearchQueriesEvent struct {
	CallID         string    `json:"call_id"`
	Queries        []string  `json:"queries"`
	SequenceNumber int       `json:"sequence_number"`
	Thought        *string   `json:"thought,omitempty"`
	Type           EventType `json:"type"`
}
type ImageSearchResultsEvent struct {
	CallID         string        `json:"call_id"`
	Results        []ImageResult `json:"results"`
	SequenceNumber int           `json:"sequence_number"`
	Thought        *string       `json:"thought,omitempty"`
	Type           EventType     `json:"type"`
	Usage          *Usage        `json:"usage,omitempty"`
}

func NewInputItemFromMCPApprovalResponse(v MCPApprovalResponseInput) InputItem {
	return InputItem{value: v}
}
func (i InputItem) AsMCPApprovalResponse() (*MCPApprovalResponseInput, bool) {
	v, ok := i.value.(MCPApprovalResponseInput)
	if !ok {
		return nil, false
	}
	return &v, true
}

func NewInputItemFromMCPApprovalRequest(v MCPApprovalRequest) InputItem { return InputItem{value: v} }
func (i InputItem) AsMCPApprovalRequest() (*MCPApprovalRequest, bool) {
	v, ok := i.value.(MCPApprovalRequest)
	if !ok {
		return nil, false
	}
	return &v, true
}

func NewInputItemFromMCPCall(v MCPCallOutputItem) InputItem { return InputItem{value: v} }
func (i InputItem) AsMCPCall() (*MCPCallOutputItem, bool) {
	v, ok := i.value.(MCPCallOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func NewInputItemFromMCPListTools(v MCPListToolsInput) InputItem { return InputItem{value: v} }
func (i InputItem) AsMCPListTools() (*MCPListToolsInput, bool) {
	v, ok := i.value.(MCPListToolsInput)
	if !ok {
		return nil, false
	}
	return &v, true
}

func NewToolFromImageSearch(v ImageSearchTool) Tool { return Tool{value: v} }
func (t Tool) AsImageSearch() (*ImageSearchTool, bool) {
	v, ok := t.value.(ImageSearchTool)
	if !ok {
		return nil, false
	}
	return &v, true
}

func NewToolFromNamespace(v NamespaceTool) Tool { return Tool{value: v} }
func (t Tool) AsNamespace() (*NamespaceTool, bool) {
	v, ok := t.value.(NamespaceTool)
	if !ok {
		return nil, false
	}
	return &v, true
}

func NewOutputItemFromImageSearchResults(v ImageSearchResultsOutputItem) OutputItem {
	return OutputItem{value: v}
}
func (o OutputItem) AsImageSearchResults() (*ImageSearchResultsOutputItem, bool) {
	v, ok := o.value.(ImageSearchResultsOutputItem)
	if !ok {
		return nil, false
	}
	return &v, true
}

func NewOutputItemFromMCPApprovalRequest(v MCPApprovalRequest) OutputItem {
	return OutputItem{value: v}
}
func (o OutputItem) AsMCPApprovalRequest() (*MCPApprovalRequest, bool) {
	v, ok := o.value.(MCPApprovalRequest)
	if !ok {
		return nil, false
	}
	return &v, true
}

func NewStreamEventFromImageSearchQueries(v ImageSearchQueriesEvent) StreamEvent {
	return StreamEvent{value: v}
}
func (e StreamEvent) AsImageSearchQueries() (*ImageSearchQueriesEvent, bool) {
	v, ok := e.value.(ImageSearchQueriesEvent)
	if !ok {
		return nil, false
	}
	return &v, true
}

func NewStreamEventFromImageSearchResults(v ImageSearchResultsEvent) StreamEvent {
	return StreamEvent{value: v}
}
func (e StreamEvent) AsImageSearchResults() (*ImageSearchResultsEvent, bool) {
	v, ok := e.value.(ImageSearchResultsEvent)
	if !ok {
		return nil, false
	}
	return &v, true
}

type OutputTokensDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}
type ToolCallTokenUsage struct {
	InputTokens         int                 `json:"input_tokens"`
	InputTokensDetails  InputTokensDetails  `json:"input_tokens_details"`
	Model               *string             `json:"model,omitempty"`
	OutputTokens        int                 `json:"output_tokens"`
	OutputTokensDetails OutputTokensDetails `json:"output_tokens_details"`
}
