// Code generated from Perplexity Python SDK 0.46.1 models. DO NOT EDIT.
package files

type FilePurpose = string
type FileDeleted struct {
	Deleted bool   `json:"deleted"`
	ID      string `json:"id"`
	Object  string `json:"object"`
}
type FileObject struct {
	Bytes     int         `json:"bytes"`
	CreatedAt int         `json:"created_at"`
	ExpiresAt *int        `json:"expires_at,omitempty"`
	Filename  string      `json:"filename"`
	ID        string      `json:"id"`
	Object    string      `json:"object"`
	Purpose   FilePurpose `json:"purpose"`
}
type FileObjectList struct {
	Data    []FileObject `json:"data"`
	FirstID *string      `json:"first_id,omitempty"`
	HasMore bool         `json:"has_more"`
	LastID  *string      `json:"last_id,omitempty"`
	Object  string       `json:"object"`
}
