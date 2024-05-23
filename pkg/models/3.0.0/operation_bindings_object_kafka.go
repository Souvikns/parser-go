
package models

// OperationBindingsObjectKafka represents a OperationBindingsObjectKafka model.
type OperationBindingsObjectKafka struct {
  BindingVersion *OperationBindingsObjectKafkaBindingVersion `json:"bindingVersion,omitempty"`
  GroupId *Schema `json:"groupId,omitempty"`
  ClientId *Schema `json:"clientId,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}