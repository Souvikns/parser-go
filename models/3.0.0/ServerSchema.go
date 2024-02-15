
package models

// ServerSchema represents a ServerSchema model.
type ServerSchema struct {
  JmsConnectionFactory string
  Properties []ServerSchema
  ClientId string
  BindingVersion *BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion
  AdditionalProperties map[string]interface{}
}