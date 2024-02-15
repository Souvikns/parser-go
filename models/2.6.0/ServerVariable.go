
package models

// ServerVariable represents a ServerVariable model.
type ServerVariable struct {
  Enum []string
  ReservedDefault string
  Description string
  Examples []string
  AdditionalProperties map[string]interface{}
}