
package models

// ServerBindingsObjectJms represents a ServerBindingsObjectJms model.
type ServerBindingsObjectJms struct {
  BindingVersion *BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerBindingVersion `json:"bindingVersion"`
  JmsConnectionFactory string `json:"jmsConnectionFactory"`
  Properties []ServerSchema `json:"properties"`
  ClientId string `json:"clientID"`
  AdditionalProperties map[string]interface{} `json:"-"`
}