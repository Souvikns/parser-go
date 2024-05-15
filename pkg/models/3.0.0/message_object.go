
package models

// MessageObject represents a MessageObject model.
type MessageObject struct {
  ContentType string `json:"contentType"`
  Headers *AnySchema `json:"headers"`
  Payload *AnySchema `json:"payload"`
  CorrelationId *MessageObjectCorrelationId `json:"correlationId"`
  Tags []MessageObjectTagsItem `json:"tags"`
  Summary string `json:"summary"`
  Name string `json:"name"`
  Title string `json:"title"`
  Description string `json:"description"`
  ExternalDocs *MessageObjectExternalDocs `json:"externalDocs"`
  Deprecated bool `json:"deprecated"`
  Examples []MessageObjectExamplesItem `json:"examples"`
  Bindings *MessageObjectBindings `json:"bindings"`
  Traits []MessageObjectTraitsItem `json:"traits"`
  AdditionalProperties map[string]interface{} `json:"-"`
}