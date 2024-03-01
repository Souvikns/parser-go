
package models

// ServerVariable represents a ServerVariable model.
type ServerVariable struct {
  Enum []string `json:"enum"`
  ReservedDefault string `json:"default"`
  Description string `json:"description"`
  Examples []string `json:"examples"`
  AdditionalProperties map[string]interface{} `json:"-"`
}