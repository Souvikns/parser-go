
package models

// TopicItem represents a TopicItem model.
type TopicItem struct {
  Ref string
  Parameters []Parameter
  Publish interface{}
  Subscribe interface{}
  Deprecated bool
  AdditionalProperties map[string]interface{}
}