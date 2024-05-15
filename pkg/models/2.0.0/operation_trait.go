
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Summary string `json:"summary"`
  Description string `json:"description"`
  Tags []Tag `json:"tags"`
  ExternalDocs *ExternalDocs `json:"externalDocs"`
  OperationId string `json:"operationId"`
  Bindings *BindingsObject `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}