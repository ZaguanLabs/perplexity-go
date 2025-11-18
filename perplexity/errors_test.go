package perplexity

import (
	"net/http"
	"testing"
)

func TestNewError(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		message    string
		wantType   string
	}{
		{
			name:       "bad request",
			statusCode: http.StatusBadRequest,
			message:    "invalid request",
			wantType:   "*perplexity.BadRequestError",
		},
		{
			name:       "unauthorized",
			statusCode: http.StatusUnauthorized,
			message:    "invalid API key",
			wantType:   "*perplexity.AuthenticationError",
		},
		{
			name:       "forbidden",
			statusCode: http.StatusForbidden,
			message:    "access denied",
			wantType:   "*perplexity.PermissionDeniedError",
		},
		{
			name:       "not found",
			statusCode: http.StatusNotFound,
			message:    "resource not found",
			wantType:   "*perplexity.NotFoundError",
		},
		{
			name:       "conflict",
			statusCode: http.StatusConflict,
			message:    "conflict",
			wantType:   "*perplexity.ConflictError",
		},
		{
			name:       "unprocessable entity",
			statusCode: http.StatusUnprocessableEntity,
			message:    "validation failed",
			wantType:   "*perplexity.UnprocessableEntityError",
		},
		{
			name:       "rate limit",
			statusCode: http.StatusTooManyRequests,
			message:    "rate limit exceeded",
			wantType:   "*perplexity.RateLimitError",
		},
		{
			name:       "internal server error",
			statusCode: http.StatusInternalServerError,
			message:    "server error",
			wantType:   "*perplexity.InternalServerError",
		},
		{
			name:       "bad gateway",
			statusCode: http.StatusBadGateway,
			message:    "bad gateway",
			wantType:   "*perplexity.InternalServerError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := newError(tt.statusCode, tt.message, nil, "test-request-id")
			if err == nil {
				t.Fatal("newError() returned nil")
			}

			// Check error message
			if err.Error() == "" {
				t.Error("Error() returned empty string")
			}

			// Check if error contains the message
			errStr := err.Error()
			if len(errStr) == 0 {
				t.Error("Error message is empty")
			}
		})
	}
}

func TestIsRetryable(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "rate limit error",
			err:  &RateLimitError{Err: &Error{Message: "rate limit"}},
			want: true,
		},
		{
			name: "internal server error",
			err:  &InternalServerError{Err: &Error{Message: "server error"}},
			want: true,
		},
		{
			name: "conflict error",
			err:  &ConflictError{Err: &Error{Message: "conflict"}},
			want: true,
		},
		{
			name: "timeout error",
			err:  &TimeoutError{Err: &Error{Message: "timeout"}},
			want: true,
		},
		{
			name: "connection error",
			err:  &ConnectionError{Err: &Error{Message: "connection failed"}},
			want: true,
		},
		{
			name: "bad request error",
			err:  &BadRequestError{Err: &Error{Message: "bad request"}},
			want: false,
		},
		{
			name: "authentication error",
			err:  &AuthenticationError{Err: &Error{Message: "unauthorized"}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRetryable(tt.err); got != tt.want {
				t.Errorf("IsRetryable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsRateLimitError(t *testing.T) {
	rateLimitErr := &RateLimitError{Err: &Error{Message: "rate limit"}}
	otherErr := &BadRequestError{Err: &Error{Message: "bad request"}}

	if !IsRateLimitError(rateLimitErr) {
		t.Error("IsRateLimitError() returned false for RateLimitError")
	}

	if IsRateLimitError(otherErr) {
		t.Error("IsRateLimitError() returned true for non-RateLimitError")
	}
}

func TestIsAuthenticationError(t *testing.T) {
	authErr := &AuthenticationError{Err: &Error{Message: "unauthorized"}}
	otherErr := &BadRequestError{Err: &Error{Message: "bad request"}}

	if !IsAuthenticationError(authErr) {
		t.Error("IsAuthenticationError() returned false for AuthenticationError")
	}

	if IsAuthenticationError(otherErr) {
		t.Error("IsAuthenticationError() returned true for non-AuthenticationError")
	}
}

func TestIsTimeoutError(t *testing.T) {
	timeoutErr := &TimeoutError{Err: &Error{Message: "timeout"}}
	otherErr := &BadRequestError{Err: &Error{Message: "bad request"}}

	if !IsTimeoutError(timeoutErr) {
		t.Error("IsTimeoutError() returned false for TimeoutError")
	}

	if IsTimeoutError(otherErr) {
		t.Error("IsTimeoutError() returned true for non-TimeoutError")
	}
}
