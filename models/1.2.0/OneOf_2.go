
package models

// OneOf_2 represents a OneOf_2 model.
type OneOf_2 struct {
  Asyncapi *Asyncapi
  Info *Info
  BaseTopic string
  Servers []Server
  Topics *Topics
  Stream *StreamObject
  Events interface{}
  Components *Components
  Tags []Tag
  Security []map[string][]string
  ExternalDocs *ExternalDocs
  AdditionalProperties map[string]interface{}
}