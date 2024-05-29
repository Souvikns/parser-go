
package models

// OneOf_1 represents a OneOf_1 model.
type OneOf_1 struct {
  Asyncapi *Asyncapi `json:"asyncapi,omitempty"`
  Info *Info `json:"info,omitempty"`
  BaseTopic string `json:"baseTopic,omitempty"`
  Servers []Server `json:"servers,omitempty"`
  Topics *Topics `json:"topics,omitempty"`
  Stream *StreamObject `json:"stream,omitempty"`
  Events *EventsObject `json:"events,omitempty"`
  Components *Components `json:"components,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  Security []map[string][]string `json:"security,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}