# Type Reference from Python SDK

This document catalogs the key types from the Python SDK (v0.20.0) for accurate Go implementation.

## Core Response Types

### UsageInfo
```python
class UsageInfo(BaseModel):
    completion_tokens: int
    cost: Cost
    prompt_tokens: int
    total_tokens: int
    citation_tokens: Optional[int] = None
    num_search_queries: Optional[int] = None
    reasoning_tokens: Optional[int] = None
    search_context_size: Optional[str] = None

class Cost(BaseModel):
    input_tokens_cost: float
    output_tokens_cost: float
    total_cost: float
    citation_tokens_cost: Optional[float] = None
    reasoning_tokens_cost: Optional[float] = None
    request_cost: Optional[float] = None
    search_queries_cost: Optional[float] = None
```

### Choice
```python
class Choice(BaseModel):
    delta: ChatMessageOutput
    index: int
    message: ChatMessageOutput
    finish_reason: Optional[Literal["stop", "length"]] = None
```

### APIPublicSearchResult
```python
class APIPublicSearchResult(BaseModel):
    title: str
    url: str
    date: Optional[str] = None
    last_updated: Optional[str] = None
    snippet: Optional[str] = None
    source: Optional[Literal["web", "attachment"]] = None
```

## Message Types

### ChatMessageOutput (and ChatMessageInput - same structure)
```python
class ChatMessageOutput(BaseModel):
    content: Union[str, List[ContentStructuredContent]]
    role: Literal["system", "user", "assistant", "tool"]
    reasoning_steps: Optional[List[ReasoningStep]] = None
    tool_calls: Optional[List[ToolCall]] = None
```

### Content Chunks (Union Type)
```python
ContentStructuredContent: TypeAlias = Union[
    ContentStructuredContentChatMessageContentTextChunk,
    ContentStructuredContentChatMessageContentImageChunk,
    ContentStructuredContentChatMessageContentFileChunk,
    ContentStructuredContentChatMessageContentPdfChunk,
    ContentStructuredContentChatMessageContentVideoChunk,
]

# Text Chunk
class ContentStructuredContentChatMessageContentTextChunk(BaseModel):
    text: str
    type: Literal["text"]

# Image Chunk
class ContentStructuredContentChatMessageContentImageChunk(BaseModel):
    image_url: Union[ImageURLObject, str]  # Can be string or object
    type: Literal["image_url"]

# File Chunk
class ContentStructuredContentChatMessageContentFileChunk(BaseModel):
    file_url: Union[FileURLObject, str]
    type: Literal["file_url"]
    file_name: Optional[str] = None

# PDF Chunk
class ContentStructuredContentChatMessageContentPdfChunk(BaseModel):
    pdf_url: Union[PdfURLObject, str]
    type: Literal["pdf_url"]

# Video Chunk
class ContentStructuredContentChatMessageContentVideoChunk(BaseModel):
    type: Literal["video_url"]
    video_url: Union[VideoURLObject, str]
```

### ReasoningStep
```python
class ReasoningStep(BaseModel):
    thought: str
    agent_progress: Optional[ReasoningStepAgentProgress] = None
    browser_agent: Optional[ReasoningStepBrowserAgent] = None
    browser_tool_execution: Optional[ReasoningStepBrowserToolExecution] = None
    execute_python: Optional[ReasoningStepExecutePython] = None
    fetch_url_content: Optional[ReasoningStepFetchURLContent] = None
    file_attachment_search: Optional[ReasoningStepFileAttachmentSearch] = None
    type: Optional[str] = None
    web_search: Optional[ReasoningStepWebSearch] = None
```

### ToolCall
```python
class ToolCall(BaseModel):
    id: Optional[str] = None
    function: Optional[ToolCallFunction] = None
    type: Optional[Literal["function"]] = None

class ToolCallFunction(BaseModel):
    arguments: Optional[str] = None
    name: Optional[str] = None
```

## Request Parameter Types

### CompletionCreateParams (60+ fields)
Key required fields:
- `messages: Required[Iterable[ChatMessageInput]]`
- `model: Required[str]`

