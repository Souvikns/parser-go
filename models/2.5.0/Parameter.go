
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string
  Schema interface{}
  Location string
  AdditionalProperties map[string]interface{}
}