
package models

// ChannelBindingsObjectWs represents a ChannelBindingsObjectWs model.
type ChannelBindingsObjectWs struct {
  BindingVersion *ChannelBindingsObjectWsBindingVersion `json:"bindingVersion"`
  Method *BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod `json:"method"`
  Query *BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelQuery `json:"query"`
  Headers *BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelHeaders `json:"headers"`
  AdditionalProperties map[string]interface{} `json:"-"`
}