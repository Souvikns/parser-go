
package models

// BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumer represents a BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumer model.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumer struct {
  Protocol *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol `json:"protocol"`
  Endpoint *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationIdentifier `json:"endpoint"`
  FilterPolicy map[string]BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicy `json:"filterPolicy"`
  FilterPolicyScope *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope `json:"filterPolicyScope"`
  RawMessageDelivery bool `json:"rawMessageDelivery"`
  RedrivePolicy *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationRedrivePolicy `json:"redrivePolicy"`
  DeliveryPolicy *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy `json:"deliveryPolicy"`
  DisplayName string `json:"displayName"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}