
package models

// OperationBindingsObjectHttpBindingVersion represents an enum of OperationBindingsObjectHttpBindingVersion.
type OperationBindingsObjectHttpBindingVersion uint

const (
  OperationBindingsObjectHttpBindingVersionNumber_0Dot_2Dot_0 OperationBindingsObjectHttpBindingVersion = iota
  OperationBindingsObjectHttpBindingVersionNumber_0Dot_3Dot_0
)

// Value returns the value of the enum.
func (op OperationBindingsObjectHttpBindingVersion) Value() any {
	if op >= OperationBindingsObjectHttpBindingVersion(len(OperationBindingsObjectHttpBindingVersionValues)) {
		return nil
	}
	return OperationBindingsObjectHttpBindingVersionValues[op]
}

var OperationBindingsObjectHttpBindingVersionValues = []any{"0.2.0","0.3.0"}
var ValuesToOperationBindingsObjectHttpBindingVersion = map[any]OperationBindingsObjectHttpBindingVersion{
  OperationBindingsObjectHttpBindingVersionValues[OperationBindingsObjectHttpBindingVersionNumber_0Dot_2Dot_0]: OperationBindingsObjectHttpBindingVersionNumber_0Dot_2Dot_0,
  OperationBindingsObjectHttpBindingVersionValues[OperationBindingsObjectHttpBindingVersionNumber_0Dot_3Dot_0]: OperationBindingsObjectHttpBindingVersionNumber_0Dot_3Dot_0,
}
