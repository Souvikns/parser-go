
package models
import (  
  "encoding/json"
)
// BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator represents an enum of BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator.
type BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator uint

const (
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorNumber_0 BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator = iota
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorNumber_1
)

// Value returns the value of the enum.
func (op BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator) Value() any {
	if op >= BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator(len(BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorValues)) {
		return nil
	}
	return BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorValues[op]
}

var BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorValues = []any{0,1}
var ValuesToBindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator = map[any]BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator{
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorValues[BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorNumber_0]: BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorNumber_0,
  BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorValues[BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorNumber_1]: BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicatorNumber_1,
}

 
          
func (op *BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToBindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator[v]
	return nil
}

func (op BindingsMinusMqttMinus_0Dot_2Dot_0MinusMessagePayloadFormatIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          