
package models

// MessageBindingsObjectAnypointmq represents a MessageBindingsObjectAnypointmq model.
type MessageBindingsObjectAnypointmq struct {
  BindingVersion *MessageBindingsObjectAnypointmqBindingVersion `json:"bindingVersion,omitempty"`
  Headers *BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusMessageHeaders `json:"headers,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}