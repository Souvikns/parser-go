
package models

// MessageBindingsObjectKafka represents a MessageBindingsObjectKafka model.
type MessageBindingsObjectKafka struct {
  BindingVersion *MessageBindingsObjectKafkaBindingVersion `json:"bindingVersion,omitempty"`
  Key *BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageKey `json:"key,omitempty"`
  SchemaIdLocation *BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation `json:"schemaIdLocation,omitempty"`
  SchemaIdPayloadEncoding string `json:"schemaIdPayloadEncoding,omitempty"`
  SchemaLookupStrategy string `json:"schemaLookupStrategy,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}