package contextualizedembeddings

import "github.com/ZaguanLabs/perplexity-go/perplexity/embeddings"

type Model string

type EncodingFormat string

const (
	ModelEmbedContextV106B Model = "pplx-embed-context-v1-0.6b"
	ModelEmbedContextV14B  Model = "pplx-embed-context-v1-4b"

	EncodingFormatBase64Int8   EncodingFormat = "base64_int8"
	EncodingFormatBase64Binary EncodingFormat = "base64_binary"
)

type CreateParams struct {
	Input          [][]string      `json:"input"`
	Model          Model           `json:"model"`
	Dimensions     *int            `json:"dimensions,omitempty"`
	EncodingFormat *EncodingFormat `json:"encoding_format,omitempty"`
}

type EmbeddingObject = embeddings.EmbeddingObject

type Usage = embeddings.Usage

type ContextualizedEmbeddingObject struct {
	Data   []EmbeddingObject `json:"data,omitempty"`
	Index  *int              `json:"index,omitempty"`
	Object *string           `json:"object,omitempty"`
}

type CreateResponse struct {
	Data   []ContextualizedEmbeddingObject `json:"data,omitempty"`
	Model  *string                         `json:"model,omitempty"`
	Object *string                         `json:"object,omitempty"`
	Usage  *Usage                          `json:"usage,omitempty"`
}
