
package models

// ChannelBindingsObjectJms represents a ChannelBindingsObjectJms model.
type ChannelBindingsObjectJms struct {
  BindingVersion *ChannelBindingsObjectJmsBindingVersion `json:"bindingVersion,omitempty"`
  Destination string `json:"destination,omitempty"`
  DestinationType *BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType `json:"destinationType,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}