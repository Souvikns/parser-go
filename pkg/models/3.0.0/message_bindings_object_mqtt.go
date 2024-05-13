
package models

// MessageBindingsObjectMqtt represents a MessageBindingsObjectMqtt model.
type MessageBindingsObjectMqtt struct {
  BindingVersion *MessageBindingsObjectMqttBindingVersion `json:"bindingVersion"`
  PayloadFormatIndicator *BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator `json:"payloadFormatIndicator"`
  CorrelationData *BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessageCorrelationData `json:"correlationData"`
  ContentType string `json:"contentType"`
  ResponseTopic *BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessageResponseTopic `json:"responseTopic"`
  AdditionalProperties map[string]interface{} `json:"-"`
}