
package models

// MessageBindingsObjectAmqp represents a MessageBindingsObjectAmqp model.
type MessageBindingsObjectAmqp struct {
  BindingVersion *MessageBindingsObjectAmqpBindingVersion `json:"bindingVersion,omitempty"`
  ContentEncoding string `json:"contentEncoding,omitempty"`
  MessageType string `json:"messageType,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}