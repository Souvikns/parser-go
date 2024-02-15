
package models

// StreamObject represents a StreamObject model.
type StreamObject struct {
  Framing interface{}
  Read []Message
  Write []Message
  AdditionalProperties map[string]interface{}
}