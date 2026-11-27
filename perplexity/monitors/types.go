// Code generated from the Perplexity Python SDK 0.50.0 schema. DO NOT EDIT.

package monitors

import "encoding/json"
import "fmt"

type CreateMonitorRequest struct {
	Goal     *string                `json:"goal,omitempty"`
	Metadata *MonitorMetadata       `json:"metadata,omitempty"`
	Name     string                 `json:"name"`
	Schedule MonitorScheduleRequest `json:"schedule"`
	Targets  MonitorTargets         `json:"targets"`
	Webhook  *MonitorWebhookRequest `json:"webhook,omitempty"`
}
type DeleteMonitorResponse struct {
	Deleted bool   `json:"deleted"`
	ID      string `json:"id"`
	Object  string `json:"object"`
}
type EventMonitor struct {
	ID          string              `json:"id"`
	Metadata    map[string]string   `json:"metadata"`
	Name        string              `json:"name"`
	PauseReason *MonitorPauseReason `json:"pause_reason,omitempty"`
	Status      MonitorStatus       `json:"status"`
}
type Monitor struct {
	CreatedAt    MonitorTimestamp    `json:"created_at"`
	CurrentRunID *string             `json:"current_run_id,omitempty"`
	Goal         string              `json:"goal"`
	ID           string              `json:"id"`
	LastRun      *MonitorLastRun     `json:"last_run,omitempty"`
	Metadata     MonitorMetadata     `json:"metadata"`
	Name         string              `json:"name"`
	NextRunAt    *MonitorTimestamp   `json:"next_run_at,omitempty"`
	Object       string              `json:"object"`
	PauseReason  *MonitorPauseReason `json:"pause_reason,omitempty"`
	Schedule     MonitorSchedule     `json:"schedule"`
	Status       MonitorStatus       `json:"status"`
	Targets      []MonitorTarget     `json:"targets"`
	UpdatedAt    MonitorTimestamp    `json:"updated_at"`
	Webhook      *MonitorWebhook     `json:"webhook,omitempty"`
}
type MonitorChange struct {
	SearchResultChange *MonitorSearchResultChange
	PageChange         *MonitorPageChange
	Unknown            json.RawMessage
}

func (v MonitorChange) MarshalJSON() ([]byte, error) {
	var active any
	count := 0
	if v.SearchResultChange != nil {
		active = v.SearchResultChange
		count++
	}
	if v.PageChange != nil {
		active = v.PageChange
		count++
	}
	if v.Unknown != nil {
		active = v.Unknown
		count++
	}
	if count != 1 {
		return nil, fmt.Errorf("MonitorChange requires one variant")
	}
	return json.Marshal(active)
}
func (v *MonitorChange) UnmarshalJSON(raw []byte) error {
	var tag struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &tag); err != nil {
		return err
	}
	var decoded MonitorChange
	switch tag.Type {
	case "search_result":
		if err := json.Unmarshal(raw, &decoded.SearchResultChange); err != nil {
			return err
		}
	case "page_change":
		if err := json.Unmarshal(raw, &decoded.PageChange); err != nil {
			return err
		}
	default:
		decoded.Unknown = append(json.RawMessage(nil), raw...)
	}
	*v = decoded
	return nil
}

type MonitorChangeId = string
type MonitorChangesDetectedEvent struct {
	Data      MonitorRunEventData `json:"data"`
	ID        MonitorEventId      `json:"id"`
	Object    MonitorEventObject  `json:"object"`
	Timestamp MonitorTimestamp    `json:"timestamp"`
	Type      string              `json:"type"`
}
type MonitorCronSchedule struct {
	Cron     string `json:"cron"`
	Timezone string `json:"timezone"`
	Type     string `json:"type"`
}
type MonitorCronScheduleRequest struct {
	Cron     string  `json:"cron"`
	Timezone *string `json:"timezone,omitempty"`
	Type     string  `json:"type"`
}
type MonitorCurrency = string
type MonitorDiff struct {
	AddedLines   int    `json:"added_lines"`
	RemovedLines int    `json:"removed_lines"`
	Truncated    bool   `json:"truncated"`
	Unified      string `json:"unified"`
}
type MonitorDomain = string
type MonitorEvent struct {
	ChangesDetectedEvent *MonitorChangesDetectedEvent
	RunCompletedEvent    *MonitorRunCompletedEvent
	PausedEvent          *MonitorPausedEvent
	WebhookTestEvent     *MonitorWebhookTestEvent
	Unknown              json.RawMessage
}

