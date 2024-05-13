
package models

// ChannelBindingsObjectAnypointmq represents a ChannelBindingsObjectAnypointmq model.
type ChannelBindingsObjectAnypointmq struct {
  BindingVersion *ChannelBindingsObjectAnypointmqBindingVersion `json:"bindingVersion"`
  Destination string `json:"destination"`
  DestinationType *BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType `json:"destinationType"`
  AdditionalProperties map[string]interface{} `json:"-"`
}