
package models
import (  
  "encoding/json"
)
// BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode represents an enum of BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode.
type BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode uint

const (
  BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeNumber_1 BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode = iota
  BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeNumber_2
)

// Value returns the value of the enum.
func (op BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode) Value() any {
	if op >= BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode(len(BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeValues)) {
		return nil
	}
	return BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeValues[op]
}

var BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeValues = []any{1,2}
var ValuesToBindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode = map[any]BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode{
  BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeValues[BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeNumber_1]: BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeNumber_1,
  BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeValues[BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeNumber_2]: BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryModeNumber_2,
}

 
func (op *BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode[v]
  return nil
}

func (op BindingsMinusAmqpMinus_0Dot_3Dot_0MinusOperationDeliveryMode) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}