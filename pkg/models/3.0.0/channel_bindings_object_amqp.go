
package models

// ChannelBindingsObjectAmqp represents a ChannelBindingsObjectAmqp model.
type ChannelBindingsObjectAmqp struct {
  BindingVersion *ChannelBindingsObjectAmqpBindingVersion `json:"bindingVersion,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}