
package models

// BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatement represents a BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatement model.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatement struct {
  Effect *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect `json:"effect"`
  Principal *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementPrincipal `json:"principal"`
  Action *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementAction `json:"action"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}