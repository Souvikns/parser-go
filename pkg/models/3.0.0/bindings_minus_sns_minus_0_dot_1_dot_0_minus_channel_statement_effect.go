
package models
import (  
  "encoding/json"
)
// BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect represents an enum of BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect.
type BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect uint

const (
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectAllow BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect = iota
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectDeny
)

// Value returns the value of the enum.
func (op BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect) Value() any {
	if op >= BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect(len(BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectValues)) {
		return nil
	}
	return BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectValues[op]
}

var BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectValues = []any{"Allow","Deny"}
var ValuesToBindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect = map[any]BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect{
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectAllow]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectAllow,
  BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectValues[BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectDeny]: BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffectDeny,
}

 
func (op *BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect[v]
  return nil
}

func (op BindingsMinusSnsMinus_0Dot_1Dot_0MinusChannelStatementEffect) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}