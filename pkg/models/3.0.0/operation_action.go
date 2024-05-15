
package models
import (  
  "encoding/json"
)
// OperationAction represents an enum of OperationAction.
type OperationAction uint

const (
  OperationActionSend OperationAction = iota
  OperationActionReceive
)

// Value returns the value of the enum.
func (op OperationAction) Value() any {
	if op >= OperationAction(len(OperationActionValues)) {
		return nil
	}
	return OperationActionValues[op]
}

var OperationActionValues = []any{"send","receive"}
var ValuesToOperationAction = map[any]OperationAction{
  OperationActionValues[OperationActionSend]: OperationActionSend,
  OperationActionValues[OperationActionReceive]: OperationActionReceive,
}

 
func (op *OperationAction) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToOperationAction[v]
  return nil
}

func (op OperationAction) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}