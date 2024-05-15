
package models

// OperationReply represents a OperationReply model.
type OperationReply struct {
  Address *OperationReplyAddress `json:"address"`
  Channel *Reference `json:"channel"`
  Messages []Reference `json:"messages"`
  AdditionalProperties map[string]interface{} `json:"-"`
}