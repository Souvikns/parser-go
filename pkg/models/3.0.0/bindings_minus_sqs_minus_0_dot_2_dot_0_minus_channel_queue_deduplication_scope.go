
package models
import (  
  "encoding/json"
)
// BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope represents an enum of BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope.
type BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope uint

const (
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeQueue BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope = iota
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeMessageGroup
)

// Value returns the value of the enum.
func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope) Value() any {
	if op >= BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope(len(BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeValues)) {
		return nil
	}
	return BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeValues[op]
}

var BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeValues = []any{"queue","messageGroup"}
var ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope = map[any]BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope{
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeQueue]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeQueue,
  BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeValues[BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeMessageGroup]: BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScopeMessageGroup,
}

 
func (op *BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope[v]
  return nil
}

func (op BindingsMinusSqsMinus_0Dot_2Dot_0MinusChannelQueueDeduplicationScope) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}