
package models

// MessageOneOf_1OneOf_1 represents a MessageOneOf_1OneOf_1 model.
type MessageOneOf_1OneOf_1 struct {
  SchemaFormat string `json:"schemaFormat,omitempty"`
  ContentType string `json:"contentType,omitempty"`
  Headers *MessageOneOf_1OneOf_1Headers `json:"headers,omitempty"`
  Payload *MessageOneOf_1OneOf_1Payload `json:"payload,omitempty"`
  CorrelationId *MessageOneOf_1OneOf_1CorrelationId `json:"correlationId,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  Summary string `json:"summary,omitempty"`
  Name string `json:"name,omitempty"`
  Title string `json:"title,omitempty"`
  Description string `json:"description,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  Examples []MessageOneOf_1OneOf_1ExamplesItem `json:"examples,omitempty"`
  Bindings *BindingsObject `json:"bindings,omitempty"`
  Traits []MessageOneOf_1OneOf_1TraitsItem `json:"traits,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}