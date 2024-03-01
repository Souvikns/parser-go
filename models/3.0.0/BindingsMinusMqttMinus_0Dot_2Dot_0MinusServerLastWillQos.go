
package models
import (  
  "encoding/json"
)
// BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos represents an enum of BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos.
type BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos uint

const (
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosNumber_0 BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos = iota
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosNumber_1
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosNumber_2
)

// Value returns the value of the enum.
func (op BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos) Value() any {
	if op >= BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos(len(BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosValues)) {
		return nil
	}
	return BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosValues[op]
}

var BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosValues = []any{0,1,2}
var ValuesToBindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos = map[any]BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos{
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosValues[BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosNumber_0]: BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosNumber_0,
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosValues[BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosNumber_1]: BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosNumber_1,
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosValues[BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosNumber_2]: BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQosNumber_2,
}

 
          
func (op *BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToBindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos[v]
	return nil
}

func (op BindingsMinusMqttMinus_0Dot_2Dot_0MinusServerLastWillQos) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          