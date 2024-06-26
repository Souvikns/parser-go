
package models
import (  
  "encoding/json"
)
// BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType represents an enum of BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType.
type BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType uint

const (
  BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeQueue BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType = iota
  BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeFifoMinusQueue
)

// Value returns the value of the enum.
func (op BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType) Value() any {
	if op >= BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType(len(BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues)) {
		return nil
	}
	return BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues[op]
}

var BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues = []any{"queue","fifo-queue"}
var ValuesToBindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType = map[any]BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType{
  BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues[BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeQueue]: BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeQueue,
  BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues[BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeFifoMinusQueue]: BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationTypeFifoMinusQueue,
}

 
func (op *BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType[v]
  return nil
}

func (op BindingsMinusJmsMinus_0Dot_0Dot_1MinusChannelDestinationType) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}