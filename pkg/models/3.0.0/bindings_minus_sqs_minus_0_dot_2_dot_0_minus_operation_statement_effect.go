
package models
import (  
  "encoding/json"
)
// BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect represents an enum of BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect uint

const (
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectAllow BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect = iota
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectDeny
)

// Value returns the value of the enum.
func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect) Value() any {
	if op >= BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect(len(BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectValues)) {
		return nil
	}
	return BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectValues[op]
}

var BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectValues = []any{"Allow","Deny"}
var ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect = map[any]BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect{
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectAllow]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectAllow,
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectDeny]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffectDeny,
}

 
func (op *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect[v]
  return nil
}

func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationStatementEffect) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}