
package models

// ReservedMap represents a ReservedMap model.
type ReservedMap struct {
  ReservedType string `json:"type"`
  Name string `json:"name"`
  Namespace string `json:"namespace"`
  Doc string `json:"doc"`
  Aliases []string `json:"aliases"`
  Values interface{} `json:"values"`
  AdditionalProperties map[string]interface{} `json:"-"`
}