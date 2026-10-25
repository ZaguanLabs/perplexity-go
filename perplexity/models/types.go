// Code generated from Perplexity Python SDK 0.46.1 models. DO NOT EDIT.
package models

type Model struct {
	Created int    `json:"created"`
	ID      string `json:"id"`
	Object  string `json:"object"`
	OwnedBy string `json:"owned_by"`
}
type ModelList struct {
	Data   []Model `json:"data"`
	Object string  `json:"object"`
}
