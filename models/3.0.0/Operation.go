
package models

// Operation represents a Operation model.
type Operation struct {
  Action *OperationAction `json:"action"`
  Channel *Reference `json:"channel"`
  Messages []Reference `json:"messages"`
  Reply *OperationReply `json:"reply"`
  Traits []OperationTraitsItem `json:"traits"`
  Title string `json:"title"`
  Summary string `json:"summary"`
  Description string `json:"description"`
  Security []SecurityRequirementsItem `json:"security"`
  Tags []OperationTagsItem `json:"tags"`
  ExternalDocs *OperationExternalDocs `json:"externalDocs"`
  Bindings *OperationBindings `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}