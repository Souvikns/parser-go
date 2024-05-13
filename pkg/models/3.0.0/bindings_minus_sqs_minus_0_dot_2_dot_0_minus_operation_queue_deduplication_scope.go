
package models
import (  
  "encoding/json"
)
// BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope represents an enum of BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope uint

const (
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeQueue BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope = iota
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeMessageGroup
)

// Value returns the value of the enum.
func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope) Value() any {
	if op >= BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope(len(BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeValues)) {
		return nil
	}
	return BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeValues[op]
}

var BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeValues = []any{"queue","messageGroup"}
var ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope = map[any]BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope{
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeQueue]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeQueue,
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeMessageGroup]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScopeMessageGroup,
}

 
func (op *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope[v]
  return nil
}

func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueueDeduplicationScope) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}