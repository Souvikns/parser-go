
package models

// BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy represents a BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy model.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicy struct {
  MinDelayTarget int `json:"minDelayTarget"`
  MaxDelayTarget int `json:"maxDelayTarget"`
  NumRetries int `json:"numRetries"`
  NumNoDelayRetries int `json:"numNoDelayRetries"`
  NumMinDelayRetries int `json:"numMinDelayRetries"`
  NumMaxDelayRetries int `json:"numMaxDelayRetries"`
  BackoffFunction *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationDeliveryPolicyBackoffFunction `json:"backoffFunction"`
  MaxReceivesPerSecond int `json:"maxReceivesPerSecond"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}