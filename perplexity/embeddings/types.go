package embeddings

import "encoding/json"

type Model string

type EncodingFormat string

type Currency string

const (
	ModelEmbedV106B Model = "pplx-embed-v1-0.6b"
	ModelEmbedV14B  Model = "pplx-embed-v1-4b"

	EncodingFormatBase64Int8   EncodingFormat = "base64_int8"
	EncodingFormatBase64Binary EncodingFormat = "base64_binary"

	CurrencyUSD Currency = "USD"
)

type Input struct {
	Text  *string
	Texts []string
}

func (i Input) MarshalJSON() ([]byte, error) {
	if i.Text != nil {
		return json.Marshal(*i.Text)
	}
	if i.Texts == nil {
		return []byte("null"), nil
	}
	return json.Marshal(i.Texts)
}

func (i *Input) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || string(data) == "null" {
		*i = Input{}
		return nil
	}
	var text string
	if err := json.Unmarshal(data, &text); err == nil {
		i.Text = &text
		i.Texts = nil
		return nil
	}
	var texts []string
	if err := json.Unmarshal(data, &texts); err != nil {
		return err
	}
	i.Text = nil
	i.Texts = texts
	return nil
}

type CreateParams struct {
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
	Embedding *string `json:"embedding,omitempty"`
	Index     *int    `json:"index,omitempty"`
	Object    *string `json:"object,omitempty"`
}

type CreateResponse struct {
	Data   []EmbeddingObject `json:"data,omitempty"`
	Model  *string           `json:"model,omitempty"`
	Object *string           `json:"object,omitempty"`
	Usage  *Usage            `json:"usage,omitempty"`
}
