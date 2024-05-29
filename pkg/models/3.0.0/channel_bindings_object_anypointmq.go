
package models

// ChannelBindingsObjectAnypointmq represents a ChannelBindingsObjectAnypointmq model.
type ChannelBindingsObjectAnypointmq struct {
  BindingVersion *ChannelBindingsObjectAnypointmqBindingVersion `json:"bindingVersion,omitempty"`
  Destination string `json:"destination,omitempty"`
  DestinationType *BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType `json:"destinationType,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}