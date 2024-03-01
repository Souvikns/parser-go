
package models

// ChannelBindingsObjectWs represents a ChannelBindingsObjectWs model.
type ChannelBindingsObjectWs struct {
  BindingVersion *ChannelBindingsObjectWsBindingVersion `json:"bindingVersion"`
  Method *BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod `json:"method"`
  Query interface{} `json:"query"`
  Headers interface{} `json:"headers"`
  AdditionalProperties map[string]interface{} `json:"-"`
}