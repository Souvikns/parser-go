
package models

// MessageBindingsObjectKafka represents a MessageBindingsObjectKafka model.
type MessageBindingsObjectKafka struct {
  BindingVersion *MessageBindingsObjectKafkaBindingVersion `json:"bindingVersion"`
  Key *BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageKey `json:"key"`
  SchemaIdLocation *BindingsMinusKafkaMinus_0Dot_5Dot_0MinusMessageSchemaIdLocation `json:"schemaIdLocation"`
  SchemaIdPayloadEncoding string `json:"schemaIdPayloadEncoding"`
  SchemaLookupStrategy string `json:"schemaLookupStrategy"`
  AdditionalProperties map[string]interface{} `json:"-"`
}