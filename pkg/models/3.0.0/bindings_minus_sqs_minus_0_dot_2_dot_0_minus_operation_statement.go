
package models

// BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatement represents a BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatement model.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatement struct {
  Effect *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect `json:"effect"`
  Principal *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementPrincipal `json:"principal"`
  Action *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementAction `json:"action"`
  AdditionalProperties map[string]AdditionalProperties `json:"-"`
}