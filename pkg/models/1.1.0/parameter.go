
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string `json:"description"`
  Name string `json:"name"`
  Schema *Schema `json:"schema"`
  AdditionalProperties map[string]interface{} `json:"-"`
}