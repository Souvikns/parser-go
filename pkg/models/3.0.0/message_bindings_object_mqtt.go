
package models

// MessageBindingsObjectMqtt represents a MessageBindingsObjectMqtt model.
type MessageBindingsObjectMqtt struct {
  BindingVersion *MessageBindingsObjectMqttBindingVersion `json:"bindingVersion,omitempty"`
  PayloadFormatIndicator *BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator `json:"payloadFormatIndicator,omitempty"`
  CorrelationData *BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessageCorrelationData `json:"correlationData,omitempty"`
  ContentType string `json:"contentType,omitempty"`
  ResponseTopic *BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessageResponseTopic `json:"responseTopic,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}