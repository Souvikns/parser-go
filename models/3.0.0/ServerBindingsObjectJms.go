
package models

// ServerBindingsObjectJms represents a ServerBindingsObjectJms model.
type ServerBindingsObjectJms struct {
  BindingVersion *BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion
  JmsConnectionFactory string
  Properties []ServerSchema
  ClientId string
  AdditionalProperties map[string]interface{}
}