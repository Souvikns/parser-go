
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelRedrivePolicy represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelRedrivePolicy model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelRedrivePolicy struct {
  DeadLetterQueue *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelIdentifier `json:"deadLetterQueue"`
  MaxReceiveCount int `json:"maxReceiveCount"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}