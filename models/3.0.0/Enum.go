
package models

// Enum represents a Enum model.
type Enum struct {
  ReservedType string
  Name string
  Namespace string
  Doc string
  Aliases []string
  Symbols []string
  AdditionalProperties map[string]interface{}
}