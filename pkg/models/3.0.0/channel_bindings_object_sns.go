
package models

// ChannelBindingsObjectSns represents a ChannelBindingsObjectSns model.
type ChannelBindingsObjectSns struct {
  BindingVersion *ChannelBindingsObjectSnsBindingVersion `json:"bindingVersion"`
  Name string `json:"name"`
  Ordering *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrdering `json:"ordering"`
  Policy *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelPolicy `json:"policy"`
  Tags map[string]interface{} `json:"tags"`
  AdditionalProperties map[string]interface{} `json:"-"`
}