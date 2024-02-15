
package models

// ChannelBindingsObjectJms represents a ChannelBindingsObjectJms model.
type ChannelBindingsObjectJms struct {
  BindingVersion *ChannelBindingsObjectJmsBindingVersion
  Destination string
  DestinationType *BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType
  AdditionalProperties map[string]interface{}
}