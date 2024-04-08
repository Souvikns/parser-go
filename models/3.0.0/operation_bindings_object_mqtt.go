
package models

// OperationBindingsObjectMqtt represents a OperationBindingsObjectMqtt model.
type OperationBindingsObjectMqtt struct {
  BindingVersion *OperationBindingsObjectMqttBindingVersion `json:"bindingVersion"`
  Qos *BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos `json:"qos"`
  Retain bool `json:"retain"`
  MessageExpiryInterval *BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationMessageExpiryInterval `json:"messageExpiryInterval"`
  AdditionalProperties map[string]interface{} `json:"-"`
}