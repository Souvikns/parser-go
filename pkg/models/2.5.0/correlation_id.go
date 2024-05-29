
package models

// CorrelationId represents a CorrelationId model.
type CorrelationId struct {
  Description string `json:"description,omitempty"`
  Location string `json:"location,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}