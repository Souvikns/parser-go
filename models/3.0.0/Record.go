
package models

// Record represents a Record model.
type Record struct {
  ReservedType string `json:"type"`
  Name string `json:"name"`
  Namespace string `json:"namespace"`
  Doc string `json:"doc"`
  Aliases []string `json:"aliases"`
  Fields []Field `json:"fields"`
  AdditionalProperties map[string]interface{} `json:"-"`
}