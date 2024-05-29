
package models

// OperationReply represents a OperationReply model.
type OperationReply struct {
  Address *OperationReplyAddress `json:"address,omitempty"`
  Channel *Reference `json:"channel,omitempty"`
  Messages []Reference `json:"messages,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}