
package models

// MessageTrait represents a MessageTrait model.
type MessageTrait struct {
  SchemaFormat string `json:"schemaFormat,omitempty"`
  ContentType string `json:"contentType,omitempty"`
  Headers map[string]MessageTraitHeadersAdditionalProperty `json:"headers,omitempty"`
  CorrelationId *MessageTraitCorrelationId `json:"correlationId,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  Summary string `json:"summary,omitempty"`
  Name string `json:"name,omitempty"`
  Title string `json:"title,omitempty"`
  Description string `json:"description,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  Examples []map[string]interface{} `json:"examples,omitempty"`
  ProtocolInfo map[string]map[string]interface{} `json:"protocolInfo,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}