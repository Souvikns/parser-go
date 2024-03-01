
package models

// MessageObject represents a MessageObject model.
type MessageObject struct {
  ContentType string `json:"contentType"`
  Headers interface{} `json:"headers"`
  Payload interface{} `json:"payload"`
  CorrelationId interface{} `json:"correlationId"`
  Tags []interface{} `json:"tags"`
  Summary string `json:"summary"`
  Name string `json:"name"`
  Title string `json:"title"`
  Description string `json:"description"`
  ExternalDocs interface{} `json:"externalDocs"`
  Deprecated bool `json:"deprecated"`
  Examples []interface{} `json:"examples"`
  Bindings interface{} `json:"bindings"`
  Traits []interface{} `json:"traits"`
  AdditionalProperties map[string]interface{} `json:"-"`
}