
package models

// Message represents a Message model.
type Message struct {
  SchemaFormat string `json:"schemaFormat,omitempty"`
  ContentType string `json:"contentType,omitempty"`
  Headers map[string]MessageHeadersAdditionalProperty `json:"headers,omitempty"`
  Payload interface{} `json:"payload,omitempty"`
  CorrelationId *MessageCorrelationId `json:"correlationId,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  Summary string `json:"summary,omitempty"`
  Name string `json:"name,omitempty"`
  Title string `json:"title,omitempty"`
  Description string `json:"description,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  Examples []map[string]interface{} `json:"examples,omitempty"`
  ProtocolInfo map[string]map[string]interface{} `json:"protocolInfo,omitempty"`
  Traits []MessageTraitsItem `json:"traits,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}