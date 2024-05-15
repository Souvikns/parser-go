
package models
import (  
  "encoding/json"
)
// BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType represents an enum of BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType.
type BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType uint

const (
  BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeExchange BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType = iota
  BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeQueue
  BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeFifoMinusQueue
)

// Value returns the value of the enum.
func (op BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType) Value() any {
	if op >= BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType(len(BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues)) {
		return nil
	}
	return BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues[op]
}

var BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues = []any{"exchange","queue","fifo-queue"}
var ValuesToBindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType = map[any]BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType{
  BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues[BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeExchange]: BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeExchange,
  BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues[BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeQueue]: BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeQueue,
  BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeValues[BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeFifoMinusQueue]: BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationTypeFifoMinusQueue,
}

 
func (op *BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType[v]
  return nil
}

func (op BindingsMinusAnypointmqMinus_0Dot_0Dot_1MinusChannelDestinationType) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}