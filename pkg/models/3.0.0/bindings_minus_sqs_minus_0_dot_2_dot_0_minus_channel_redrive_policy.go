
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelRedrivePolicy represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelRedrivePolicy model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelRedrivePolicy struct {
  DeadLetterQueue *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelIdentifier `json:"deadLetterQueue,omitempty"`
  MaxReceiveCount int `json:"maxReceiveCount,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-",omitempty`
}