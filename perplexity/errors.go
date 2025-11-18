package perplexity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Error is the base error type for all Perplexity API errors.
type Error struct {
	Message    string          `json:"message"`
	StatusCode int             `json:"status_code"`
	Body       json.RawMessage `json:"body,omitempty"`
	RequestID  string          `json:"request_id,omitempty"`
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.RequestID != "" {
		return fmt.Sprintf("perplexity: %s (status: %d, request_id: %s)", e.Message, e.StatusCode, e.RequestID)
	}
	return fmt.Sprintf("perplexity: %s (status: %d)", e.Message, e.StatusCode)
}

// BadRequestError represents a 400 Bad Request error.
type BadRequestError struct {
	Err *Error
}

func (e *BadRequestError) Error() string { return e.Err.Error() }
func (e *BadRequestError) Unwrap() error { return e.Err }

// AuthenticationError represents a 401 Unauthorized error.
type AuthenticationError struct {
	Err *Error
}

func (e *AuthenticationError) Error() string { return e.Err.Error() }
func (e *AuthenticationError) Unwrap() error { return e.Err }

// PermissionDeniedError represents a 403 Forbidden error.
type PermissionDeniedError struct {
	Err *Error
}

func (e *PermissionDeniedError) Error() string { return e.Err.Error() }
func (e *PermissionDeniedError) Unwrap() error { return e.Err }

// NotFoundError represents a 404 Not Found error.
type NotFoundError struct {
	Err *Error
}

func (e *NotFoundError) Error() string { return e.Err.Error() }
func (e *NotFoundError) Unwrap() error { return e.Err }

// ConflictError represents a 409 Conflict error.
type ConflictError struct {
	Err *Error
}

func (e *ConflictError) Error() string { return e.Err.Error() }
func (e *ConflictError) Unwrap() error { return e.Err }

// UnprocessableEntityError represents a 422 Unprocessable Entity error.
type UnprocessableEntityError struct {
	Err *Error
}

func (e *UnprocessableEntityError) Error() string { return e.Err.Error() }
func (e *UnprocessableEntityError) Unwrap() error { return e.Err }

// RateLimitError represents a 429 Too Many Requests error.
type RateLimitError struct {
	Err *Error
}

func (e *RateLimitError) Error() string { return e.Err.Error() }
func (e *RateLimitError) Unwrap() error { return e.Err }

// InternalServerError represents a 5xx server error.
type InternalServerError struct {
	Err *Error
}

func (e *InternalServerError) Error() string { return e.Err.Error() }
func (e *InternalServerError) Unwrap() error { return e.Err }

// ConnectionError represents a network connection error.
type ConnectionError struct {
	Err *Error
}

func (e *ConnectionError) Error() string { return e.Err.Error() }
func (e *ConnectionError) Unwrap() error { return e.Err }

// TimeoutError represents a request timeout error.
type TimeoutError struct {
	Err *Error
}

func (e *TimeoutError) Error() string { return e.Err.Error() }
func (e *TimeoutError) Unwrap() error { return e.Err }

// ValidationError represents a client-side validation error.
type ValidationError struct {
	Err *Error
}

func (e *ValidationError) Error() string { return e.Err.Error() }
func (e *ValidationError) Unwrap() error { return e.Err }

// newError creates a new Error from an HTTP response.
func newError(statusCode int, message string, body []byte, requestID string) error {
	baseErr := &Error{
		Message:    message,
		StatusCode: statusCode,
		Body:       body,
		RequestID:  requestID,
	}

	switch statusCode {
	case http.StatusBadRequest:
		return &BadRequestError{Err: baseErr}
	case http.StatusUnauthorized:
		return &AuthenticationError{Err: baseErr}
	case http.StatusForbidden:
		return &PermissionDeniedError{Err: baseErr}
	case http.StatusNotFound:
		return &NotFoundError{Err: baseErr}
	case http.StatusConflict:
		return &ConflictError{Err: baseErr}
	case http.StatusUnprocessableEntity:
		return &UnprocessableEntityError{Err: baseErr}
	case http.StatusTooManyRequests:
		return &RateLimitError{Err: baseErr}
	default:
		if statusCode >= 500 {
			return &InternalServerError{Err: baseErr}
		}
		return baseErr
	}
}

// IsRetryable returns true if the error is retryable.
func IsRetryable(err error) bool {
	switch err.(type) {
	case *RateLimitError, *InternalServerError, *ConflictError, *TimeoutError, *ConnectionError:
		return true
	default:
		return false
	}
}

// IsRateLimitError returns true if the error is a rate limit error.
func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	return ok
}

// IsAuthenticationError returns true if the error is an authentication error.
func IsAuthenticationError(err error) bool {
	_, ok := err.(*AuthenticationError)
	return ok
}

// IsTimeoutError returns true if the error is a timeout error.
func IsTimeoutError(err error) bool {
	_, ok := err.(*TimeoutError)
	return ok
}
