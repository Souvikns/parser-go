
package models

// BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationRedrivePolicy represents a BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationRedrivePolicy model.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationRedrivePolicy struct {
  DeadLetterQueue *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationIdentifier `json:"deadLetterQueue,omitempty"`
  MaxReceiveCount int `json:"maxReceiveCount,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-,omitempty"`
}