
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string
  Enum []string
  ReservedDefault string
  Examples []string
  Location string
  AdditionalProperties map[string]interface{}
}