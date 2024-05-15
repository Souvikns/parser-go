
package models

// MessageTrait represents a MessageTrait model.
type MessageTrait struct {
  SchemaFormat string `json:"schemaFormat"`
  ContentType string `json:"contentType"`
  Headers map[string]MessageTraitHeadersAdditionalProperty `json:"headers"`
  CorrelationId *MessageTraitCorrelationId `json:"correlationId"`
  Tags []Tag `json:"tags"`
  Summary string `json:"summary"`
  Name string `json:"name"`
  Title string `json:"title"`
  Description string `json:"description"`
  ExternalDocs *ExternalDocs `json:"externalDocs"`
  Deprecated bool `json:"deprecated"`
  Examples []map[string]interface{} `json:"examples"`
  ProtocolInfo map[string]map[string]interface{} `json:"protocolInfo"`
  AdditionalProperties map[string]interface{} `json:"-"`
}