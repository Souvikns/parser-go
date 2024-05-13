
package models
import (  
  "encoding/json"
)
// BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit represents an enum of BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit uint

const (
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitPerQueue BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit = iota
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitPerMessageGroupId
)

// Value returns the value of the enum.
func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit) Value() any {
	if op >= BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit(len(BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitValues)) {
		return nil
	}
	return BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitValues[op]
}

var BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitValues = []any{"perQueue","perMessageGroupId"}
var ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit = map[any]BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit{
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitPerQueue]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitPerQueue,
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitPerMessageGroupId]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimitPerMessageGroupId,
}

 
func (op *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit[v]
  return nil
}

func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueFifoThroughputLimit) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}