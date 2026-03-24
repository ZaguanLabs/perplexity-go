package browser

type SessionStatus string

const (
	SessionStatusRunning SessionStatus = "running"
	SessionStatusStopped SessionStatus = "stopped"
)

type SessionResponse struct {
	SessionID *string        `json:"session_id,omitempty"`
	Status    *SessionStatus `json:"status,omitempty"`
}
