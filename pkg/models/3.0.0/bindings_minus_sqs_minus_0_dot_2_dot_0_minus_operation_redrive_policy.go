
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationRedrivePolicy represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationRedrivePolicy model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationRedrivePolicy struct {
  DeadLetterQueue *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationIdentifier `json:"deadLetterQueue,omitempty"`
  MaxReceiveCount int `json:"maxReceiveCount,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-,omitempty"`
}