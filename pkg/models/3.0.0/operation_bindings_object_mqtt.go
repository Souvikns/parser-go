
package models

// OperationBindingsObjectMqtt represents a OperationBindingsObjectMqtt model.
type OperationBindingsObjectMqtt struct {
  BindingVersion *OperationBindingsObjectMqttBindingVersion `json:"bindingVersion,omitempty"`
  Qos *BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos `json:"qos,omitempty"`
  Retain bool `json:"retain,omitempty"`
  MessageExpiryInterval *BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationMessageExpiryInterval `json:"messageExpiryInterval,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}