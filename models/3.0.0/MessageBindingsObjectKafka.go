
package models

// MessageBindingsObjectKafka represents a MessageBindingsObjectKafka model.
type MessageBindingsObjectKafka struct {
  BindingVersion *MessageBindingsObjectKafkaBindingVersion
  Key interface{}
  SchemaIdLocation *BindingsMinusKafkaMinus_0Dot_4Dot_0MinusMessageSchemaIdLocation
  SchemaIdPayloadEncoding string
  SchemaLookupStrategy string
  AdditionalProperties map[string]interface{}
}