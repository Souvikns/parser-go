
package models

// OperationBindingsObjectSqs represents a OperationBindingsObjectSqs model.
type OperationBindingsObjectSqs struct {
  BindingVersion *OperationBindingsObjectSqsBindingVersion `json:"bindingVersion,omitempty"`
  Queues []BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationQueue `json:"queues,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}