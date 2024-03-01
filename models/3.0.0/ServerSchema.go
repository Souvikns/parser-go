
package models

// ServerSchema represents a ServerSchema model.
type ServerSchema struct {
  JmsConnectionFactory string `json:"jmsConnectionFactory"`
  Properties []ServerSchema `json:"properties"`
  ClientId string `json:"clientID"`
  BindingVersion *BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion `json:"bindingVersion"`
  AdditionalProperties map[string]interface{} `json:"-"`
}