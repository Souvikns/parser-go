
package models
import (  
  "encoding/json"
)
// OperationBindingsObjectSqsBindingVersion represents an enum of OperationBindingsObjectSqsBindingVersion.
type OperationBindingsObjectSqsBindingVersion uint

const (
  OperationBindingsObjectSqsBindingVersionNumber_0Dot_2Dot_0 OperationBindingsObjectSqsBindingVersion = iota
  OperationBindingsObjectSqsBindingVersionNumber_0Dot_1Dot_0
)

// Value returns the value of the enum.
func (op OperationBindingsObjectSqsBindingVersion) Value() any {
	if op >= OperationBindingsObjectSqsBindingVersion(len(OperationBindingsObjectSqsBindingVersionValues)) {
		return nil
	}
	return OperationBindingsObjectSqsBindingVersionValues[op]
}

var OperationBindingsObjectSqsBindingVersionValues = []any{"0.2.0","0.1.0"}
var ValuesToOperationBindingsObjectSqsBindingVersion = map[any]OperationBindingsObjectSqsBindingVersion{
  OperationBindingsObjectSqsBindingVersionValues[OperationBindingsObjectSqsBindingVersionNumber_0Dot_2Dot_0]: OperationBindingsObjectSqsBindingVersionNumber_0Dot_2Dot_0,
  OperationBindingsObjectSqsBindingVersionValues[OperationBindingsObjectSqsBindingVersionNumber_0Dot_1Dot_0]: OperationBindingsObjectSqsBindingVersionNumber_0Dot_1Dot_0,
}

 
func (op *OperationBindingsObjectSqsBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToOperationBindingsObjectSqsBindingVersion[v]
  return nil
}

func (op OperationBindingsObjectSqsBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}