func (v MonitorEvent) MarshalJSON() ([]byte, error) {
	var active any
	count := 0
	if v.ChangesDetectedEvent != nil {
		active = v.ChangesDetectedEvent
		count++
	}
	if v.RunCompletedEvent != nil {
		active = v.RunCompletedEvent
		count++
	}
	if v.PausedEvent != nil {
		active = v.PausedEvent
		count++
	}
	if v.WebhookTestEvent != nil {
		active = v.WebhookTestEvent
		count++
	}
	if v.Unknown != nil {
		active = v.Unknown
		count++
	}
	if count != 1 {
		return nil, fmt.Errorf("MonitorEvent requires one variant")
	}
	return json.Marshal(active)
}
func (v *MonitorEvent) UnmarshalJSON(raw []byte) error {
	var tag struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &tag); err != nil {
		return err
	}
	var decoded MonitorEvent
	switch tag.Type {
	case "monitor.changes.detected":
		if err := json.Unmarshal(raw, &decoded.ChangesDetectedEvent); err != nil {
			return err
		}
	case "monitor.run.completed":
		if err := json.Unmarshal(raw, &decoded.RunCompletedEvent); err != nil {
			return err
		}
	case "monitor.paused":
		if err := json.Unmarshal(raw, &decoded.PausedEvent); err != nil {
			return err
		}
	case "monitor.webhook.test":
		if err := json.Unmarshal(raw, &decoded.WebhookTestEvent); err != nil {
			return err
		}
	default:
		decoded.Unknown = append(json.RawMessage(nil), raw...)
	}
	*v = decoded
	return nil
}

