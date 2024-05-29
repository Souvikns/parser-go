
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string `json:"description,omitempty"`
  Name string `json:"name,omitempty"`
  Schema *Schema `json:"schema,omitempty"`
  Ref string `json:"$ref,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}