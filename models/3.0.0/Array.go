
package models

// Array represents a Array model.
type Array struct {
  ReservedType string
  Name string
  Namespace string
  Doc string
  Aliases []string
  Items interface{}
  AdditionalProperties map[string]interface{}
}