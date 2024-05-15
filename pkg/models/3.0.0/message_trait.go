
package models

// MessageTrait represents a MessageTrait model.
type MessageTrait struct {
  ContentType string `json:"contentType"`
  Headers *AnySchema `json:"headers"`
  CorrelationId *MessageTraitCorrelationId `json:"correlationId"`
  Tags []MessageTraitTagsItem `json:"tags"`
  Summary string `json:"summary"`
  Name string `json:"name"`
  Title string `json:"title"`
  Description string `json:"description"`
  ExternalDocs *MessageTraitExternalDocs `json:"externalDocs"`
  Deprecated bool `json:"deprecated"`
  Examples []map[string]interface{} `json:"examples"`
  Bindings *MessageTraitBindings `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}