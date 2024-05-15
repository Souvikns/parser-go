
package models

// ServerBindingsObjectJms represents a ServerBindingsObjectJms model.
type ServerBindingsObjectJms struct {
  BindingVersion *ServerBindingsObjectJmsBindingVersion `json:"bindingVersion"`
  JmsConnectionFactory string `json:"jmsConnectionFactory"`
  Properties []BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerProperty `json:"properties"`
  ClientId string `json:"clientID"`
  AdditionalProperties map[string]interface{} `json:"-"`
}