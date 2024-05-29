
package models

// Message represents a Message model.
type Message struct {
  Ref string `json:"$ref,omitempty"`
  Headers *Schema `json:"headers,omitempty"`
  Payload *Schema `json:"payload,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  Summary string `json:"summary,omitempty"`
  Description string `json:"description,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  Example interface{} `json:"example,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}