
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string `json:"description,omitempty"`
  Schema *Schema `json:"schema,omitempty"`
  Location string `json:"location,omitempty"`
  Ref string `json:"$ref,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}