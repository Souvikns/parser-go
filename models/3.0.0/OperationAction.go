
package models

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
