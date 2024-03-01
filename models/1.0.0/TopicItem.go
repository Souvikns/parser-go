
package models

// TopicItem represents a TopicItem model.
type TopicItem struct {
  Ref string `json:$ref`
  Publish *Message `json:publish`
  Subscribe *Message `json:subscribe`
  Deprecated bool `json:deprecated`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}