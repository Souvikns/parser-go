
package models

// Operation represents a Operation model.
type Operation struct {
  Action *OperationAction `json:"action,omitempty"`
  Channel *Reference `json:"channel,omitempty"`
  Messages []Reference `json:"messages,omitempty"`
  Reply *OperationReply `json:"reply,omitempty"`
  Traits []OperationTraitsItem `json:"traits,omitempty"`
  Title string `json:"title,omitempty"`
  Summary string `json:"summary,omitempty"`
  Description string `json:"description,omitempty"`
  Security []SecurityRequirementsItem `json:"security,omitempty"`
  Tags []OperationTagsItem `json:"tags,omitempty"`
  ExternalDocs *OperationExternalDocs `json:"externalDocs,omitempty"`
  Bindings *OperationBindings `json:"bindings,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}