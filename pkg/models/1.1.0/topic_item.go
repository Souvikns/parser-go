
package models

// TopicItem represents a TopicItem model.
type TopicItem struct {
  Ref string `json:"$ref"`
  Parameters []Parameter `json:"parameters"`
  Publish *Operation `json:"publish"`
  Subscribe *Operation `json:"subscribe"`
  Deprecated bool `json:"deprecated"`
  AdditionalProperties map[string]interface{} `json:"-"`
}