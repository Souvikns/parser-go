
package models

// OperationBindingsObjectKafka represents a OperationBindingsObjectKafka model.
type OperationBindingsObjectKafka struct {
  BindingVersion *OperationBindingsObjectKafkaBindingVersion `json:"bindingVersion"`
  GroupId interface{} `json:"groupId"`
  ClientId interface{} `json:"clientId"`
  AdditionalProperties map[string]interface{} `json:"-"`
}