
package models

// OperationBindingsObjectAmqpBindingVersion represents an enum of OperationBindingsObjectAmqpBindingVersion.
type OperationBindingsObjectAmqpBindingVersion uint

const (
  OperationBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0 OperationBindingsObjectAmqpBindingVersion = iota
)

// Value returns the value of the enum.
func (op OperationBindingsObjectAmqpBindingVersion) Value() any {
	if op >= OperationBindingsObjectAmqpBindingVersion(len(OperationBindingsObjectAmqpBindingVersionValues)) {
		return nil
	}
	return OperationBindingsObjectAmqpBindingVersionValues[op]
}

var OperationBindingsObjectAmqpBindingVersionValues = []any{"0.3.0"}
var ValuesToOperationBindingsObjectAmqpBindingVersion = map[any]OperationBindingsObjectAmqpBindingVersion{
  OperationBindingsObjectAmqpBindingVersionValues[OperationBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0]: OperationBindingsObjectAmqpBindingVersionNumber_0Dot_3Dot_0,
}
