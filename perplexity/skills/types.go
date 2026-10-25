// Code generated from Perplexity Python SDK 0.46.1 models. DO NOT EDIT.
package skills

type ManagedSkill struct {
	CreatedAt      *string `json:"created_at,omitempty"`
	Description    *string `json:"description,omitempty"`
	LatestRevision *string `json:"latest_revision,omitempty"`
	Name           *string `json:"name,omitempty"`
	Revision       *string `json:"revision,omitempty"`
	SkillID        string  `json:"skill_id"`
	UpdatedAt      *string `json:"updated_at,omitempty"`
}
type SkillDownload struct {
	DownloadURL string  `json:"download_url"`
	Name        *string `json:"name,omitempty"`
	Revision    *string `json:"revision,omitempty"`
}
type SkillList struct {
	NextPageToken *string        `json:"next_page_token,omitempty"`
	Skills        []SkillSummary `json:"skills"`
}
type SkillRevision struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Name      *string `json:"name,omitempty"`
	Revision  string  `json:"revision"`
}
type SkillRevisionList struct {
	NextPageToken *string         `json:"next_page_token,omitempty"`
	Revisions     []SkillRevision `json:"revisions"`
}
type SkillSummary struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Name      *string `json:"name,omitempty"`
	Revision  *string `json:"revision,omitempty"`
	SkillID   string  `json:"skill_id"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}
