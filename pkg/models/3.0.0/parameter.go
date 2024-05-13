
package models

// Parameter represents a Parameter model.
type Parameter struct {
  Description string `json:"description"`
  Enum []string `json:"enum"`
  ReservedDefault string `json:"default"`
  Examples []string `json:"examples"`
  Location string `json:"location"`
  AdditionalProperties map[string]interface{} `json:"-"`
}