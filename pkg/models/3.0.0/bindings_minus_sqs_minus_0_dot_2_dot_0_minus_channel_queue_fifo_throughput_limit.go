
package models
import (  
  "encoding/json"
)
// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit represents an enum of BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit uint

const (
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitPerQueue BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit = iota
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitPerMessageGroupId
)

// Value returns the value of the enum.
func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit) Value() any {
	if op >= BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit(len(BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitValues)) {
		return nil
	}
	return BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitValues[op]
}

var BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitValues = []any{"perQueue","perMessageGroupId"}
var ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit = map[any]BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit{
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitPerQueue]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitPerQueue,
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitPerMessageGroupId]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimitPerMessageGroupId,
}

 
func (op *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit[v]
  return nil
}

func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueFifoThroughputLimit) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}