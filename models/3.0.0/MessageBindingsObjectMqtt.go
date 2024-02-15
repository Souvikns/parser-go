
package models

// MessageBindingsObjectMqtt represents a MessageBindingsObjectMqtt model.
type MessageBindingsObjectMqtt struct {
  BindingVersion *MessageBindingsObjectMqttBindingVersion
  PayloadFormatIndicator *BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator
  CorrelationData interface{}
  ContentType string
  ResponseTopic interface{}
  AdditionalProperties map[string]interface{}
}