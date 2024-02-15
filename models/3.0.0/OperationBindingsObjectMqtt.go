
package models

// OperationBindingsObjectMqtt represents a OperationBindingsObjectMqtt model.
type OperationBindingsObjectMqtt struct {
  BindingVersion *OperationBindingsObjectMqttBindingVersion
  Qos *BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos
  Retain bool
  MessageExpiryInterval interface{}
  AdditionalProperties map[string]interface{}
}