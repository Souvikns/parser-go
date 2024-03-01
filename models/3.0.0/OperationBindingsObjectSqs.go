
package models

// OperationBindingsObjectSqs represents a OperationBindingsObjectSqs model.
type OperationBindingsObjectSqs struct {
  BindingVersion *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion `json:"bindingVersion"`
  Queues []OperationSchema `json:"queues"`
  AdditionalProperties map[string]interface{} `json:"-"`
}