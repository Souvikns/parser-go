
package models
import (  
  "encoding/json"
)
// BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode represents an enum of BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode.
type BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode uint

const (
  BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModeDirect BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode = iota
  BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModePersistent
)

// Value returns the value of the enum.
func (op BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode) Value() any {
	if op >= BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode(len(BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModeValues)) {
		return nil
	}
	return BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModeValues[op]
}

var BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModeValues = []any{"direct","persistent"}
var ValuesToBindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode = map[any]BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode{
  BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModeValues[BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModeDirect]: BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModeDirect,
  BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModeValues[BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModePersistent]: BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryModePersistent,
}

 
          
func (op *BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToBindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode[v]
	return nil
}

func (op BindingsMinusSolaceMinus_0Dot_4Dot_0MinusOperationDestinationsItemDeliveryMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          