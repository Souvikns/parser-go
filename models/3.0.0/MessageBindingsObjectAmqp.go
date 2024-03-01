
package models

// MessageBindingsObjectAmqp represents a MessageBindingsObjectAmqp model.
type MessageBindingsObjectAmqp struct {
  BindingVersion *MessageBindingsObjectAmqpBindingVersion `json:"bindingVersion"`
  ContentEncoding string `json:"contentEncoding"`
  MessageType string `json:"messageType"`
  AdditionalProperties map[string]interface{} `json:"-"`
}