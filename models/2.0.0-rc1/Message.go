
package models

// Message represents a Message model.
type Message struct {
  SchemaFormat string `json:schemaFormat`
  ContentType string `json:contentType`
  Headers map[string]interface{} `json:headers`
  Payload interface{} `json:payload`
  CorrelationId interface{} `json:correlationId`
  Tags []Tag `json:tags`
  Summary string `json:summary`
  Name string `json:name`
  Title string `json:title`
  Description string `json:description`
  ExternalDocs *ExternalDocs `json:externalDocs`
  Deprecated bool `json:deprecated`
  Examples []map[string]interface{} `json:examples`
  ProtocolInfo map[string]map[string]interface{} `json:protocolInfo`
  Traits []interface{} `json:traits`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}