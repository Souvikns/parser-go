
package models

// ChannelBindingsObjectWs represents a ChannelBindingsObjectWs model.
type ChannelBindingsObjectWs struct {
  BindingVersion *ChannelBindingsObjectWsBindingVersion `json:"bindingVersion,omitempty"`
  Method *BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod `json:"method,omitempty"`
  Query *BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelQuery `json:"query,omitempty"`
  Headers *BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelHeaders `json:"headers,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}