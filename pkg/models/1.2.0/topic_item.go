
package models

// TopicItem represents a TopicItem model.
type TopicItem struct {
  Ref string `json:"$ref,omitempty"`
  Parameters []Parameter `json:"parameters,omitempty"`
  Publish *Operation `json:"publish,omitempty"`
  Subscribe *Operation `json:"subscribe,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}