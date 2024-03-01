
package models
import (  
  "encoding/json"
)
// BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos represents an enum of BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos.
type BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos uint

const (
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosNumber_0 BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos = iota
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosNumber_1
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosNumber_2
)

// Value returns the value of the enum.
func (op BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos) Value() any {
	if op >= BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos(len(BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosValues)) {
		return nil
	}
	return BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosValues[op]
}

var BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosValues = []any{0,1,2}
var ValuesToBindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos = map[any]BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos{
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosValues[BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosNumber_0]: BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosNumber_0,
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosValues[BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosNumber_1]: BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosNumber_1,
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosValues[BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosNumber_2]: BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQosNumber_2,
}

 
          
func (op *BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToBindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos[v]
	return nil
}

func (op BindingsMinusMqttMinus_0Dot_2Dot_0MinusOperationQos) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          