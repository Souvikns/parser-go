
package models

// OperationBindingsObjectSqs represents a OperationBindingsObjectSqs model.
type OperationBindingsObjectSqs struct {
  BindingVersion *OperationBindingsObjectSqsBindingVersion `json:"bindingVersion"`
  Queues []BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueue `json:"queues"`
  AdditionalProperties map[string]interface{} `json:"-"`
}