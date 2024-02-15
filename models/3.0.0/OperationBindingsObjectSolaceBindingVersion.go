
package models

// OperationBindingsObjectSolaceBindingVersion represents an enum of OperationBindingsObjectSolaceBindingVersion.
type OperationBindingsObjectSolaceBindingVersion uint

const (
  OperationBindingsObjectSolaceBindingVersionNumber_0Dot_4Dot_0 OperationBindingsObjectSolaceBindingVersion = iota
  OperationBindingsObjectSolaceBindingVersionNumber_0Dot_3Dot_0
  OperationBindingsObjectSolaceBindingVersionNumber_0Dot_2Dot_0
)

// Value returns the value of the enum.
func (op OperationBindingsObjectSolaceBindingVersion) Value() any {
	if op >= OperationBindingsObjectSolaceBindingVersion(len(OperationBindingsObjectSolaceBindingVersionValues)) {
		return nil
	}
	return OperationBindingsObjectSolaceBindingVersionValues[op]
}

var OperationBindingsObjectSolaceBindingVersionValues = []any{"0.4.0","0.3.0","0.2.0"}
var ValuesToOperationBindingsObjectSolaceBindingVersion = map[any]OperationBindingsObjectSolaceBindingVersion{
  OperationBindingsObjectSolaceBindingVersionValues[OperationBindingsObjectSolaceBindingVersionNumber_0Dot_4Dot_0]: OperationBindingsObjectSolaceBindingVersionNumber_0Dot_4Dot_0,
  OperationBindingsObjectSolaceBindingVersionValues[OperationBindingsObjectSolaceBindingVersionNumber_0Dot_3Dot_0]: OperationBindingsObjectSolaceBindingVersionNumber_0Dot_3Dot_0,
  OperationBindingsObjectSolaceBindingVersionValues[OperationBindingsObjectSolaceBindingVersionNumber_0Dot_2Dot_0]: OperationBindingsObjectSolaceBindingVersionNumber_0Dot_2Dot_0,
}