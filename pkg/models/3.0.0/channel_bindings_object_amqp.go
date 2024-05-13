
package models

// ChannelBindingsObjectAmqp represents a ChannelBindingsObjectAmqp model.
type ChannelBindingsObjectAmqp struct {
  BindingVersion *ChannelBindingsObjectAmqpBindingVersion `json:"bindingVersion"`
  AdditionalProperties map[string]interface{} `json:"-"`
}