Optional fields include:
- `stream: Optional[Literal[False]]` or `Literal[True]`
- `max_tokens: Optional[int]`
- `temperature: Optional[float]`
- `top_p: Optional[float]`
- `frequency_penalty: Optional[float]`
- `presence_penalty: Optional[float]`
- `stop: Union[str, SequenceNotStr[str], None]`
- `tools: Optional[Iterable[Tool]]`
- `tool_choice: Optional[Literal["none", "auto", "required"]]`
- `web_search_options: WebSearchOptions`
- `search_domain_filter: Optional[SequenceNotStr[str]]`
- `search_recency_filter: Optional[Literal["hour", "day", "week", "month", "year"]]`
- `search_mode: Optional[Literal["web", "academic", "sec"]]`
- `return_images: Optional[bool]`
- `return_related_questions: Optional[bool]`
- `reasoning_effort: Optional[Literal["minimal", "low", "medium", "high"]]`
- ... and 40+ more

### Tool Definition
```python
class Tool(TypedDict, total=False):
    function: Required[ToolFunction]
    type: Required[Literal["function"]]

class ToolFunction(TypedDict, total=False):
    name: Required[str]
    description: str
    parameters: ToolFunctionParameters
```

### WebSearchOptions
```python
class WebSearchOptions(TypedDict, total=False):
    user_location: WebSearchOptionsUserLocation

class WebSearchOptionsUserLocation(TypedDict, total=False):
    latitude: Required[float]
    longitude: Required[float]
```

### ResponseFormat (Union Type)
```python
ResponseFormat: TypeAlias = Union[
    ResponseFormatResponseFormatText,
    ResponseFormatResponseFormatJsonSchema,
    ResponseFormatResponseFormatRegex,
]

class ResponseFormatResponseFormatText(TypedDict, total=False):
    type: Required[Literal["text"]]

class ResponseFormatResponseFormatJsonSchema(TypedDict, total=False):
    json_schema: Required[ResponseFormatResponseFormatJsonSchemaJsonSchema]
    type: Required[Literal["json_schema"]]

class ResponseFormatResponseFormatRegex(TypedDict, total=False):
    regex: Required[ResponseFormatResponseFormatRegexRegex]
    type: Required[Literal["regex"]]
```

## Search API Types

### SearchCreateParams
```python
class SearchCreateParams(TypedDict, total=False):
    query: Required[Union[str, SequenceNotStr[str]]]
    country: Optional[str]
    display_server_time: bool
    max_results: int
    max_tokens: int
    max_tokens_per_page: int
    search_after_date_filter: Optional[str]
    search_before_date_filter: Optional[str]
    search_domain_filter: Optional[SequenceNotStr[str]]
    search_language_filter: Optional[SequenceNotStr[str]]
    search_mode: Optional[Literal["web", "academic", "sec"]]
    search_recency_filter: Optional[Literal["hour", "day", "week", "month", "year"]]
```

### SearchCreateResponse
```python
class SearchCreateResponse(BaseModel):
    id: str
    results: List[Result]
    server_time: Optional[str] = None

class Result(BaseModel):
    snippet: str
    title: str
    url: str
    date: Optional[str] = None
    last_updated: Optional[str] = None
```

## StreamChunk
```python
class StreamChunk(BaseModel):
    id: str
    choices: List[Choice]
    created: int
    model: str
    citations: Optional[List[str]] = None
    object: Optional[str] = None
    search_results: Optional[List[APIPublicSearchResult]] = None
    status: Optional[Literal["PENDING", "COMPLETED"]] = None
    type: Optional[Literal["message", "info", "end_of_stream"]] = None
    usage: Optional[UsageInfo] = None
```

## Key Observations for Go Implementation

1. **Union Types**: Python uses `Union[str, List[...]]` extensively
   - Go approach: Use interface{} with type assertions or custom unmarshaling

2. **Optional Fields**: Python uses `Optional[T]` 
   - Go approach: Use pointers `*T`

3. **Literal Types**: Python uses `Literal["value1", "value2"]`
   - Go approach: Use string constants with validation

4. **Nested Structures**: Deep nesting in content chunks
   - Go approach: Proper struct composition with json tags

5. **Type Aliases**: Python uses `TypeAlias` for union types
   - Go approach: Use interfaces with type switches

6. **Required vs Optional**: Python uses `Required[T]` in TypedDict
   - Go approach: Non-pointer for required, pointer for optional

7. **Cost Tracking**: UsageInfo includes detailed cost breakdown
   - Important for billing/monitoring features
