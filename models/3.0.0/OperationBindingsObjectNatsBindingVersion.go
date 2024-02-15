
package models

// OperationBindingsObjectNatsBindingVersion represents an enum of OperationBindingsObjectNatsBindingVersion.
type OperationBindingsObjectNatsBindingVersion uint

const (
  OperationBindingsObjectNatsBindingVersionNumber_0Dot_1Dot_0 OperationBindingsObjectNatsBindingVersion = iota
)

// Value returns the value of the enum.
func (op OperationBindingsObjectNatsBindingVersion) Value() any {
	if op >= OperationBindingsObjectNatsBindingVersion(len(OperationBindingsObjectNatsBindingVersionValues)) {
		return nil
	}
	return OperationBindingsObjectNatsBindingVersionValues[op]
}

var OperationBindingsObjectNatsBindingVersionValues = []any{"0.1.0"}
var ValuesToOperationBindingsObjectNatsBindingVersion = map[any]OperationBindingsObjectNatsBindingVersion{
  OperationBindingsObjectNatsBindingVersionValues[OperationBindingsObjectNatsBindingVersionNumber_0Dot_1Dot_0]: OperationBindingsObjectNatsBindingVersionNumber_0Dot_1Dot_0,
}
