
package models

// MessageObject represents a MessageObject model.
type MessageObject struct {
  ContentType string `json:"contentType,omitempty"`
  Headers *AnySchema `json:"headers,omitempty"`
  Payload *AnySchema `json:"payload,omitempty"`
  CorrelationId *MessageObjectCorrelationId `json:"correlationId,omitempty"`
  Tags []MessageObjectTagsItem `json:"tags,omitempty"`
  Summary string `json:"summary,omitempty"`
  Name string `json:"name,omitempty"`
  Title string `json:"title,omitempty"`
  Description string `json:"description,omitempty"`
  ExternalDocs *MessageObjectExternalDocs `json:"externalDocs,omitempty"`
  Deprecated bool `json:"deprecated,omitempty"`
  Examples []MessageObjectExamplesItem `json:"examples,omitempty"`
  Bindings *MessageObjectBindings `json:"bindings,omitempty"`
  Traits []MessageObjectTraitsItem `json:"traits,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}