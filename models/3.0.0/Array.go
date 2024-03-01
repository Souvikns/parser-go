
package models

// Array represents a Array model.
type Array struct {
  ReservedType string `json:"type"`
  Name string `json:"name"`
  Namespace string `json:"namespace"`
  Doc string `json:"doc"`
  Aliases []string `json:"aliases"`
  Items interface{} `json:"items"`
  AdditionalProperties map[string]interface{} `json:"-"`
}