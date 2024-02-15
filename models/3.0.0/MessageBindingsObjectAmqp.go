
package models

// MessageBindingsObjectAmqp represents a MessageBindingsObjectAmqp model.
type MessageBindingsObjectAmqp struct {
  BindingVersion *MessageBindingsObjectAmqpBindingVersion
  ContentEncoding string
  MessageType string
  AdditionalProperties map[string]interface{}
}