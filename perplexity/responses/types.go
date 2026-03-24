package responses

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Status string

type ReasoningEffort string

type EventType string

type InputMessageRole string

type InputMessageType string

type InputMessageContentPartType string

type InputItemType string

type ToolType string

type OutputItemType string

type ResponseFormatType string

type ContentPartType string

type MessageOutputRole string

type SearchResultSource string

const (
	StatusCompleted      Status = "completed"
	StatusFailed         Status = "failed"
	StatusInProgress     Status = "in_progress"
	StatusRequiresAction Status = "requires_action"

	ReasoningEffortLow    ReasoningEffort = "low"
	ReasoningEffortMedium ReasoningEffort = "medium"
	ReasoningEffortHigh   ReasoningEffort = "high"

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

	ToolTypeWebSearch ToolType = "web_search"
	ToolTypeFetchURL  ToolType = "fetch_url"
	ToolTypeFunction  ToolType = "function"

	OutputItemTypeMessage         OutputItemType = "message"
	OutputItemTypeSearchResults   OutputItemType = "search_results"
	OutputItemTypeFetchURLResults OutputItemType = "fetch_url_results"
	OutputItemTypeFunctionCall    OutputItemType = "function_call"

	ResponseFormatTypeJSONSchema ResponseFormatType = "json_schema"

	ContentPartTypeOutputText ContentPartType = "output_text"

	MessageOutputRoleAssistant MessageOutputRole = "assistant"

	SearchResultSourceWeb SearchResultSource = "web"
)

type InputMessageContentPart struct {
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
	Type    InputMessageType    `json:"type"`
}

type FunctionCallOutputInput struct {
	CallID           string        `json:"call_id"`
	Output           string        `json:"output"`
	Type             InputItemType `json:"type"`
	Name             *string       `json:"name,omitempty"`
	ThoughtSignature *string       `json:"thought_signature,omitempty"`
}

type FunctionCallInput struct {
	Arguments        string        `json:"arguments"`
	CallID           string        `json:"call_id"`
	Name             string        `json:"name"`
	Type             InputItemType `json:"type"`
	ThoughtSignature *string       `json:"thought_signature,omitempty"`
}

type InputItem struct {
	value any
}

func NewInputItemFromMessage(v InputMessage) InputItem { return InputItem{value: v} }

func NewInputItemFromFunctionCallOutput(v FunctionCallOutputInput) InputItem {
	return InputItem{value: v}
}

func NewInputItemFromFunctionCall(v FunctionCallInput) InputItem { return InputItem{value: v} }

func (i InputItem) MarshalJSON() ([]byte, error) {
	return marshalUnionValue(i.value)
}

