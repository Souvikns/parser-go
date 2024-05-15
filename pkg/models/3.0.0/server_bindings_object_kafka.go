
package models

// ServerBindingsObjectKafka represents a ServerBindingsObjectKafka model.
type ServerBindingsObjectKafka struct {
  BindingVersion *ServerBindingsObjectKafkaBindingVersion `json:"bindingVersion"`
  SchemaRegistryUrl string `json:"schemaRegistryUrl"`
  SchemaRegistryVendor string `json:"schemaRegistryVendor"`
  AdditionalProperties map[string]interface{} `json:"-"`
}