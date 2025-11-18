package perplexity

import (
	"os"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name    string
		apiKey  string
		envKey  string
		wantErr bool
	}{
		{
			name:    "with API key",
			apiKey:  "test-key",
			wantErr: false,
		},
		{
			name:    "with env var",
			apiKey:  "",
			envKey:  "env-test-key",
			wantErr: false,
		},
		{
			name:    "no API key",
			apiKey:  "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up environment
			if tt.envKey != "" {
				os.Setenv("PERPLEXITY_API_KEY", tt.envKey)
				defer os.Unsetenv("PERPLEXITY_API_KEY")
			} else {
				os.Unsetenv("PERPLEXITY_API_KEY")
			}

			client, err := NewClient(tt.apiKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if client == nil {
					t.Error("NewClient() returned nil client")
					return
				}

				expectedKey := tt.apiKey
				if expectedKey == "" {
					expectedKey = tt.envKey
				}

				if client.APIKey() != expectedKey {
					t.Errorf("APIKey() = %v, want %v", client.APIKey(), expectedKey)
				}

				if client.BaseURL() != DefaultBaseURL {
					t.Errorf("BaseURL() = %v, want %v", client.BaseURL(), DefaultBaseURL)
				}
			}
		})
	}
}

func TestClientOptions(t *testing.T) {
	t.Run("WithBaseURL", func(t *testing.T) {
		customURL := "https://custom.api.com"
		client, err := NewClient("test-key", WithBaseURL(customURL))
		if err != nil {
			t.Fatalf("NewClient() error = %v", err)
		}

		if client.BaseURL() != customURL {
			t.Errorf("BaseURL() = %v, want %v", client.BaseURL(), customURL)
		}
	})

	t.Run("WithTimeout", func(t *testing.T) {
		timeout := 30 * time.Second
		client, err := NewClient("test-key", WithTimeout(timeout))
		if err != nil {
			t.Fatalf("NewClient() error = %v", err)
		}

		if client.httpClient.Timeout != timeout {
			t.Errorf("httpClient.Timeout = %v, want %v", client.httpClient.Timeout, timeout)
		}
	})

	t.Run("WithMaxRetries", func(t *testing.T) {
		maxRetries := 5
		client, err := NewClient("test-key", WithMaxRetries(maxRetries))
		if err != nil {
			t.Fatalf("NewClient() error = %v", err)
		}

		if client.maxRetries != maxRetries {
			t.Errorf("maxRetries = %v, want %v", client.maxRetries, maxRetries)
		}
	})

	t.Run("WithDefaultHeader", func(t *testing.T) {
		client, err := NewClient("test-key", WithDefaultHeader("X-Custom", "value"))
		if err != nil {
			t.Fatalf("NewClient() error = %v", err)
		}

		if client.defaultHeaders["X-Custom"] != "value" {
			t.Errorf("defaultHeaders[X-Custom] = %v, want value", client.defaultHeaders["X-Custom"])
		}
	})

	t.Run("multiple options", func(t *testing.T) {
		customURL := "https://custom.api.com"
		timeout := 30 * time.Second
		maxRetries := 5

		client, err := NewClient("test-key",
			WithBaseURL(customURL),
			WithTimeout(timeout),
			WithMaxRetries(maxRetries),
		)
		if err != nil {
			t.Fatalf("NewClient() error = %v", err)
		}

		if client.BaseURL() != customURL {
			t.Errorf("BaseURL() = %v, want %v", client.BaseURL(), customURL)
		}
		if client.httpClient.Timeout != timeout {
			t.Errorf("httpClient.Timeout = %v, want %v", client.httpClient.Timeout, timeout)
		}
		if client.maxRetries != maxRetries {
			t.Errorf("maxRetries = %v, want %v", client.maxRetries, maxRetries)
		}
	})
}
