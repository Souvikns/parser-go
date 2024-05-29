
package models

// OperationBindingsObjectSns represents a OperationBindingsObjectSns model.
type OperationBindingsObjectSns struct {
  BindingVersion *OperationBindingsObjectSnsBindingVersion `json:"bindingVersion,omitempty"`
  Topic *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationIdentifier `json:"topic,omitempty"`
  Consumers []BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumer `json:"consumers,omitempty"`
  DeliveryPolicy *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy `json:"deliveryPolicy,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}