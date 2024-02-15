
package models

// Record represents a Record model.
type Record struct {
  ReservedType string
  Name string
  Namespace string
  Doc string
  Aliases []string
  Fields []Field
  AdditionalProperties map[string]interface{}
}