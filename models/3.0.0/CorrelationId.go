
package models

// CorrelationId represents a CorrelationId model.
type CorrelationId struct {
  Description string `json:"description"`
  Location string `json:"location"`
  AdditionalProperties map[string]interface{} `json:"-"`
}