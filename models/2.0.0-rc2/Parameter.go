
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string `json:"description"`
  Schema *Schema `json:"schema"`
  Location string `json:"location"`
  Ref string `json:"$ref"`
  AdditionalProperties map[string]interface{} `json:"-"`
}