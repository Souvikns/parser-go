
package models

// TopicItem represents a TopicItem model.
type TopicItem struct {
  Ref string
  Publish *Message
  Subscribe *Message
  Deprecated bool
  AdditionalProperties map[string]interface{}
}