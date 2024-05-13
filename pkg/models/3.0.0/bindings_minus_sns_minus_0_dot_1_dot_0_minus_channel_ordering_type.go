
package models
import (  
  "encoding/json"
)
// BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType represents an enum of BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType uint

const (
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeStandard BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType = iota
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeFifo
)

// Value returns the value of the enum.
func (op BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType) Value() any {
	if op >= BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType(len(BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeValues)) {
		return nil
	}
	return BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeValues[op]
}

var BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeValues = []any{"standard","FIFO"}
var ValuesToBindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType = map[any]BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType{
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeStandard]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeStandard,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeFifo]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingTypeFifo,
}

 
func (op *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType[v]
  return nil
}

func (op BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelOrderingType) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}