
package models

// OperationSchema represents a OperationSchema model.
type OperationSchema struct {
  Queues []OperationSchema
  BindingVersion *BindingsMinusSqsMinus_0Dot_2Dot_0MinusOperationBindingVersion
  AdditionalProperties map[string]interface{}
}