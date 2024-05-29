
package models

// ServerVariable represents a ServerVariable model.
type ServerVariable struct {
  Enum []string `json:"enum,omitempty"`
  ReservedDefault string `json:"default,omitempty"`
  Description string `json:"description,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}