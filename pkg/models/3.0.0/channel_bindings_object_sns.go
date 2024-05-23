
package models

// ChannelBindingsObjectSns represents a ChannelBindingsObjectSns model.
type ChannelBindingsObjectSns struct {
  BindingVersion *ChannelBindingsObjectSnsBindingVersion `json:"bindingVersion,omitempty"`
  Name string `json:"name,omitempty"`
  Ordering *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrdering `json:"ordering,omitempty"`
  Policy *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelPolicy `json:"policy,omitempty"`
  Tags map[string]interface{} `json:"tags,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}