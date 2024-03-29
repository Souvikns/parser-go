
package models

// MessageBindingsObjectKafka represents a MessageBindingsObjectKafka model.
type MessageBindingsObjectKafka struct {
  BindingVersion *MessageBindingsObjectKafkaBindingVersion `json:"bindingVersion"`
  Key interface{} `json:"key"`
  SchemaIdLocation *BindingsMinusKafkaMinus_0Dot_4Dot_0MinusMessageSchemaIdLocation `json:"schemaIdLocation"`
  SchemaIdPayloadEncoding string `json:"schemaIdPayloadEncoding"`
  SchemaLookupStrategy string `json:"schemaLookupStrategy"`
  AdditionalProperties map[string]interface{} `json:"-"`
}