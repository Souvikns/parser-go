
package models
import (  
  "encoding/json"
)
// BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod represents an enum of BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod.
type BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod uint

const (
  BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodGet BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod = iota
  BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodPost
)

// Value returns the value of the enum.
func (op BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod) Value() any {
	if op >= BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod(len(BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodValues)) {
		return nil
	}
	return BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodValues[op]
}

var BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodValues = []any{"GET","POST"}
var ValuesToBindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod = map[any]BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod{
  BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodValues[BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodGet]: BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodGet,
  BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodValues[BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodPost]: BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethodPost,
}

 
func (op *BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToBindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod[v]
  return nil
}

func (op BindingsMinusWebsocketsMinus_0Dot_1Dot_0MinusChannelMethod) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}