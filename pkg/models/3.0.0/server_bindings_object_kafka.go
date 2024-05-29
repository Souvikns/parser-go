
package models

// ServerBindingsObjectKafka represents a ServerBindingsObjectKafka model.
type ServerBindingsObjectKafka struct {
  BindingVersion *ServerBindingsObjectKafkaBindingVersion `json:"bindingVersion,omitempty"`
  SchemaRegistryUrl string `json:"schemaRegistryUrl,omitempty"`
  SchemaRegistryVendor string `json:"schemaRegistryVendor,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}