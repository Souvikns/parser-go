
package models

// ChannelBindingsObjectSns represents a ChannelBindingsObjectSns model.
type ChannelBindingsObjectSns struct {
  BindingVersion string
  Name string
  Ordering *ChannelSchema
  Policy *ChannelSchema
  Tags map[string]interface{}
  AdditionalProperties map[string]interface{}
}