
package models

// BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy represents a BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy model.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy struct {
  MinDelayTarget int `json:"minDelayTarget,omitempty"`
  MaxDelayTarget int `json:"maxDelayTarget,omitempty"`
  NumRetries int `json:"numRetries,omitempty"`
  NumNoDelayRetries int `json:"numNoDelayRetries,omitempty"`
  NumMinDelayRetries int `json:"numMinDelayRetries,omitempty"`
  NumMaxDelayRetries int `json:"numMaxDelayRetries,omitempty"`
  BackoffFunction *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicyBackoffFunction `json:"backoffFunction,omitempty"`
  MaxReceivesPerSecond int `json:"maxReceivesPerSecond,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-",omitempty`
}