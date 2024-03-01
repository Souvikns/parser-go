
package models

// Enum represents a Enum model.
type Enum struct {
  ReservedType string `json:"type"`
  Name string `json:"name"`
  Namespace string `json:"namespace"`
  Doc string `json:"doc"`
  Aliases []string `json:"aliases"`
  Symbols []string `json:"symbols"`
  AdditionalProperties map[string]interface{} `json:"-"`
}