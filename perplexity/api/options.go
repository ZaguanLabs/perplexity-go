package api

import (
	"net/http"
	"time"
)

type RequestOptions struct {
	Headers   map[string]string
	Query     map[string]any
	ExtraBody map[string]any
	Timeout   time.Duration
}

type RequestOption func(*RequestOptions)

type RawResponse[T any] struct {
	Data       *T
	StatusCode int
	Headers    http.Header
	Body       []byte
	RequestID  string
}

func WithHeader(key, value string) RequestOption {
	return func(o *RequestOptions) {
		if o.Headers == nil {
			o.Headers = make(map[string]string)
		}
		o.Headers[key] = value
	}
}

func WithHeaders(headers map[string]string) RequestOption {
	return func(o *RequestOptions) {
		if o.Headers == nil {
			o.Headers = make(map[string]string, len(headers))
		}
		for key, value := range headers {
			o.Headers[key] = value
		}
	}
}

func WithQuery(key string, value any) RequestOption {
	return func(o *RequestOptions) {
		if o.Query == nil {
			o.Query = make(map[string]any)
		}
		o.Query[key] = value
	}
}

func WithQueryParams(query map[string]any) RequestOption {
	return func(o *RequestOptions) {
		if o.Query == nil {
			o.Query = make(map[string]any, len(query))
		}
		for key, value := range query {
			o.Query[key] = value
		}
	}
}

func WithExtraBody(key string, value any) RequestOption {
	return func(o *RequestOptions) {
		if o.ExtraBody == nil {
			o.ExtraBody = make(map[string]any)
		}
		o.ExtraBody[key] = value
	}
}

func WithExtraBodyFields(fields map[string]any) RequestOption {
	return func(o *RequestOptions) {
		if o.ExtraBody == nil {
			o.ExtraBody = make(map[string]any, len(fields))
		}
		for key, value := range fields {
			o.ExtraBody[key] = value
		}
	}
}

func WithTimeout(timeout time.Duration) RequestOption {
	return func(o *RequestOptions) {
		o.Timeout = timeout
	}
}

func ApplyRequestOptions(opts []RequestOption) RequestOptions {
	var options RequestOptions
	for _, opt := range opts {
		if opt != nil {
			opt(&options)
		}
	}
	return options
}
