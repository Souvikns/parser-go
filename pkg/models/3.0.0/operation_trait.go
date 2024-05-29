
package models

// OperationTrait represents a OperationTrait model.
type OperationTrait struct {
  Title string `json:"title,omitempty"`
  Summary string `json:"summary,omitempty"`
  Description string `json:"description,omitempty"`
  Security []SecurityRequirementsItem `json:"security,omitempty"`
  Tags []OperationTagsItem `json:"tags,omitempty"`
  ExternalDocs *OperationExternalDocs `json:"externalDocs,omitempty"`
  Bindings *OperationTraitBindings `json:"bindings,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}