
package models

// BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatement represents a BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatement model.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatement struct {
  Effect *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect `json:"effect,omitempty"`
  Principal *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementPrincipal `json:"principal,omitempty"`
  Action *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementAction `json:"action,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-,omitempty"`
}