
package models

// Operation represents a Operation model.
type Operation struct {
  Traits []OperationTraitsItem `json:"traits"`
  Summary string `json:"summary"`
  Description string `json:"description"`
  Tags []Tag `json:"tags"`
  ExternalDocs *ExternalDocs `json:"externalDocs"`
  OperationId string `json:"operationId"`
  Bindings *BindingsObject `json:"bindings"`
  Message *Message `json:"message"`
  AdditionalProperties map[string]interface{} `json:"-"`
}