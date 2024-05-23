
package models

// AsyncApi_1Dot_0SchemaDot represents a AsyncApi_1Dot_0SchemaDot model.
type AsyncApi_1Dot_0SchemaDot struct {
  Asyncapi *Asyncapi `json:"asyncapi,omitempty"`
  Info *Info `json:"info,omitempty"`
  BaseTopic string `json:"baseTopic,omitempty"`
  Servers []Server `json:"servers,omitempty"`
  Topics *Topics `json:"topics,omitempty"`
  Components *Components `json:"components,omitempty"`
  Tags []Tag `json:"tags,omitempty"`
  Security []map[string][]string `json:"security,omitempty"`
  ExternalDocs *ExternalDocs `json:"externalDocs,omitempty"`
  AdditionalProperties map[string]interface{} `json:"-",omitempty`
}