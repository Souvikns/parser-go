
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string
  Name string
  Schema *Schema
  AdditionalProperties map[string]interface{}
}