
package models

// ServerBindingsObjectKafka represents a ServerBindingsObjectKafka model.
type ServerBindingsObjectKafka struct {
  BindingVersion *ServerBindingsObjectKafkaBindingVersion
  SchemaRegistryUrl string
  SchemaRegistryVendor string
  AdditionalProperties map[string]interface{}
}