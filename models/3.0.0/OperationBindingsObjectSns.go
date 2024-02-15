
package models

// OperationBindingsObjectSns represents a OperationBindingsObjectSns model.
type OperationBindingsObjectSns struct {
  BindingVersion string
  Topic *OperationSchema
  Consumers []OperationSchema
  DeliveryPolicy *OperationSchema
  AdditionalProperties map[string]interface{}
}