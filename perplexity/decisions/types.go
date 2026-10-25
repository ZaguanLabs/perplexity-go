// Code generated from Perplexity Python SDK 0.46.1 models. DO NOT EDIT.
package decisions

type DecisionQuestion = any
type DecisionAnswer = any
type DecisionChoiceAnswer struct {
	Choice        string             `json:"choice"`
	Confidence    float64            `json:"confidence"`
	Probabilities map[string]float64 `json:"probabilities"`
	Type          string             `json:"type"`
}
type DecisionChoiceQuestion struct {
	Criteria     map[string]any `json:"criteria"`
	Instructions any            `json:"instructions,omitempty"`
	Type         string         `json:"type"`
}
type DecisionNoulAnswer struct {
	Noul float64 `json:"noul"`
	Type string  `json:"type"`
}
type DecisionNoulCriteria struct {
	False any `json:"false,omitempty"`
	True  any `json:"true,omitempty"`
}
type DecisionNoulQuestion struct {
	Criteria     *DecisionNoulCriteria `json:"criteria,omitempty"`
	Instructions any                   `json:"instructions,omitempty"`
	Type         string                `json:"type"`
}
type DecisionScoreAnswer struct {
	Confidence    float64            `json:"confidence"`
	Legend        map[string]any     `json:"legend"`
	Probabilities map[string]float64 `json:"probabilities"`
	Score         float64            `json:"score"`
	Type          string             `json:"type"`
}
type DecisionScoreQuestion struct {
	Criteria     []any  `json:"criteria"`
	Instructions any    `json:"instructions,omitempty"`
	Type         string `json:"type"`
}
type DecisionUsage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}
type DecisionsRequest struct {
	Model     string                      `json:"model"`
	Questions map[string]DecisionQuestion `json:"questions"`
	State     any                         `json:"state"`
}
type DecisionsResponse struct {
	Answers map[string]DecisionAnswer `json:"answers"`
	Model   string                    `json:"model"`
	Usage   DecisionUsage             `json:"usage"`
}
