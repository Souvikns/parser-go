
package models

// BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrdering represents a BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrdering model.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrdering struct {
  ReservedType *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType `json:"type,omitempty"`
  ContentBasedDeduplication bool `json:"contentBasedDeduplication,omitempty"`
  AdditionalProperties map[string]AdditionalProperties `json:"-",omitempty`
}