
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Summary string `json:"summary,omitempty"`
  Description string `json:"description,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  OperationId string `json:"operationId,omitempty"`
  Security []map[string][]string `json:"security,omitempty"`
  Bindings *BindingsObject `json:"bindings,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}