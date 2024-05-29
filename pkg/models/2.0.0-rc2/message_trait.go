
package models

// MessageTrait represents a MessageTrait model.
type MessageTrait struct {
  SchemaFormat string `json:"schemaFormat,omitempty"`
  ContentType string `json:"contentType,omitempty"`
  Headers *MessageTraitHeaders `json:"headers,omitempty"`
  CorrelationId *MessageTraitCorrelationId `json:"correlationId,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  Summary string `json:"summary,omitempty"`
  Name string `json:"name,omitempty"`
  Title string `json:"title,omitempty"`
  Description string `json:"description,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  Examples []map[string]interface{} `json:"examples,omitempty"`
  Bindings *BindingsObject `json:"bindings,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}