
package models

// Fixed represents a Fixed model.
type Fixed struct {
  ReservedType string
  Name string
  Namespace string
  Doc string
  Aliases []string
  Size float64
  AdditionalProperties map[string]interface{}
}