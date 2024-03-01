
package models

// Message represents a Message model.
type Message struct {
  Ref string `json:$ref`
  Headers *Schema `json:headers`
  Payload *Schema `json:payload`
  Tags []Tag `json:tags`
  Summary string `json:summary`
  Description string `json:description`
  ExternalDocs *ExternalDocs `json:externalDocs`
  Deprecated bool `json:deprecated`
  Example interface{} `json:example`
  AdditionalProperties map[string]interface{} `json:additionalProperties`
}