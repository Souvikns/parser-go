
package models

// TopicItem represents a TopicItem model.
type TopicItem struct {
  Ref string `json:"$ref,omitempty"`
  Publish *Message `json:"publish,omitempty"`
  Subscribe *Message `json:"subscribe,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}