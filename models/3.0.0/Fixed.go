
package models

// Fixed represents a Fixed model.
type Fixed struct {
  ReservedType string `json:"type"`
  Name string `json:"name"`
  Namespace string `json:"namespace"`
  Doc string `json:"doc"`
  Aliases []string `json:"aliases"`
  Size float64 `json:"size"`
  AdditionalProperties map[string]interface{} `json:"-"`
}