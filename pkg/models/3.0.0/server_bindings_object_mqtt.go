
package models

// ServerBindingsObjectMqtt represents a ServerBindingsObjectMqtt model.
type ServerBindingsObjectMqtt struct {
  BindingVersion *ServerBindingsObjectMqttBindingVersion `json:"bindingVersion,omitempty"`
  ClientId string `json:"clientId,omitempty"`
  CleanSession bool `json:"cleanSession,omitempty"`
  LastWill *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill `json:"lastWill,omitempty"`
  KeepAlive int `json:"keepAlive,omitempty"`
  SessionExpiryInterval *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerSessionExpiryInterval `json:"sessionExpiryInterval,omitempty"`
  MaximumPacketSize *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerMaximumPacketSize `json:"maximumPacketSize,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}