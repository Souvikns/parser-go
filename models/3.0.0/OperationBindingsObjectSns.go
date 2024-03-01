
package models

// OperationBindingsObjectSns represents a OperationBindingsObjectSns model.
type OperationBindingsObjectSns struct {
  BindingVersion string `json:"bindingVersion"`
  Topic *OperationSchema `json:"topic"`
  Consumers []OperationSchema `json:"consumers"`
  DeliveryPolicy *OperationSchema `json:"deliveryPolicy"`
  AdditionalProperties map[string]interface{} `json:"-"`
}