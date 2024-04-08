
package models

// MessageOneOf_1OneOf_1 represents a MessageOneOf_1OneOf_1 model.
type MessageOneOf_1OneOf_1 struct {
  SchemaFormat string `json:"schemaFormat"`
  ContentType string `json:"contentType"`
  Headers *MessageOneOf_1OneOf_1Headers `json:"headers"`
  MessageId string `json:"messageId"`
  Payload *MessageOneOf_1OneOf_1Payload `json:"payload"`
  CorrelationId *MessageOneOf_1OneOf_1CorrelationId `json:"correlationId"`
  Tags []Tag `json:"tags"`
  Summary string `json:"summary"`
  Name string `json:"name"`
  Title string `json:"title"`
  Description string `json:"description"`
  ExternalDocs *ExternalDocs `json:"externalDocs"`
  Deprecated bool `json:"deprecated"`
  Examples []MessageOneOf_1OneOf_1ExamplesItem `json:"examples"`
  Bindings *BindingsObject `json:"bindings"`
  Traits []MessageOneOf_1OneOf_1TraitsItem `json:"traits"`
  AdditionalProperties map[string]interface{} `json:"-"`
}