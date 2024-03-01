
package models

// TopicItem represents a TopicItem model.
type TopicItem struct {
  Ref string `json:$ref`
  Parameters []Parameter `json:parameters`
  Publish interface{} `json:publish`
  Subscribe interface{} `json:subscribe`
  Deprecated bool `json:deprecated`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}