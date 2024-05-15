
package models

// MessageTrait represents a MessageTrait model.
type MessageTrait struct {
  SchemaFormat string `json:"schemaFormat"`
  ContentType string `json:"contentType"`
  Headers *MessageTraitHeaders `json:"headers"`
  MessageId string `json:"messageId"`
  CorrelationId *MessageTraitCorrelationId `json:"correlationId"`
  Tags []Tag `json:"tags"`
  Summary string `json:"summary"`
  Name string `json:"name"`
  Title string `json:"title"`
  Description string `json:"description"`
  ExternalDocs *ExternalDocs `json:"externalDocs"`
  Deprecated bool `json:"deprecated"`
  Examples []map[string]interface{} `json:"examples"`
  Bindings *BindingsObject `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}