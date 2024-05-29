
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string `json:"description,omitempty"`
  Enum []string `json:"enum,omitempty"`
  ReservedDefault string `json:"default,omitempty"`
  Examples []string `json:"examples,omitempty"`
  Location string `json:"location,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}