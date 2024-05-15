
package models

// OperationBindingsObjectKafka represents a OperationBindingsObjectKafka model.
type OperationBindingsObjectKafka struct {
  BindingVersion *OperationBindingsObjectKafkaBindingVersion `json:"bindingVersion"`
  GroupId *Schema `json:"groupId"`
  ClientId *Schema `json:"clientId"`
  AdditionalProperties map[string]interface{} `json:"-"`
}