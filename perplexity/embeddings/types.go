package embeddings

import (
	"encoding/json"
	"fmt"
)

type Model string

type EncodingFormat string

type Currency string

const (
	ModelEmbedV2Context06B Model = "pplx-embed-v2-context-0.6b"
	ModelEmbedV2Context9B  Model = "pplx-embed-v2-context-9b"
	ModelEmbedV106B        Model = "pplx-embed-v1-0.6b"
	ModelEmbedV14B         Model = "pplx-embed-v1-4b"

	EncodingFormatBase64Int8   EncodingFormat = "base64_int8"
	EncodingFormatBase64Binary EncodingFormat = "base64_binary"

	CurrencyUSD Currency = "USD"
)

type ImageInput struct {
	Image string `json:"image"`
	Type  string `json:"type"`
}
type Input struct {
	Documents [][]string
	Images    []ImageInput
	Text      *string
	Texts     []string
}

func (i Input) MarshalJSON() ([]byte, error) {
	if i.Text != nil {
		return json.Marshal(*i.Text)
	}
	if i.Texts != nil {
		return json.Marshal(i.Texts)
	}
	if i.Documents != nil {
		return json.Marshal(i.Documents)
	}
	if i.Images != nil {
		return json.Marshal(i.Images)
	}
	return []byte("null"), nil
}

func (i *Input) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*i = Input{}
		return nil
	}
	var text string
	if err := json.Unmarshal(data, &text); err == nil {
		*i = Input{Text: &text}
		return nil
	}
	var texts []string
	if err := json.Unmarshal(data, &texts); err == nil {
		*i = Input{Texts: texts}
		return nil
	}
	var documents [][]string
	if err := json.Unmarshal(data, &documents); err == nil {
		*i = Input{Documents: documents}
		return nil
	}
	var images []ImageInput
	if err := json.Unmarshal(data, &images); err != nil {
		return fmt.Errorf("invalid embedding input: %w", err)
	}
	*i = Input{Images: images}
	return nil
}

type CreateParams struct {
	InputType      *string         `json:"input_type,omitempty"`
	NumChunks      *int            `json:"num_chunks,omitempty"`
	Input          Input           `json:"input"`
	Model          Model           `json:"model"`
	Dimensions     *int            `json:"dimensions,omitempty"`
	EncodingFormat *EncodingFormat `json:"encoding_format,omitempty"`
}

type Cost struct {
	Currency  *Currency `json:"currency,omitempty"`
	InputCost *float64  `json:"input_cost,omitempty"`
	TotalCost *float64  `json:"total_cost,omitempty"`
}

type Usage struct {
	Cost         *Cost `json:"cost,omitempty"`
	PromptTokens *int  `json:"prompt_tokens,omitempty"`
	TotalTokens  *int  `json:"total_tokens,omitempty"`
}

type EmbeddingObject struct {
	Data      []EmbeddingObject `json:"data,omitempty"`
	Embedding *string           `json:"embedding,omitempty"`
	Index     *int              `json:"index,omitempty"`
	Object    *string           `json:"object,omitempty"`
}

type CreateResponse struct {
	Data   []EmbeddingObject `json:"data,omitempty"`
	Model  *string           `json:"model,omitempty"`
	Object *string           `json:"object,omitempty"`
	Usage  *Usage            `json:"usage,omitempty"`
}
