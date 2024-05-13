
package models

// ServerBindingsObjectMqtt represents a ServerBindingsObjectMqtt model.
type ServerBindingsObjectMqtt struct {
  BindingVersion *ServerBindingsObjectMqttBindingVersion `json:"bindingVersion"`
  ClientId string `json:"clientId"`
  CleanSession bool `json:"cleanSession"`
  LastWill *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill `json:"lastWill"`
  KeepAlive int `json:"keepAlive"`
  SessionExpiryInterval *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerSessionExpiryInterval `json:"sessionExpiryInterval"`
  MaximumPacketSize *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerMaximumPacketSize `json:"maximumPacketSize"`
  AdditionalProperties map[string]interface{} `json:"-"`
}