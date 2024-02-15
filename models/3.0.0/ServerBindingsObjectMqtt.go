
package models

// ServerBindingsObjectMqtt represents a ServerBindingsObjectMqtt model.
type ServerBindingsObjectMqtt struct {
  BindingVersion *ServerBindingsObjectMqttBindingVersion
  ClientId string
  CleanSession bool
  LastWill *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWill
  KeepAlive int
  SessionExpiryInterval interface{}
  MaximumPacketSize interface{}
  AdditionalProperties map[string]interface{}
}