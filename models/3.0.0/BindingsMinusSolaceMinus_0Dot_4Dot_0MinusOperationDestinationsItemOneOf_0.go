
package models

// BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0 represents a BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0 model.
type BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0 struct {
  DeliveryMode *BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode `json:"deliveryMode"`
  DestinationType string `json:"destinationType"`
  Queue *BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemOneOf_0Queue `json:"queue"`
  AdditionalProperties map[string]interface{} `json:"-"`
}