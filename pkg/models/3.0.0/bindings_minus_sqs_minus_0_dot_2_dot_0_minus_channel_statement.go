
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatement represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatement model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatement struct {
  Effect *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect `json:"effect,omitempty"`
  Principal *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementPrincipal `json:"principal,omitempty"`
  Action *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementAction `json:"action,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-,omitempty"`
}