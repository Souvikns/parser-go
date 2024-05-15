
package models

// BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationRedrivePolicy represents a BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationRedrivePolicy model.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationRedrivePolicy struct {
  DeadLetterQueue *BindingsMinusSnsMinus_0Dot_1Dot_0MinusOperationIdentifier `json:"deadLetterQueue"`
  MaxReceiveCount int `json:"maxReceiveCount"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}