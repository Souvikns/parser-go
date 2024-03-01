
package models
import (  
  "encoding/json"
)
// OperationBindingsObjectKafkaBindingVersion represents an enum of OperationBindingsObjectKafkaBindingVersion.
type OperationBindingsObjectKafkaBindingVersion uint

const (
  OperationBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0 OperationBindingsObjectKafkaBindingVersion = iota
  OperationBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0
)

// Value returns the value of the enum.
func (op OperationBindingsObjectKafkaBindingVersion) Value() any {
	if op >= OperationBindingsObjectKafkaBindingVersion(len(OperationBindingsObjectKafkaBindingVersionValues)) {
		return nil
	}
	return OperationBindingsObjectKafkaBindingVersionValues[op]
}

var OperationBindingsObjectKafkaBindingVersionValues = []any{"0.4.0","0.3.0"}
var ValuesToOperationBindingsObjectKafkaBindingVersion = map[any]OperationBindingsObjectKafkaBindingVersion{
  OperationBindingsObjectKafkaBindingVersionValues[OperationBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0]: OperationBindingsObjectKafkaBindingVersionNumber_0Dot_4Dot_0,
  OperationBindingsObjectKafkaBindingVersionValues[OperationBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0]: OperationBindingsObjectKafkaBindingVersionNumber_0Dot_3Dot_0,
}

 
          
func (op *OperationBindingsObjectKafkaBindingVersion) UnmarshalJSON(raw []byte) error {
	var v any
	if err := json.Unmarshal(raw, &v); err != nil {
		return err
	}
	*op = ValuesToOperationBindingsObjectKafkaBindingVersion[v]
	return nil
}

func (op OperationBindingsObjectKafkaBindingVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(op.Value())
} 
          