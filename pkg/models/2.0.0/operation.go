
package models

// Operation represents a Operation model.
type Operation struct {
  Traits []OperationTraitsItem `json:"traits,omitempty"`
  Summary string `json:"summary,omitempty"`
  Description string `json:"description,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  OperationId string `json:"operationId,omitempty"`
  Bindings *BindingsObject `json:"bindings,omitempty"`
  Message *Message `json:"message,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}