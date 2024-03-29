
package models
import (  
  "encoding/json"
)
// BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion represents an enum of BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion uint

const (
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionNumber_0Dot_1Dot_0 BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion = iota
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionNumber_0Dot_2Dot_0
)

// Value returns the value of the enum.
func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion) Value() any {
	if op >= BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion(len(BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionValues)) {
		return nil
	}
	return BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionValues[op]
}

var BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionValues = []any{"0.1.0","0.2.0"}
var ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion = map[any]BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion{
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionNumber_0Dot_1Dot_0]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionNumber_0Dot_1Dot_0,
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionNumber_0Dot_2Dot_0]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersionNumber_0Dot_2Dot_0,
}

 
          
func (op *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion[v]
	return nil
}

func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          