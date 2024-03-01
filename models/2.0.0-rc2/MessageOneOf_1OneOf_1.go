
package models

// MessageOneOf_1OneOf_1 represents a MessageOneOf_1OneOf_1 model.
type MessageOneOf_1OneOf_1 struct {
  SchemaFormat string `json:schemaFormat`
  ContentType string `json:contentType`
  Headers interface{} `json:headers`
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
  Bindings *BindingsObject `json:bindings`
  Traits []interface{} `json:traits`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}