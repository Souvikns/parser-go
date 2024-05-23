
package models

// BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_1 represents a BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_1 model.
type BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_1 struct {
  DeliveryMode *BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode `json:"deliveryMode,omitempty"`
  DestinationType string `json:"destinationType,omitempty"`
  TopicSubscriptions []string `json:"topicSubscriptions,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}