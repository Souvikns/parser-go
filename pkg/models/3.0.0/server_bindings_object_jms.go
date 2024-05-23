
package models

// ServerBindingsObjectJms represents a ServerBindingsObjectJms model.
type ServerBindingsObjectJms struct {
  BindingVersion *ServerBindingsObjectJmsBindingVersion `json:"bindingVersion,omitempty"`
  JmsConnectionFactory string `json:"jmsConnectionFactory,omitempty"`
  Properties []BindingsMinusJmsMinus_0Dot_0Dot_1MinusServerProperty `json:"properties,omitempty"`
  ClientId string `json:"clientID,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}