type MonitorEventId = string
type MonitorEventObject = string
type MonitorEventType = string
type MonitorFreshness = string
type MonitorIntervalSchedule struct {
	Interval string `json:"interval"`
	Type     string `json:"type"`
}
type MonitorIntervalScheduleRequest struct {
	Interval string `json:"interval"`
	Type     string `json:"type"`
}
type MonitorLastRun struct {
	ChangeCount int              `json:"change_count"`
	FinishedAt  MonitorTimestamp `json:"finished_at"`
	ID          string           `json:"id"`
	Status      MonitorRunStatus `json:"status"`
}
type MonitorList struct {
	Data    []Monitor `json:"data"`
	HasMore bool      `json:"has_more"`
	Object  string    `json:"object"`
}
type MonitorMetadata map[string]any
type MonitorPageChange struct {
	Confidence        float64          `json:"confidence"`
	Diff              MonitorDiff      `json:"diff"`
	FetchedAt         MonitorTimestamp `json:"fetched_at"`
	ID                MonitorChangeId  `json:"id"`
	PreviousFetchedAt MonitorTimestamp `json:"previous_fetched_at"`
	Title             *string          `json:"title,omitempty"`
	Type              string           `json:"type"`
	URL               string           `json:"url"`
}
type MonitorPageTarget struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
type MonitorPageTargetError struct {
	Code       MonitorTargetErrorCode `json:"code"`
	Message    string                 `json:"message"`
	StatusCode *int                   `json:"status_code,omitempty"`
	Type       string                 `json:"type"`
	URL        string                 `json:"url"`
}
type MonitorPageTargetRequest struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
type MonitorPauseReason = string
type MonitorPausedEvent struct {
	Data      MonitorPausedEventData `json:"data"`
	ID        MonitorEventId         `json:"id"`
	Object    MonitorEventObject     `json:"object"`
	Timestamp MonitorTimestamp       `json:"timestamp"`
	Type      string                 `json:"type"`
}
type MonitorPausedEventData struct {
	Monitor PausedEventMonitor `json:"monitor"`
	Run     *MonitorRun        `json:"run,omitempty"`
}
type MonitorRun struct {
	Changes      []MonitorChange      `json:"changes"`
	Counts       MonitorRunCounts     `json:"counts"`
	CreatedAt    MonitorTimestamp     `json:"created_at"`
	Error        *MonitorRunError     `json:"error,omitempty"`
	FinishedAt   *MonitorTimestamp    `json:"finished_at,omitempty"`
	ID           string               `json:"id"`
	MonitorID    string               `json:"monitor_id"`
	Object       string               `json:"object"`
	Status       MonitorRunStatus     `json:"status"`
	TargetErrors []MonitorTargetError `json:"target_errors"`
	Trigger      MonitorRunTrigger    `json:"trigger"`
	Usage        MonitorRunUsage      `json:"usage"`
	Window       *MonitorRunWindow    `json:"window,omitempty"`
}
type MonitorRunCompletedEvent struct {
	Data      MonitorRunEventData `json:"data"`
	ID        MonitorEventId      `json:"id"`
	Object    MonitorEventObject  `json:"object"`
	Timestamp MonitorTimestamp    `json:"timestamp"`
	Type      string              `json:"type"`
}
type MonitorRunCost struct {
	Currency  MonitorCurrency `json:"currency"`
	TotalCost float64         `json:"total_cost"`
}
type MonitorRunCounts struct {
	Candidates int `json:"candidates"`
	Changes    int `json:"changes"`
	Errors     int `json:"errors"`
}
type MonitorRunError struct {
	Code       *MonitorRunFailureCode `json:"code,omitempty"`
	Message    string                 `json:"message"`
	SkipReason *MonitorPauseReason    `json:"skip_reason,omitempty"`
}
type MonitorRunEventData struct {
	Monitor EventMonitor `json:"monitor"`
	Run     MonitorRun   `json:"run"`
}
type MonitorRunFailureCode = string
type MonitorRunList struct {
	Data    []MonitorRun `json:"data"`
	HasMore bool         `json:"has_more"`
	Object  string       `json:"object"`
}
type MonitorRunStatus = string
type MonitorRunTrigger = string
type MonitorRunUsage struct {
	Cost        MonitorRunCost `json:"cost"`
	FetchedURLs int            `json:"fetched_urls"`
}
type MonitorRunWindow struct {
	From MonitorTimestamp `json:"from"`
	To   MonitorTimestamp `json:"to"`
}
type MonitorSchedule struct {
	IntervalSchedule *MonitorIntervalSchedule
	CronSchedule     *MonitorCronSchedule
	Unknown          json.RawMessage
}

func (v MonitorSchedule) MarshalJSON() ([]byte, error) {
	var active any
	count := 0
	if v.IntervalSchedule != nil {
		active = v.IntervalSchedule
		count++
	}
	if v.CronSchedule != nil {
		active = v.CronSchedule
		count++
	}
	if v.Unknown != nil {
		active = v.Unknown
		count++
	}
	if count != 1 {
		return nil, fmt.Errorf("MonitorSchedule requires one variant")
	}
	return json.Marshal(active)
}
func (v *MonitorSchedule) UnmarshalJSON(raw []byte) error {
	var tag struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &tag); err != nil {
		return err
	}
	var decoded MonitorSchedule
	switch tag.Type {
	case "interval":
		if err := json.Unmarshal(raw, &decoded.IntervalSchedule); err != nil {
			return err
		}
	case "cron":
		if err := json.Unmarshal(raw, &decoded.CronSchedule); err != nil {
			return err
		}
	default:
		decoded.Unknown = append(json.RawMessage(nil), raw...)
	}
	*v = decoded
	return nil
}

