
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Title string `json:"title"`
  Summary string `json:"summary"`
  Description string `json:"description"`
  Security []interface{} `json:"security"`
  Tags []interface{} `json:"tags"`
  ExternalDocs interface{} `json:"externalDocs"`
  Bindings interface{} `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}