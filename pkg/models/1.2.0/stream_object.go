
package models

// StreamObject represents a StreamObject model.
type StreamObject struct {
  Framing *StreamFramingObject `json:"framing,omitempty"`
  Read []Message `json:"read,omitempty"`
  Write []Message `json:"write,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}