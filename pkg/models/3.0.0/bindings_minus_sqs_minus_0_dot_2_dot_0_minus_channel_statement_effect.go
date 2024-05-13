
package models
import (  
  "encoding/json"
)
// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect represents an enum of BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect uint

const (
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectAllow BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect = iota
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectDeny
)

// Value returns the value of the enum.
func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect) Value() any {
	if op >= BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect(len(BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectValues)) {
		return nil
	}
	return BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectValues[op]
}

var BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectValues = []any{"Allow","Deny"}
var ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect = map[any]BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect{
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectAllow]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectAllow,
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectDeny]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffectDeny,
}

 
func (op *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect[v]
  return nil
}

func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelStatementEffect) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}