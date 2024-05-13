
package models

// OperationBindingsObjectSns represents a OperationBindingsObjectSns model.
type OperationBindingsObjectSns struct {
  BindingVersion *OperationBindingsObjectSnsBindingVersion `json:"bindingVersion"`
  Topic *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationIdentifier `json:"topic"`
  Consumers []BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumer `json:"consumers"`
  DeliveryPolicy *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy `json:"deliveryPolicy"`
  AdditionalProperties map[string]interface{} `json:"-"`
}