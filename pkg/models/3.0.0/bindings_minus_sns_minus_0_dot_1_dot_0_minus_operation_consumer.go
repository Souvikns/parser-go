
package models

// BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumer represents a BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumer model.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumer struct {
  Protocol *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerProtocol `json:"protocol,omitempty"`
  Endpoint *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationIdentifier `json:"endpoint,omitempty"`
  FilterPolicy map[string]BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicy `json:"filterPolicy,omitempty"`
  FilterPolicyScope *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationConsumerFilterPolicyScope `json:"filterPolicyScope,omitempty"`
  RawMessageDelivery bool `json:"rawMessageDelivery,omitempty"`
  RedrivePolicy *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationRedrivePolicy `json:"redrivePolicy,omitempty"`
  DeliveryPolicy *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy `json:"deliveryPolicy,omitempty"`
  DisplayName string `json:"displayName,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-",omitempty`
}