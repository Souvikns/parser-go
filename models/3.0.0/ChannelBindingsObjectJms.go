
package models

// ChannelBindingsObjectJms represents a ChannelBindingsObjectJms model.
type ChannelBindingsObjectJms struct {
  BindingVersion *ChannelBindingsObjectJmsBindingVersion `json:"bindingVersion"`
  Destination string `json:"destination"`
  DestinationType *BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType `json:"destinationType"`
  AdditionalProperties map[string]interface{} `json:"-"`
}