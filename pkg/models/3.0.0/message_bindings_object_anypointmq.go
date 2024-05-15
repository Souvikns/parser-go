
package models

// MessageBindingsObjectAnypointmq represents a MessageBindingsObjectAnypointmq model.
type MessageBindingsObjectAnypointmq struct {
  BindingVersion *MessageBindingsObjectAnypointmqBindingVersion `json:"bindingVersion"`
  Headers *BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusMessageHeaders `json:"headers"`
  AdditionalProperties map[string]interface{} `json:"-"`
}