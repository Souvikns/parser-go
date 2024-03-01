
package models
import (  
  "encoding/json"
)
// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion represents an enum of BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion uint

const (
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionNumber_0Dot_1Dot_0 BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion = iota
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionNumber_0Dot_2Dot_0
)

// Value returns the value of the enum.
func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion) Value() any {
	if op >= BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion(len(BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionValues)) {
		return nil
	}
	return BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionValues[op]
}

var BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionValues = []any{"0.1.0","0.2.0"}
var ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion = map[any]BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion{
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionNumber_0Dot_1Dot_0]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionNumber_0Dot_1Dot_0,
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionNumber_0Dot_2Dot_0]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersionNumber_0Dot_2Dot_0,
}

 
          
func (op *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion[v]
	return nil
}

func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          