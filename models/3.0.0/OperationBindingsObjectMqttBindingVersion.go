
package models
import (  
  "encoding/json"
)
// OperationBindingsObjectMqttBindingVersion represents an enum of OperationBindingsObjectMqttBindingVersion.
type OperationBindingsObjectMqttBindingVersion uint

const (
  OperationBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0 OperationBindingsObjectMqttBindingVersion = iota
)

// Value returns the value of the enum.
func (op OperationBindingsObjectMqttBindingVersion) Value() any {
	if op >= OperationBindingsObjectMqttBindingVersion(len(OperationBindingsObjectMqttBindingVersionValues)) {
		return nil
	}
	return OperationBindingsObjectMqttBindingVersionValues[op]
}

var OperationBindingsObjectMqttBindingVersionValues = []any{"0.2.0"}
var ValuesToOperationBindingsObjectMqttBindingVersion = map[any]OperationBindingsObjectMqttBindingVersion{
  OperationBindingsObjectMqttBindingVersionValues[OperationBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0]: OperationBindingsObjectMqttBindingVersionNumber_0Dot_2Dot_0,
}

 
          
func (op *OperationBindingsObjectMqttBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToOperationBindingsObjectMqttBindingVersion[v]
	return nil
}

func (op OperationBindingsObjectMqttBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          