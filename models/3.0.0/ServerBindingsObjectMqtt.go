
package models

// ServerBindingsObjectMqtt represents a ServerBindingsObjectMqtt model.
type ServerBindingsObjectMqtt struct {
  BindingVersion *ServerBindingsObjectMqttBindingVersion `json:"bindingVersion"`
  ClientId string `json:"clientId"`
  CleanSession bool `json:"cleanSession"`
  LastWill *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill `json:"lastWill"`
  KeepAlive int `json:"keepAlive"`
  SessionExpiryInterval interface{} `json:"sessionExpiryInterval"`
  MaximumPacketSize interface{} `json:"maximumPacketSize"`
  AdditionalProperties map[string]interface{} `json:"-"`
}