type MonitorScheduleRequest struct {
	IntervalScheduleRequest *MonitorIntervalScheduleRequest
	CronScheduleRequest     *MonitorCronScheduleRequest
	Unknown                 json.RawMessage
}

func (v MonitorScheduleRequest) MarshalJSON() ([]byte, error) {
	var active any
	count := 0
	if v.IntervalScheduleRequest != nil {
		active = v.IntervalScheduleRequest
		count++
	}
	if v.CronScheduleRequest != nil {
		active = v.CronScheduleRequest
		count++
	}
	if v.Unknown != nil {
		active = v.Unknown
		count++
	}
	if count != 1 {
		return nil, fmt.Errorf("MonitorScheduleRequest requires one variant")
	}
	return json.Marshal(active)
}
func (v *MonitorScheduleRequest) UnmarshalJSON(raw []byte) error {
	var tag struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &tag); err != nil {
		return err
	}
	var decoded MonitorScheduleRequest
	switch tag.Type {
	case "interval":
		if err := json.Unmarshal(raw, &decoded.IntervalScheduleRequest); err != nil {
			return err
		}
	case "cron":
		if err := json.Unmarshal(raw, &decoded.CronScheduleRequest); err != nil {
			return err
		}
	default:
		decoded.Unknown = append(json.RawMessage(nil), raw...)
	}
	*v = decoded
	return nil
}

type MonitorSearchFilters struct {
	ExcludeDomains []MonitorDomain    `json:"exclude_domains"`
	Freshness      []MonitorFreshness `json:"freshness"`
	IncludeDomains []MonitorDomain    `json:"include_domains"`
}
type MonitorSearchFiltersRequest struct {
	ExcludeDomains *[]string           `json:"exclude_domains,omitempty"`
	Freshness      *[]MonitorFreshness `json:"freshness,omitempty"`
	IncludeDomains *[]string           `json:"include_domains,omitempty"`
}
type MonitorSearchResultChange struct {
	Confidence    float64           `json:"confidence"`
	ID            MonitorChangeId   `json:"id"`
	LastUpdatedAt *MonitorTimestamp `json:"last_updated_at,omitempty"`
	PublishedAt   *MonitorTimestamp `json:"published_at,omitempty"`
	Snippet       *string           `json:"snippet,omitempty"`
	Title         *string           `json:"title,omitempty"`
	Type          string            `json:"type"`
	URL           string            `json:"url"`
}
type MonitorSearchTarget struct {
	Filters    MonitorSearchFilters `json:"filters"`
	MaxResults int                  `json:"max_results"`
	Query      string               `json:"query"`
	Type       string               `json:"type"`
}
type MonitorSearchTargetRequest struct {
	Filters    *MonitorSearchFiltersRequest `json:"filters,omitempty"`
	MaxResults *int                         `json:"max_results,omitempty"`
	Query      string                       `json:"query"`
	Type       string                       `json:"type"`
}
type MonitorStatus = string
type MonitorTarget struct {
	SearchTarget *MonitorSearchTarget
	PageTarget   *MonitorPageTarget
	Unknown      json.RawMessage
}

func (v MonitorTarget) MarshalJSON() ([]byte, error) {
	var active any
	count := 0
	if v.SearchTarget != nil {
		active = v.SearchTarget
		count++
	}
	if v.PageTarget != nil {
		active = v.PageTarget
		count++
	}
	if v.Unknown != nil {
		active = v.Unknown
		count++
	}
	if count != 1 {
		return nil, fmt.Errorf("MonitorTarget requires one variant")
	}
	return json.Marshal(active)
}
func (v *MonitorTarget) UnmarshalJSON(raw []byte) error {
	var tag struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &tag); err != nil {
		return err
	}
	var decoded MonitorTarget
	switch tag.Type {
	case "search":
		if err := json.Unmarshal(raw, &decoded.SearchTarget); err != nil {
			return err
		}
	case "page":
		if err := json.Unmarshal(raw, &decoded.PageTarget); err != nil {
			return err
		}
	default:
		decoded.Unknown = append(json.RawMessage(nil), raw...)
	}
	*v = decoded
	return nil
}

