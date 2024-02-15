
package models

// ReservedMap represents a ReservedMap model.
type ReservedMap struct {
  ReservedType string
  Name string
  Namespace string
  Doc string
  Aliases []string
  Values interface{}
  AdditionalProperties map[string]interface{}
}