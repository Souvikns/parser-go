
package models

// ChannelBindingsObjectSns represents a ChannelBindingsObjectSns model.
type ChannelBindingsObjectSns struct {
  BindingVersion string `json:"bindingVersion"`
  Name string `json:"name"`
  Ordering *ChannelSchema `json:"ordering"`
  Policy *ChannelSchema `json:"policy"`
  Tags map[string]interface{} `json:"tags"`
  AdditionalProperties map[string]interface{} `json:"-"`
}