type MonitorTargetError = MonitorPageTargetError
type MonitorTargetErrorCode = string
type MonitorTargetRequest struct {
	SearchTargetRequest *MonitorSearchTargetRequest
	PageTargetRequest   *MonitorPageTargetRequest
	Unknown             json.RawMessage
}

func (v MonitorTargetRequest) MarshalJSON() ([]byte, error) {
	var active any
	count := 0
	if v.SearchTargetRequest != nil {
		active = v.SearchTargetRequest
		count++
	}
	if v.PageTargetRequest != nil {
		active = v.PageTargetRequest
		count++
	}
	if v.Unknown != nil {
		active = v.Unknown
		count++
	}
	if count != 1 {
		return nil, fmt.Errorf("MonitorTargetRequest requires one variant")
	}
	return json.Marshal(active)
}
func (v *MonitorTargetRequest) UnmarshalJSON(raw []byte) error {
	var tag struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(raw, &tag); err != nil {
		return err
	}
	var decoded MonitorTargetRequest
	switch tag.Type {
	case "search":
		if err := json.Unmarshal(raw, &decoded.SearchTargetRequest); err != nil {
			return err
		}
	case "page":
		if err := json.Unmarshal(raw, &decoded.PageTargetRequest); err != nil {
			return err
		}
	default:
		decoded.Unknown = append(json.RawMessage(nil), raw...)
	}
	*v = decoded
	return nil
}

type MonitorTargets = []MonitorTargetRequest
type MonitorTimestamp = string
type MonitorWebhook struct {
	Events []MonitorEventType    `json:"events"`
	Secret *MonitorWebhookSecret `json:"secret,omitempty"`
	URL    string                `json:"url"`
}
type MonitorWebhookRequest = MonitorWebhookRequestVariant1
type MonitorWebhookRequestVariant1 struct {
	Events *[]MonitorEventType   `json:"events,omitempty"`
	Secret *MonitorWebhookSecret `json:"secret,omitempty"`
	URL    string                `json:"url"`
}
type MonitorWebhookSecret = string
type MonitorWebhookTest struct {
	Delivered  bool                     `json:"delivered"`
	DurationMS int                      `json:"duration_ms"`
	Error      *MonitorWebhookTestError `json:"error,omitempty"`
	EventID    string                   `json:"event_id"`
	Object     string                   `json:"object"`
	StatusCode *int                     `json:"status_code,omitempty"`
}
type MonitorWebhookTestError = string
type MonitorWebhookTestEvent struct {
	Data      MonitorWebhookTestEventData `json:"data"`
	ID        MonitorEventId              `json:"id"`
	Object    MonitorEventObject          `json:"object"`
	Timestamp MonitorTimestamp            `json:"timestamp"`
	Type      string                      `json:"type"`
}
type MonitorWebhookTestEventData struct {
	Monitor EventMonitor `json:"monitor"`
}
type PausedEventMonitor struct {
	ID          string             `json:"id"`
	Metadata    map[string]string  `json:"metadata"`
	Name        string             `json:"name"`
	PauseReason MonitorPauseReason `json:"pause_reason"`
	Status      string             `json:"status"`
}
type UpdateMonitorRequest struct {
	Goal     **string                `json:"goal,omitempty"`
	Metadata *map[string]string      `json:"metadata,omitempty"`
	Name     *string                 `json:"name,omitempty"`
	Schedule *MonitorScheduleRequest `json:"schedule,omitempty"`
	Targets  *MonitorTargets         `json:"targets,omitempty"`
	Webhook  **MonitorWebhookRequest `json:"webhook,omitempty"`
}
