
package models

// MessageOneOf_1OneOf_1 represents a MessageOneOf_1OneOf_1 model.
type MessageOneOf_1OneOf_1 struct {
  SchemaFormat string `json:"schemaFormat,omitempty"`
  ContentType string `json:"contentType,omitempty"`
  Headers *Schema `json:"headers,omitempty"`
  Payload interface{} `json:"payload,omitempty"`
  CorrelationId *MessageOneOf_1OneOf_1CorrelationId `json:"correlationId,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  Summary string `json:"summary,omitempty"`
  Name string `json:"name,omitempty"`
  Title string `json:"title,omitempty"`
  Description string `json:"description,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  Examples []map[string]interface{} `json:"examples,omitempty"`
  Bindings *BindingsObject `json:"bindings,omitempty"`
  Traits []MessageOneOf_1OneOf_1TraitsItem `json:"traits,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}