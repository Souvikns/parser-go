
package models

// StreamObject represents a StreamObject model.
type StreamObject struct {
  Framing interface{} `json:framing`
  Read []Message `json:read`
  Write []Message `json:write`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}