
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Title string `json:"title"`
  Summary string `json:"summary"`
  Description string `json:"description"`
  Security []SecurityRequirementsItem `json:"security"`
  Tags []OperationTagsItem `json:"tags"`
  ExternalDocs *OperationExternalDocs `json:"externalDocs"`
  Bindings *OperationTraitBindings `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}