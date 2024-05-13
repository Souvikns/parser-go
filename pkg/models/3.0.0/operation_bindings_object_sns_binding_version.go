
package models
import (  
  "encoding/json"
)
// OperationBindingsObjectSnsBindingVersion represents an enum of OperationBindingsObjectSnsBindingVersion.
type OperationBindingsObjectSnsBindingVersion uint

const (
  OperationBindingsObjectSnsBindingVersionNumber_0Dot_1Dot_0 OperationBindingsObjectSnsBindingVersion = iota
)

// Value returns the value of the enum.
func (op OperationBindingsObjectSnsBindingVersion) Value() any {
	if op >= OperationBindingsObjectSnsBindingVersion(len(OperationBindingsObjectSnsBindingVersionValues)) {
		return nil
	}
	return OperationBindingsObjectSnsBindingVersionValues[op]
}

var OperationBindingsObjectSnsBindingVersionValues = []any{"0.1.0"}
var ValuesToOperationBindingsObjectSnsBindingVersion = map[any]OperationBindingsObjectSnsBindingVersion{
  OperationBindingsObjectSnsBindingVersionValues[OperationBindingsObjectSnsBindingVersionNumber_0Dot_1Dot_0]: OperationBindingsObjectSnsBindingVersionNumber_0Dot_1Dot_0,
}

 
func (op *OperationBindingsObjectSnsBindingVersion) UnmarshalJSON(raw []byte) error {
  var v any
  if err := json.Unmarshal(raw, &v); err != nil {
  return err
  }
  *op = ValuesToOperationBindingsObjectSnsBindingVersion[v]
  return nil
}

func (op OperationBindingsObjectSnsBindingVersion) MarshalJSON() ([]byte, error) {
  return json.Marshal(op.Value())
}