func (i *InputItem) UnmarshalJSON(data []byte) error {
	variant, err := unmarshalByTypeField(data, map[string]func() any{
		string(InputItemTypeMessage):            func() any { return &InputMessage{} },
		string(InputItemTypeFunctionCallOutput): func() any { return &FunctionCallOutputInput{} },
		string(InputItemTypeFunctionCall):       func() any { return &FunctionCallInput{} },
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
	Type             ToolType              `json:"type"`
	Filters          *WebSearchToolFilters `json:"filters,omitempty"`
	MaxTokens        *int                  `json:"max_tokens,omitempty"`
	MaxTokensPerPage *int                  `json:"max_tokens_per_page,omitempty"`
	UserLocation     *UserLocation         `json:"user_location,omitempty"`
}

type FetchURLTool struct {
	Type    ToolType `json:"type"`
	MaxURLs *int     `json:"max_urls,omitempty"`
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

func (t Tool) MarshalJSON() ([]byte, error) {
	return marshalUnionValue(t.value)
}

func (t *Tool) UnmarshalJSON(data []byte) error {
	variant, err := unmarshalByTypeField(data, map[string]func() any{
		string(ToolTypeWebSearch): func() any { return &WebSearchTool{} },
		string(ToolTypeFetchURL):  func() any { return &FetchURLTool{} },
		string(ToolTypeFunction):  func() any { return &FunctionTool{} },
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
	Input              Input           `json:"input"`
	Instructions       *string         `json:"instructions,omitempty"`
	LanguagePreference *string         `json:"language_preference,omitempty"`
	MaxOutputTokens    *int            `json:"max_output_tokens,omitempty"`
	MaxSteps           *int            `json:"max_steps,omitempty"`
	Model              *string         `json:"model,omitempty"`
	Models             []string        `json:"models,omitempty"`
	Preset             *string         `json:"preset,omitempty"`
	Reasoning          *Reasoning      `json:"reasoning,omitempty"`
	ResponseFormat     *ResponseFormat `json:"response_format,omitempty"`
	Stream             *bool           `json:"stream,omitempty"`
	Tools              []Tool          `json:"tools,omitempty"`
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
	CacheCreationInputTokens *int `json:"cache_creation_input_tokens,omitempty"`
	CacheReadInputTokens     *int `json:"cache_read_input_tokens,omitempty"`
}

type ToolCallsDetails struct {
	Invocation *int `json:"invocation,omitempty"`
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

func (o OutputItem) MarshalJSON() ([]byte, error) {
	return marshalUnionValue(o.value)
}

func (o *OutputItem) UnmarshalJSON(data []byte) error {
	variant, err := unmarshalByTypeField(data, map[string]func() any{
		string(OutputItemTypeMessage):         func() any { return &MessageOutputItem{} },
		string(OutputItemTypeSearchResults):   func() any { return &SearchResultsOutputItem{} },
		string(OutputItemTypeFetchURLResults): func() any { return &FetchURLResultsOutputItem{} },
		string(OutputItemTypeFunctionCall):    func() any { return &FunctionCallOutputItem{} },
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

type CreateResponse struct {
	ID        string       `json:"id"`
	CreatedAt int64        `json:"created_at"`
	Model     string       `json:"model"`
	Object    string       `json:"object"`
	Output    []OutputItem `json:"output"`
	Status    Status       `json:"status"`
	Error     *ErrorInfo   `json:"error,omitempty"`
	Usage     *Usage       `json:"usage,omitempty"`
}

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

type StreamEvent struct {
	value any
}

func (e StreamEvent) MarshalJSON() ([]byte, error) {
	return marshalUnionValue(e.value)
}

func (e *StreamEvent) UnmarshalJSON(data []byte) error {
	variant, err := unmarshalByTypeField(data, map[string]func() any{
		string(EventTypeResponseCreated):          func() any { return &ResponseCreatedEvent{} },
		string(EventTypeResponseInProgress):       func() any { return &ResponseInProgressEvent{} },
		string(EventTypeResponseCompleted):        func() any { return &ResponseCompletedEvent{} },
		string(EventTypeResponseFailed):           func() any { return &ResponseFailedEvent{} },
		string(EventTypeOutputItemAdded):          func() any { return &OutputItemAddedEvent{} },
		string(EventTypeOutputItemDone):           func() any { return &OutputItemDoneEvent{} },
		string(EventTypeOutputTextDelta):          func() any { return &TextDeltaEvent{} },
		string(EventTypeOutputTextDone):           func() any { return &TextDoneEvent{} },
		string(EventTypeReasoningStarted):         func() any { return &ReasoningStartedEvent{} },
		string(EventTypeReasoningSearchQueries):   func() any { return &SearchQueriesEvent{} },
		string(EventTypeReasoningSearchResults):   func() any { return &SearchResultsEvent{} },
		string(EventTypeReasoningFetchURLQueries): func() any { return &FetchURLQueriesEvent{} },
		string(EventTypeReasoningFetchURLResults): func() any { return &FetchURLResultsEvent{} },
		string(EventTypeReasoningStopped):         func() any { return &ReasoningStoppedEvent{} },
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
	case *InputMessage:
		return *v
	case *FunctionCallOutputInput:
		return *v
	case *FunctionCallInput:
		return *v
	case *WebSearchTool:
		return *v
	case *FetchURLTool:
		return *v
	case *FunctionTool:
		return *v
	case *MessageOutputItem:
		return *v
	case *SearchResultsOutputItem:
		return *v
	case *FetchURLResultsOutputItem:
		return *v
	case *FunctionCallOutputItem:
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
	default:
		return value
	}
}
