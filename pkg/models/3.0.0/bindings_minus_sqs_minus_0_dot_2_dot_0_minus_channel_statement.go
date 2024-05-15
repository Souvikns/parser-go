
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatement represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatement model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatement struct {
  Effect *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect `json:"effect"`
  Principal *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementPrincipal `json:"principal"`
  Action *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementAction `json:"action"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}