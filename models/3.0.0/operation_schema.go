
package models

// OperationSchema represents a OperationSchema model.
type OperationSchema struct {
  Queues []OperationSchema `json:"queues"`
  BindingVersion *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion `json:"bindingVersion"`
  AdditionalProperties map[string]interface{} `json:"-"`
}