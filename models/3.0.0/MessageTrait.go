
package models

// MessageTrait represents a MessageTrait model.
type MessageTrait struct {
  ContentType string `json:"contentType"`
  Headers interface{} `json:"headers"`
  CorrelationId interface{} `json:"correlationId"`
  Tags []interface{} `json:"tags"`
  Summary string `json:"summary"`
  Name string `json:"name"`
  Title string `json:"title"`
  Description string `json:"description"`
  ExternalDocs interface{} `json:"externalDocs"`
  Deprecated bool `json:"deprecated"`
  Examples []map[string]interface{} `json:"examples"`
  Bindings interface{} `json:"bindings"`
  AdditionalProperties map[string]interface{} `json:"-"`
}